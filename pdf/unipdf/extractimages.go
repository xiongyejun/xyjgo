package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unipdf/extractor"
	"github.com/unidoc/unipdf/model"
)

// Extracts images and properties of a PDF specified by inputPath.
func extractImages(inputPath, outputPath string) (err error) {
	if outputPath, err = filepath.Abs(outputPath); err != nil {
		return err
	}

	tmp := string(os.PathSeparator)
	if !strings.HasSuffix(outputPath, tmp) {
		outputPath += tmp
	}

	var f *os.File
	if f, err = os.Open(inputPath); err != nil {
		return err
	}

	defer f.Close()

	var pdfReader *model.PdfReader
	if pdfReader, err = model.NewPdfReader(f); err != nil {
		return err
	}

	var isEncrypted bool
	if isEncrypted, err = pdfReader.IsEncrypted(); err != nil {
		return err
	}

	// Try decrypting with an empty one.
	if isEncrypted {
		var auth bool
		if auth, err = pdfReader.Decrypt([]byte("")); err != nil {
			// Encrypted and we cannot do anything about it.
			return err
		}
		if !auth {
			fmt.Println("Need to decrypt with password")
			return nil
		}
	}

	var numPages int
	if numPages, err = pdfReader.GetNumPages(); err != nil {
		return err
	}
	fmt.Printf("PDF Num Pages: %d\n", numPages)

	totalImages := 0
	for i := 0; i < numPages; i++ {
		fmt.Printf("-----\nPage %d:\n", i+1)

		var page *model.PdfPage
		if page, err = pdfReader.GetPage(i + 1); err != nil {
			return err
		}

		var pextract *extractor.Extractor
		if pextract, err = extractor.New(page); err != nil {
			return err
		}

		var pimages *extractor.PageImages
		if pimages, err = pextract.ExtractPageImages(nil); err != nil {
			return err
		}

		fmt.Printf("%d Images\n", len(pimages.Images))
		for idx, img := range pimages.Images {
			fmt.Printf("Image %d - X: %.2f Y: %.2f, Width: %.2f, Height: %.2f\n",
				totalImages+idx+1, img.X, img.Y, img.Width, img.Height)
			fname := fmt.Sprintf("p%d_%d.jpg", i+1, idx)

			var gimg image.Image
			if gimg, err = img.Image.ToGoImage(); err != nil {
				return err
			}

			savename := outputPath + fname
			var fsave *os.File
			if fsave, err = os.Create(savename); err != nil {
				return
			}
			defer fsave.Close()

			opt := jpeg.Options{Quality: 100}
			if err = jpeg.Encode(fsave, gimg, &opt); err != nil {
				return err
			}
		}
		totalImages += len(pimages.Images)
	}
	fmt.Printf("Total: %d images\n", totalImages)

	return nil
}
