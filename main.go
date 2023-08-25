package main

import (
	"context"
	"fmt"
	"github.com/cn-golang/minio/oss"
	"log"
)

func main() {
	//err := oss.UploadFile(context.Background(), "test", "file2.txt", "upload/file.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("upload OK")

	err := oss.Download(context.Background(), "test", "file2.txt", "download/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download OK")
}
