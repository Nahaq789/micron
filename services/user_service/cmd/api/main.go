package main

import (
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"user_service/internal/shared/infrastructure/logger"
	"user_service/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	contextHandler := logger.NewContextHandler(handler)
	logger := slog.New(contextHandler)
	slog.SetDefault(logger)

	go func() {
		listen, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		proto.RegisterUserServiceServer(s)
		logger.Info("Server is running on :50051")
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r := gin.Default()
	if err := Router(r); err != nil {
		slog.Error("failed to start router", "error", err)
	}
	http.ListenAndServe(":8080", r)
}
