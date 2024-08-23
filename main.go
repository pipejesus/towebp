package main

import (
	"fmt"
	"towebp/lib"
)

func main() {
	imagesList, count, err := lib.NewImagesList(".").Scan()
	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}

	fmt.Println("Found", count, "images")

	for _, inputPath := range imagesList {
		conversion := lib.NewConversion(inputPath, 80, false)
		outputPath, err := conversion.ToWebp()
		if err != nil {
			fmt.Println("Error converting file:", err)
			continue
		}
		fmt.Println("Converted", inputPath, "to", outputPath)
	}
}
