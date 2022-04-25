package filesprotohdl

import (
	"bufio"
	"bytes"
	"entity/internal/core/ports"
	"entity/pb"
	"entity/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const maxfileSize = 1 << 23

type protoHandler struct {
	fs ports.FilesService
}

func NewProtoHandler(fs ports.FilesService) *protoHandler {
	return &protoHandler{fs: fs}
}

// UploadFile implements pb.FileServer
func (ph *protoHandler) UploadFile(stream pb.Files_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return errors.LogError(status.Errorf(codes.Internal, "cannot receive file info"))
	}
	fileType := req.GetInfo().GetFileType()
	log.Printf("receive an upload-file request with file type %s", fileType)

	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		err := errors.ContextError(stream.Context())
		if err != nil {
			return errors.LogError(err)
		}

		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return errors.LogError(status.Errorf(codes.Internal, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		fileSize += size
		if fileSize > maxfileSize {
			return errors.LogError(status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", fileSize, maxfileSize))
		}

		// write slowly
		// time.Sleep(time.Second)

		_, err = fileData.Write(chunk)
		if err != nil {
			return errors.LogError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	fileID, err := ph.fs.Upload(os.Getenv("FILE_PATH"), fileType, fileData)
	if err != nil {
		return errors.LogError(status.Errorf(codes.Internal, "cannot save file to the store: %v", err))
	}

	res := &pb.UploadFileResponse{
		Id:   fileID,
		Size: uint32(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return errors.LogError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved file with id: %s, size: %d", fileID, fileSize)
	return nil
}

// DownloadFile implements pb.FileServer
func (ph *protoHandler) DownloadFile(req *pb.DownloadFileRequest, stream pb.Files_DownloadFileServer) error {
	fileType := filepath.Ext(req.Filename)
	fileId := strings.TrimSuffix(filepath.Base(req.Filename), fileType)
	file, err := ph.fs.Download(os.Getenv("FILE_PATH"), fileId, fileType)

	resp := &pb.DownloadFileResponse{
		Data: &pb.DownloadFileResponse_Info{
			Info: &pb.FileInfo{
				FileType: fileType,
			},
		},
	}

	err = stream.Send(resp)
	if err != nil {
		return errors.LogError(status.Errorf(codes.Internal, "cannot send file info to client: ", err))
	}

	reader := bufio.NewReader(&file.FileData)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.LogError(status.Errorf(codes.Internal, "cannot read chunk to buffer: ", err))
		}

		resp := &pb.DownloadFileResponse{
			Data: &pb.DownloadFileResponse_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(resp)
		if err != nil {
			return errors.LogError(status.Errorf(codes.Internal, "cannot send chunk to server: ", err))
		}

	}
	return nil
}
