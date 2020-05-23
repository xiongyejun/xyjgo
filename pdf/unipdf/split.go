package main

import (
	"os"

	pdf "github.com/unidoc/unipdf/model"
)

func splitPdf(inputPath, outputPath, spages string) (err error) {
	pdfWriter := pdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return
	}
	defer f.Close()

	pdfReader, err := pdf.NewPdfReaderLazy(f)
	if err != nil {
		return
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return
	}

	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return
		}
	}

	// 解析页面index
	var pages []int
	if pages, err = getPages(spages); err != nil {
		return
	}

	for i := range pages {
		var page *pdf.PdfPage
		if page, err = pdfReader.GetPage(pages[i]); err != nil {
			return
		}

		if err = pdfWriter.AddPage(page); err != nil {
			return
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return
	}
	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return
	}

	return nil
}
