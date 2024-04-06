package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andreluizmicro/go-driver-api/configs"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/web/handlers"
	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
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
	userHandler := handlers.NewUserHandler(createUser)

	http.HandleFunc("/user", userHandler.Create)
	http.ListenAndServe(":"+cfg.WebServerPort, nil)

	users, err := userHandler.UseCase.UserRepository.FindAll(&repository.Filters{})
	if err != nil {
		panic(err)
	}

	for key, user := range users {
		fmt.Printf("%d - %s\n", key+1, user.Name)
	}
}

func init() {
	// r.Post("/user", userHandler.Create)
}
