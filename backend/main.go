package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	pb "github.com/jeffrey2212/downloader/backend/model_downloader" // Adjusted import path
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement modeldownloader.ModelDownloaderServer.
type server struct {
	pb.UnimplementedModelDownloaderServer
}

// DownloadModel implements modeldownloader.ModelDownloaderServer
func (s *server) DownloadModel(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	// Example: Construct the external API URL using the modelVersionId from the request
	url := fmt.Sprintf("https://civitai.com/api/download/models/%s", in.GetModelVersionId())

	// Make the HTTP GET request to the external RESTful API
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error calling external API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return nil, err
	}

	// Assuming the external API response is JSON that matches the structure of DownloadResponse
	var apiResponse pb.DownloadResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Printf("Error unmarshalling response body: %v", err)
		return nil, err
	}

	return &apiResponse, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterModelDownloaderServer(s, &server{})

	// Enable reflection
	reflection.Register(s)
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
    // Add logging around the external API call and its response
    


}
