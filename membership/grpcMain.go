package main

import (
	"fmt"
	"os"

	"ecommerce-microservice/membership/db"

	"ecommerce-microservice/membership/src/presenter"
	"ecommerce-microservice/membership/src/query"
	"ecommerce-microservice/membership/src/repository"
	"ecommerce-microservice/membership/src/usecase"

	membershipGrpc "ecommerce-microservice/membership/grpc/servers"
)

//GrpcPortDefault default port for GRPC Server
const GrpcPortDefault = 3001

func grpcMain() {

	//init member handler
	memberQuery := query.NewMemberQueryInMemory(db.GetInMemoryDb())
	memberRepository := repository.NewMemberRepositoryInMemory(db.GetInMemoryDb())

	memberUseCase := usecase.NewMemberUseCase(memberRepository, memberQuery)

	memberGrpcHandler := presenter.NewGrpcHandler(memberUseCase)
	//end init member handler

	grpcServer, err := membershipGrpc.NewGrpcServer(memberGrpcHandler)

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
