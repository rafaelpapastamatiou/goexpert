package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-15-s3-upload"
}

func main() {
	// Upload files sequentially
	// sequentialUpload()

	// Upload files concurrently - UNSAFE
	// concurrentUpload()

	// Upload files concurrently - SAFE with channel to limit concurrency
	concurrentUploadWithChannel()
}

// SEQUENTIAL upload without limiting the number of goroutines
func sequentialUpload() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	for {
		filename, err := dir.Readdirnames(1)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		if len(filename) == 0 {
			continue
		}

		fmt.Printf("Found file: %s\n", filename[0])
		uploadFile(filename[0])
	}
}

// UNSAFE concurrent upload without limiting the number of goroutines
func concurrentUpload() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	filenames, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	for _, filename := range filenames {
		wg.Add(1)

		go func(filename string) {
			defer wg.Done()
			uploadFile(filename)
		}(filename)
	}

	wg.Wait()
}

func concurrentUploadWithChannel() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	filenames, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 5)             // Limit to 5 concurrent uploads
	fileUploadErrors := make(chan string, 5) // Capture a max of 5 errors to retry concurrently

	go func() {
		for filename := range fileUploadErrors {
			fmt.Print("Retrying upload for file: ", filename)

			ch <- struct{}{}
			wg.Add(1)

			go func(filename string) {
				defer wg.Done()

				err := uploadFile(filename)
				if err != nil {
					fmt.Printf("Retry failed for %s: %v\n", filename, err)
				}

				<-ch
			}(filename)
		}
	}()

	for _, filename := range filenames {
		wg.Add(1)
		// Add empty struct to channel before starting upload
		// Channel will block if it is full, until a goroutine removes an empty struct
		ch <- struct{}{}

		go func(filename string) {
			defer wg.Done()

			err := uploadFile(filename)

			<-ch // Remove empty struct from channel after upload, to allow another goroutine to start

			// If there is an error, send the filename to the errors channel
			// so we can retry the upload one more time
			if err != nil {
				fileUploadErrors <- filename
			}
		}(filename)
	}

	wg.Wait()

	// Close channels after all uploads are finished
	close(fileUploadErrors)
	close(ch)
}

func uploadFile(filename string) error {
	fullPath := fmt.Sprintf("./tmp/%s", filename)

	fmt.Printf("Uploading %s to %s\n", filename, s3Bucket)

	file, err := os.Open(fullPath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return err
	}
	defer file.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", filename, err)
		return err
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, s3Bucket)

	return nil
}
