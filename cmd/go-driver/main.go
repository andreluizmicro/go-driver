package main

import (
	"database/sql"
	"net/http"

	"github.com/andreluizmicro/go-driver-api/configs"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/web/handlers"
	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
	"github.com/andreluizmicro/go-driver-api/pkg/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	db, err := database.NewConnection(cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	userRepository := repository.NewUserRepository(db)
	createUser := user.NewCreateUser(userRepository)
	findUser := find.NewFindUser(userRepository)
	userHandler := handlers.NewUserHandler(createUser, findUser)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users/{id}", userHandler.FindById)

	http.ListenAndServe(":"+cfg.WebServerPort, mux)
}
