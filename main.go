package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"namecalculator"
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

	images := []string{}

	for _, file := range fileInfos {
		if isImage(file) {
			images = append(images, file.Name())
		}
	}

	if len(images) > 0 {

		for file := range fileInfos {
			nameCalc := &namecalculator.BasicNameCalculator{}

		}

		// TODO: send images to channels to calculate dominant image colors and rename them
		// Don't load all images at the same time though, don't have infinite RAM

	}

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Completed in %v seconds", elapsed.Seconds())
}

func rename(path string, oldName string, wg *sync.WaitGroup) {
	defer wg.Done()
	h := md5.New()

	parts := strings.Split(oldName, ".")
	name := strings.Join(parts[:len(parts)-1], ".")
	extension := parts[len(parts)-1]

	imgBytes, err := ioutil.ReadFile(path + oldName)
	if err != nil {
		fmt.Printf("%s%s.%s not found\n", path, name, extension)
		return
	}

	io.WriteString(h, string(imgBytes))
	hashString := hex.EncodeToString(h.Sum(nil))
	err = os.Rename(path+oldName, path+hashString+"."+strings.ToLower(extension))

	if err != nil {
		fmt.Println(err.Error())
	}
}

func isImage(file os.FileInfo) bool {
	if file.IsDir() {
		return false
	}
	lowered := strings.ToLower(file.Name())
	return strings.Contains(lowered, ".jpg") || strings.Contains(lowered, ".png") || strings.Contains(lowered, ".jpeg")
}
