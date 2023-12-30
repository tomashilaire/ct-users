package filess3repo

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"root/internal/core/domain"
	"root/pkg/errors"
)

type s3Repository struct {
	sess *session.Session
}

func NewFilesRepository() *s3Repository {
	return &s3Repository{sess: ConnectAws()}
}

func (r *s3Repository) Save(f *domain.FileInfo) (*domain.FileInfo, error) {
	uploader := s3manager.NewUploader(r.sess)

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		//ACL:    aws.String("public-read"),
		Key:  aws.String(fmt.Sprintf("%s%s%s", f.Path, f.Id, f.Type)),
		Body: &f.FileData,
	})

	if err != nil {
		return &domain.FileInfo{}, errors.LogError(status.Errorf(codes.Internal, "cannot upload image to S3: %v, uploader: %s", err, up))
	}

	return f, nil
}

func (r *s3Repository) Load(p string) (*bytes.Buffer, error) {
	buffer := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(r.sess)
	numBytes, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(p),
	})
	if err != nil {
		return nil, err
	}

	if numBytes < 1 {
		return nil, errors.LogError(status.Errorf(codes.Internal, "Zero bytes were downloaded"))
	}
	fileData := bytes.NewBuffer(buffer.Bytes())
	return fileData, nil
}
