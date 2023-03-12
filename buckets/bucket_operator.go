package buckets

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func WriteToBucket(bucket_name string, file_path string) {

	ctx := context.Background()

	// Set up credentials to authenticate with Google Cloud

	client := clientConnection()

	bucketName := bucket_name
	filePath := file_path

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)

	}

	defer file.Close()

	// timestamp the file
	today := time.Now().Format("2006_01_02")

	// write the date part of the file path name
	filePath = fmt.Sprintf("%s_%s", today, filePath)

	// Create a new writer for the file in the bucket
	writer := client.Bucket(bucketName).Object(filePath).NewWriter(ctx)

	// Copy the content
	if _, err := io.Copy(writer, file); err != nil {
		log.Fatalf("Failed to write file to bucket: %v", err)
	}

	// Close the writer to flush the contents to the bucket
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to close writer %v", err)

	}

	log.Printf("File %q uploaded to bucket %q. \n", filePath, bucketName)

}

func ReadFromBucket(bucket_name string, file_path string) {

	ctx := context.Background()

	client := clientConnection()

	bucketName := bucket_name
	filePath := file_path

	obj := client.Bucket(bucketName).Object(filePath)
	reader, err := obj.NewReader(ctx)

	if err != nil {
		panic(err)
	}

	defer reader.Close()

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
