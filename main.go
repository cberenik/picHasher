package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // imported with underscore to register image type with image.Decode
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
	"time"

	namecalculator "github.com/cberenik/picHasher/name-calculator"
)

func main() {

	startTime := time.Now()

	args := os.Args[1:]
	var filePath string

	if len(args) == 0 {
		filePath = "./"
	} else {
		filePath = args[0]
	}

	fileInfos, err := ioutil.ReadDir(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed read files from directory\n%s", err))
		os.Exit(-1)
	}

	imageFiles := []os.FileInfo{}

	for _, file := range fileInfos {
		if isImage(file) {
			imageFiles = append(imageFiles, file)
		}
	}

	if len(imageFiles) > 0 {

		nameCalc := &namecalculator.BasicNameCalculator{}

		for _, fileInfo := range imageFiles {

			file, err := os.Open(filePath + fileInfo.Name())
			if err != nil {
				fmt.Println(fmt.Sprintf("opening file failed: %s", err))
				continue
			}
			defer file.Close()
			// img, format, err := image.Decode(file)
			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Println(fmt.Sprintf("decoding image failed: %s", err))
				continue
			}

			newName, err := nameCalc.Rename(img)
			if err != nil {
				fmt.Println(fmt.Sprintf("calculating new name failed: %s", err))
				continue
			}
			fmt.Println(newName)

			// err = os.Rename(filePath+fileInfo.Name(), filePath+newName+"."+strings.ToLower(format))

			// if err != nil {
			// 	fmt.Println(err.Error())
			// }
		}

	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Completed in %v seconds", elapsed.Seconds())
}

func isImage(file FileData) bool {
	if file.IsDir() {
		return false
	}
	lowered := strings.ToLower(file.Name())
	return strings.Contains(lowered, ".jpg") || strings.Contains(lowered, ".png") || strings.Contains(lowered, ".jpeg")
}

// FileData is a smaller interace than FileInfo so I don't have to fulfill that interface in tests
type FileData interface {
	Name() string
	IsDir() bool
}
