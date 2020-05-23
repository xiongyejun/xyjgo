package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/unidoc/unipdf/creator"
	pdf "github.com/unidoc/unipdf/model"
)

// Rotate all pages by 90 degrees.
func rotate(inputPath, outputPath, degrees, spages string) (err error) {
	var angleDeg int64
	if angleDeg, err = strconv.ParseInt(degrees, 0, 64); err != nil {
		return
	}

	if angleDeg%90 != 0 {
		return errors.New("ERROR: Page rotation angle not a multiple of 90")
	}

	c := creator.New()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	// Try decrypting both with given password and an empty one if that fails.
	if isEncrypted {
		auth, err := pdfReader.Decrypt([]byte(""))
		if err != nil {
			return err
		}
		if !auth {
			return errors.New("Unable to decrypt pdf with empty pass")
		}
	}
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	// 解析页面index
	var pages []int
	if pages, err = getPages(spages); err != nil {
		return
	}
	// 记录需要调整的页面
	var bRotate []bool = make([]bool, numPages+1)
	for i := range pages {
		if pages[i] < 1 {
			err = errors.New("page base = 1")
			return
		}

		if pages[i] > numPages {
			err = errors.New(fmt.Sprintf("%d超过了页面总数%d", pages[i], numPages))
			return
		}
		bRotate[pages[i]] = true
	}

	for i := 0; i < numPages; i++ {
		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			return err
		}

		err = c.AddPage(page)
		if err != nil {
			return err
		}

		if bRotate[i+1] {
			// Do the rotation.
			var rotation int64
			if page.Rotate != nil {
				rotation = *(page.Rotate)
			}
			rotation += angleDeg // Rotate by angleDeg degrees.
			page.Rotate = &rotation
		}
	}

	err = c.WriteToFile(outputPath)
	return err
}
