package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "net/http"

    "google.golang.org/grpc"
    pb "github.com/jeffrey2212/downloader/backend/model_downloader"
)

type server struct {
    pb.UnimplementedModelDownloaderServer
}

func (s *server) DownloadModel(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
    url := fmt.Sprintf("https://civitai.com/api/download/models/%s", in.GetModelVersionId())
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return &pb.DownloadResponse{Data: string(body)}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterModelDownloaderServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
