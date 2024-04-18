package main

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/configs"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http/controller"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user"
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
	findUser := user.NewFindUser(userRepository)
	updateUser := user.NewUpdateUser(userRepository)
	deleteUser := user.NewDeleteUser(userRepository)

	userController := controller.NewUserController(createUser, findUser, updateUser, deleteUser)

	http.InitRoutes(userController)
}
