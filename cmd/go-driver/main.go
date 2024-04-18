package main

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/configs"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http/controller"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/find"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user/update"
	"github.com/andreluizmicro/go-driver-api/pkg/database"
	_ "github.com/go-sql-driver/mysql"
	"os"
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
	createUser := create.NewCreateUser(userRepository)
	findUser := find.NewFindUser(userRepository)
	updateUser := update.NewUpdateUser(userRepository)

	userController := controller.NewUserController(createUser, findUser, updateUser)

	http.InitRoutes(userController)
}

func init() {
	err := os.Setenv("TZ", "America/Sao_Paulo")
	if err != nil {
		return
	}
}
