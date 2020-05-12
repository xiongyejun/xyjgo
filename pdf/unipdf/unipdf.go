package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/unidoc/unipdf/creator"
	pdf "github.com/unidoc/unipdf/model"
)

// TODO 记得修改这个
/*
func (k *LicenseKey) IsLicensed() bool {
	// return k.Tier != LicenseTierUnlicensed
	return true
}
*/

func main() {
	if len(os.Args) == 1 {
		printHelp()
		return
	}

	var err error
	switch os.Args[1] {
	case "i":
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = imagesToPdf(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}

	case "s":
		if len(os.Args) != 6 {
			printHelp()
		}
		var pageFrom, pageTo int
		if pageFrom, err = strconv.Atoi(os.Args[4]); err != nil {
			fmt.Println(err)
			return
		}
		if pageTo, err = strconv.Atoi(os.Args[5]); err != nil {
			fmt.Println(err)
			return
		}
		if err = splitPdf(os.Args[2], os.Args[3], pageFrom, pageTo); err != nil {
			fmt.Println(err)
			return
		}

	case "c":
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = compress(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}

	case "e":
		if len(os.Args) != 4 {
			printHelp()
		}
		if err = extractImages(os.Args[2], os.Args[3]); err != nil {
			fmt.Println(err)
			return
		}
	default:
		printHelp()
	}
}

func imagesToPdf(imagesFolder, outputPath string) (err error) {
	var entrys []os.FileInfo
	if entrys, err = ioutil.ReadDir(imagesFolder); err != nil {
		return
	}

	c := creator.New()
	for i := range entrys {
		if entrys[i].IsDir() {
			continue
		}

		imgPath := entrys[i].Name()
		if filepath.Ext(imgPath) != ".jpg" && filepath.Ext(imgPath) != ".png" {
			continue
		}

		fmt.Printf("Image: %s\n", imgPath)

		var img *creator.Image
		if img, err = c.NewImageFromFile(imgPath); err != nil {
			return
		}
		img.ScaleToWidth(612.0)

		// Use page width of 612 points, and calculate the height proportionally based on the image.
		// Standard PPI is 72 points per inch, thus a width of 8.5"
		height := 612.0 * img.Height() / img.Width()
		c.SetPageSize(creator.PageSize{612, height})
		c.NewPage()
		img.SetPos(0, 0)
		if err = c.Draw(img); err != nil {
			return
		}
	}

	return c.WriteToFile(outputPath)
}

func splitPdf(inputPath string, outputPath string, pageFrom int, pageTo int) (err error) {
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

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return
	}

	if numPages < pageTo {
		return
	}

	for i := pageFrom; i <= pageTo; i++ {
		pageNum := i

		var page *pdf.PdfPage
		if page, err = pdfReader.GetPage(pageNum); err != nil {
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

func printHelp() {
	fmt.Println(`
 i <imagesFolder> <outputPath> -- images to pdf 图片创建pdf
 s <inputPath> <outputPath> <pageFrom>, <pageTo> -- 拆分pdf
 c <inputPath> <outputPath> -- Compress and optimize PDF 压缩pdf
 e <inputPath> <outputPath> -- Extract images 提取图片
	`)
}
