package main

import (
	"flag"
	"fmt"
	"github.com/scusi/bytesize"
	"log"
	"strconv"
	"os"
)

var fileName string
var fileSize string

func init() {
	flag.StringVar(&fileName, "f", "", "file to show filesize for")
	flag.StringVar(&fileSize, "s", "", "size in byte to convert")
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	if fileName != "" {
		// open file and output size in human readable format
		file, err := os.Open(fileName)
		checkFatal(err)
		defer file.Close()
		fileInfo, err := file.Stat()
		checkFatal(err)
		humanSize := bytesize.ByteSize(int64(fileInfo.Size()))
		fmt.Printf("'%s' %s\n", fileName, humanSize)
	}
	if fileSize != "" {
		fsInt, err := strconv.Atoi(fileSize)
		checkFatal(err)
		humanSize := bytesize.ByteSize(int64(fsInt))
		fmt.Printf("%d = %s\n", fsInt, humanSize)
	}
}
