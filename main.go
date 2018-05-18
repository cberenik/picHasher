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
)

func main() {
	startTime := time.Now()

	args := os.Args[1:]
	var filePath string
	var wg sync.WaitGroup

	if len(args) == 0 {
		filePath = "./"
	} else {
		filePath = args[0]
	}
	files, err := ioutil.ReadDir(filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	var i int

	for _, file := range files {
		if isImage(file.Name()) {
			wg.Add(1)
			go rename(filePath+"/", file.Name(), &wg)
			i = i + 1
		}
	}

	wg.Wait()
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

func isImage(fileName string) bool {
	lowered := strings.ToLower(fileName)
	return strings.Contains(lowered, ".jpg") || strings.Contains(lowered, ".png") || strings.Contains(lowered, ".jpeg")
}
