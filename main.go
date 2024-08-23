package main

import (
	"strconv"
	"towebp/lib"
)

func main() {
	imagesList, count, err := lib.NewImagesList(".").Scan()
	if err != nil {
		lib.LogError("LogError scanning directory: " + err.Error())
	}

	lib.LogInfo("Found: " + strconv.Itoa(count) + " images")

	for _, inputPath := range imagesList {
		conversion := lib.NewConversion(inputPath, 80, false)
		outputPath, err := conversion.ToWebp()
		if err != nil {
			lib.LogConversionError(inputPath, err.Error())
			continue
		}

		lib.LogConverted(inputPath, outputPath)
	}
}
