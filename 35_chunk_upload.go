package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func main() {

	source := "/Users/fatihtastemur/GolandProjects/test-image.png"
	destination := "/Users/fatihtastemur/GolandProjects/go-workspace/tmp/test-image.png"

	partialUpload(source, destination)
}

func partialUpload(source, destination string) {

	sourceFile, errorOpen := os.Open(source)
	if errorOpen != nil {
		fmt.Println("Open/Exist File Error : ", errorOpen)
	}

	// Chunk Size : 20 MB
	chunkSize := int64(20 << 20)
	buffer := make([]byte, chunkSize)

	fileInfo, _ := sourceFile.Stat()
	chunkCount := int(fileInfo.Size() / chunkSize)
	if fileInfo.Size()%chunkSize != 0 {
		chunkCount++
	}

	fmt.Println("File Size : ", fileInfo.Size())
	fmt.Println("Chunk Size : ", chunkSize)
	fmt.Println("Chunk Count : ", chunkCount)

	destinationFile, errorCreate := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if errorCreate != nil {
		fmt.Println("Create File Error : ", errorCreate)
	}

	defer func() {
		_ = sourceFile.Close()
		_ = destinationFile.Close()
	}()

	var seeker int64
	index := 0

	for {
		bytesRead, errorRead := sourceFile.Read(buffer)
		if errorRead != nil && errorRead != io.EOF {
			fmt.Println("Read File Error : ", errorRead)
		}

		if bytesRead == 0 {
			break
		}

		_, errorWrite := destinationFile.WriteAt(buffer[:bytesRead], seeker)
		if errorWrite != nil {
			fmt.Println("Write File Error : ", errorWrite)
		}

		seeker = seeker + int64(bytesRead)
		index++
		fmt.Println("Read Bytes:", bytesRead, "Seeker:", seeker, "Index:", index)
	}
}
