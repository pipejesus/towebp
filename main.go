package main

import (
	"runtime"
	"strconv"
	"sync"
	"towebp/lib"
)

func scheduleConversion(wg *sync.WaitGroup, inputPath string, quality float32, lossless bool) {
	defer wg.Done()
	conversion := lib.NewConversion(inputPath, quality, lossless)
	outputPath, err := conversion.ToWebp()
	if err != nil {
		lib.LogConversionError(inputPath, err.Error())
		return
	}

	lib.LogConverted(inputPath, outputPath)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	imagesList, count, err := lib.NewImagesList(".").Scan()
	if err != nil {
		lib.LogError("LogError scanning directory: " + err.Error())
	}

	lib.LogInfo("Found: " + strconv.Itoa(count) + " images")

	var wg sync.WaitGroup

	for _, inputPath := range imagesList {
		wg.Add(1)
		go scheduleConversion(&wg, inputPath, 80, false)
	}

	wg.Wait()
}
