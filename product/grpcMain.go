package main

import (
	"fmt"
	"os"

	productPresenterPackage "ecommerce-microservice/product/src/presenter"
	productQueryPackage "ecommerce-microservice/product/src/query"
	productUsecasePackage "ecommerce-microservice/product/src/usecase"

	inMemDB "ecommerce-microservice/product/db"

	productGrpc "ecommerce-microservice/product/grpc/server"
)

//GrpcPortDefault default port for GRPC Server
const GrpcPortDefault = 3002

// grpcMain
func grpcMain() {

	//init category grpc handler
	categoryQuery := productQueryPackage.NewCategoryQueryInMemory(inMemDB.GetCategoryInMemoryDb())
	categoryUseCase := productUsecasePackage.NewCategoryUseCase(categoryQuery)

	categoryGrpcHandler := productPresenterPackage.NewGrpcCategoryHandler(categoryUseCase)

	//init product grpc handler
	productQuery := productQueryPackage.NewProductQueryInMemory(inMemDB.GetProductInMemoryDb())
	productUseCase := productUsecasePackage.NewProductUseCase(productQuery)

	productGrpcHandler := productPresenterPackage.NewGrpcProductHandler(productUseCase)

	grpcServer, err := productGrpc.NewGrpcServer(categoryGrpcHandler, productGrpcHandler)

	if err != nil {
		fmt.Printf("Error create grpc server: %s", err.Error())
		os.Exit(1)
	}

	err = grpcServer.Serve(GrpcPortDefault)

	if err != nil {
		fmt.Printf("Error in Startup: %s", err.Error())
		os.Exit(1)
	}

}
