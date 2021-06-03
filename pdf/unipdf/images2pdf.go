package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/unidoc/unipdf/creator"
)

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
		if filepath.Ext(imgPath) != ".jpg" && filepath.Ext(imgPath) != ".png" && filepath.Ext(imgPath) != ".jpeg" {
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

func imageFilesToPdf(outputPath string, files []string) (err error) {
	c := creator.New()
	for i := range files {
		if filepath.Ext(files[i]) != ".jpg" && filepath.Ext(files[i]) != ".png" && filepath.Ext(files[i]) != ".jpeg" {
			continue
		}

		fmt.Printf("Image: %s\n", files[i])

		var img *creator.Image
		if img, err = c.NewImageFromFile(files[i]); err != nil {
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
