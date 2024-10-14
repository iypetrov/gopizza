package services

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/iypetrov/gopizza/configs"
)

type Image struct {
	s3Client *s3.Client
}

func NewImage(s3Client *s3.Client) Image {
	return Image{
		s3Client: s3Client,
	}
}

func (srv *Image) UploadImage(ctx context.Context, file io.Reader) (string, error) {
	id := uuid.New()
	_, err := srv.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(configs.Get().AWS.S3BucketName),
		Key:    aws.String(getImageKey(id)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/image/%s", configs.Get().GetBaseWebUrl(), getImageKey(id)), nil
}

func (srv *Image) GetImage(ctx context.Context, id uuid.UUID) (io.Reader, error) {
	resp, err := srv.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(configs.Get().AWS.S3BucketName),
		Key:    aws.String(getImageKey(id)),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Body, nil
}

func getImageKey(id uuid.UUID) string {
	return fmt.Sprintf("%s.png", id.String())
}
