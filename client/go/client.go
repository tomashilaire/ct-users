package client

import (
	"bitbucket.org/agroproag/am_authentication/client/go/pb"
	"bitbucket.org/agroproag/am_authentication/client/go/pkg/errors"
	"bufio"
	"bytes"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type GrpcClient struct {
	fileService pb.FilesClient
	authService pb.AuthenticationClient
}

// NewGrpcClient returns a new grpc client
func NewGrpcClient(cc *grpc.ClientConn) *GrpcClient {
	fileService := pb.NewFilesClient(cc)
	authService := pb.NewAuthenticationClient(cc)
	return &GrpcClient{fileService: fileService, authService: authService}
}

// UploadFile calls upload file RPC
func (grpcClient *GrpcClient) UploadFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("cannot open file file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := grpcClient.fileService.UploadFile(ctx)
	if err != nil {
		log.Fatal("cannot upload file: ", err)
	}

	req := &pb.UploadFileRequest{
		Data: &pb.UploadFileRequest_Info{
			Info: &pb.FileInfo{
				FileType: filepath.Ext(filePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send file info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadFileRequest{
			Data: &pb.UploadFileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("file uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

// DownloadFile calls upload file RPC
func (grpcClient *GrpcClient) DownloadFile(filename string) *bytes.Buffer {

	req := &pb.DownloadFileRequest{Filename: filename}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := grpcClient.fileService.DownloadFile(ctx, req)
	if err != nil {
		log.Fatal("cannot download file: ", err)
	}

	resp, err := stream.Recv()
	if err != nil {
		log.Fatal("cannot receive file info", err)
	}
	fileType := resp.GetInfo().GetFileType()
	log.Printf("receive an upload-file request with file type %s", fileType)

	fileData := bytes.Buffer{}
	fileSize := 0
	for {
		err := errors.ContextError(stream.Context())
		if err != nil {
			log.Fatal("error: ", err)
		}

		log.Print("waiting to receive more data")

		resp, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			log.Fatal("cannot receive chunk data: %v", err)
		}

		chunk := resp.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		fileSize += size

		// write slowly
		// time.Sleep(time.Second)

		_, err = fileData.Write(chunk)
		if err != nil {
			log.Fatal(codes.Internal, "cannot write chunk data: %v", err)
		}
	}
	return &fileData
}
