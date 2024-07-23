package main

import (
	"github.com/liuhuankun/category/handler"
	pb "github.com/liuhuankun/category/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "category"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterCategoryHandler(srv.Server(), new(handler.Category)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
