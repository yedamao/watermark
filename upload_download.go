package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	blob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return blob, nil
}

func uploadImage(data []byte, format string) error {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(env.S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			env.S3_ACCESS_KEY_ID,
			env.S3_SECRET_ACCESS_KEY,
			"",
		),
	})

	if err != nil {
		return fmt.Errorf("new session: %v", err)
	}

	fileName := uuid.NewString() + "." + format
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(env.S3_BUCKET),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(data),
		ContentLength:        aws.Int64(int64(len(data))),
		ContentType:          aws.String(http.DetectContentType(data)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	return err
}
