package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	source := "/Users/fatihtastemur/GolandProjects/test-image.png"

	destination := "/Users/fatihtastemur/GolandProjects/go-workspace/tmp/"

	chunkDestination := fmt.Sprintf("%s%s/", destination, strconv.Itoa(int(time.Now().Unix())))
	_ = os.Mkdir(chunkDestination, os.ModePerm)

	sourceFile, errorOpen := os.Open(source)
	if errorOpen != nil {
		fmt.Println("Open/Exist File Error : ", errorOpen)
	}

	// Chunk Size : 5 MB
	chunkSize := int64(5 << 20)

	fileInfo, _ := sourceFile.Stat()
	chunkCount := int(fileInfo.Size() / chunkSize)
	if fileInfo.Size()%chunkSize != 0 {
		chunkCount++
	}

	splitFileToChunks(sourceFile, chunkSize, chunkDestination)
}

func splitFileToChunks(sourceFile *os.File, chunkSize int64, chunkDestination string) {
	buffer := make([]byte, chunkSize)

	defer func() {
		_ = sourceFile.Close()
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

		partNumber := index + 1
		chunkFileName := fmt.Sprintf("%schunk-%d", chunkDestination, partNumber)
		chunkFile, errorCreate := os.Create(chunkFileName)
		if errorCreate != nil {
			fmt.Println("Create File Error : ", errorCreate)
		}

		_, errorWrite := chunkFile.Write(buffer[:bytesRead])
		if errorWrite != nil {
			fmt.Println("Write File Error : ", errorWrite)
		}

		seeker = seeker + int64(bytesRead)
		index++
		fmt.Println("Read Bytes:", bytesRead, "Seeker:", seeker, "Index:", index)
		_ = chunkFile.Close()
	}
}
