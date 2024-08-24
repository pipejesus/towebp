package main

import (
	"flag"
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
	quality := flag.Float64("quality", 80, "quality of the webp image")
	lossless := flag.Bool("lossless", false, "lossless conversion")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())

	imagesList, count, err := lib.NewImagesList(".").Scan()
	if err != nil {
		lib.LogError("Error scanning directory: " + err.Error())
	}

	lib.LogInfo("Found: " + strconv.Itoa(count) + " images")

	var wg sync.WaitGroup

	for _, inputPath := range imagesList {
		wg.Add(1)
		go scheduleConversion(&wg, inputPath, float32(*quality), *lossless)
	}

	wg.Wait()
}
