package lib

import (
	"fmt"
	"github.com/chai2010/webp"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type SourceImage struct {
	InputPath string
	Quality   float32
	Lossless  bool
}

func (s *SourceImage) ToWebp() (string, error) {
	file, err := os.Open(s.InputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var img image.Image

	switch strings.ToLower(filepath.Ext(s.InputPath)) {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		return "", fmt.Errorf("unsupported file type: %s", s.InputPath)
	}

	if err != nil {
		return "", err
	}

	outputPath := strings.TrimSuffix(s.InputPath, filepath.Ext(s.InputPath)) + ".webp"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	options := &webp.Options{Quality: s.Quality, Lossless: s.Lossless}
	err = webp.Encode(outputFile, img, options)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

func NewConversion(inputPath string, quality float32, lossless bool) *SourceImage {
	return &SourceImage{InputPath: inputPath, Quality: quality, Lossless: lossless}
}
