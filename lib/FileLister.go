package lib

import (
	"os"
	"path/filepath"
	"strings"
)

var supportedImageExtensions = map[string]string{
	".jpg":  "jpg",
	".jpeg": "jpeg",
	".png":  "png",
}

type ImagesList struct {
	DirPath string
}

func (d *ImagesList) Scan() ([]string, int, error) {
	var imagePaths []string = make([]string, 0)
	var count int = 0

	files, err := os.ReadDir(d.DirPath)
	if err != nil {
		return nil, 0, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		extension := strings.ToLower(filepath.Ext(file.Name()))

		if _, ok := supportedImageExtensions[extension]; !ok {
			continue
		}

		imagePaths = append(imagePaths, file.Name())
		count++
	}

	return imagePaths, count, nil
}

func NewImagesList(dirPath string) *ImagesList {
	return &ImagesList{DirPath: dirPath}
}
