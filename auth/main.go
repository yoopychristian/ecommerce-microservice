package main

import (
	"net/http"
	"os"
	"time"

	"ecommerce-microservice/auth/config"
	"ecommerce-microservice/auth/db"
	"ecommerce-microservice/auth/middleware"
	"ecommerce-microservice/auth/src/presenter"
	"ecommerce-microservice/auth/src/query"
	"ecommerce-microservice/auth/src/token"
	"ecommerce-microservice/auth/src/usecase"
)

func main() {

	//
	privateKey, err := config.InitPrivateKey()

	if err != nil {
		os.Exit(1)
	}

	// init auth handler

	identityQuery := query.NewIdentityQueryInMemory(db.GetInMemoryDb())
	accessTokenGenerator := token.NewJwtGenerator(privateKey, time.Minute*10)

	authUseCase := usecase.NewAuthUseCase(identityQuery, accessTokenGenerator)

	authHttpHandler := presenter.NewHttpHandler(authUseCase)
	// end init auth handler

	http.Handle("/api/auth", middleware.LogRequest(authHttpHandler.Auth()))
	http.ListenAndServe(":3000", nil)
}
