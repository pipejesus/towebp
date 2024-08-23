package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func convertToWebP(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image
	switch strings.ToLower(filepath.Ext(inputPath)) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return fmt.Errorf("unsupported file type: %s", inputPath)
	}
	if err != nil {
		return err
	}

	outputPath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath)) + ".webp"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	options := &webp.Options{Quality: 100}
	err = webp.Encode(outputFile, img, options)
	if err != nil {
		return err
	}

	fmt.Printf("Converted %s to %s\n", inputPath, outputPath)
	return nil
}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				err := convertToWebP(file.Name())
				if err != nil {
					fmt.Println("Error converting file:", err)
				}
			}
		}
	}
}
