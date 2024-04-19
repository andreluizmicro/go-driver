package main

import (
	"database/sql"
	"github.com/andreluizmicro/go-driver-api/configs"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/http/controller"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/folder"
	"github.com/andreluizmicro/go-driver-api/internal/usecase/user"
	"github.com/andreluizmicro/go-driver-api/pkg/database"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	cfg, err := configs.LoadConfig("../")
	if err != nil {
		panic(err)
	}

	InitInDebugMode(cfg, os.Args)

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
	folderRepository := repository.NewFolderRepository(db)

	createUser := user.NewCreateUser(userRepository)
	findUser := user.NewFindUser(userRepository)
	updateUser := user.NewUpdateUser(userRepository)
	deleteUser := user.NewDeleteUser(userRepository)
	listUser := user.NewListUser(userRepository)

	createFolder := folder.NewCreateFolder(folderRepository)
	updateFolder := folder.NewUpdateFolder(folderRepository)
	deleteFolder := folder.NewDeleteFolder(folderRepository)

	userController := controller.NewUserController(listUser, createUser, findUser, updateUser, deleteUser)
	folderController := controller.NewFolderController(createFolder, updateFolder, deleteFolder)

	http.InitRoutes(userController, folderController, cfg.WebServerPort)
}

func InitInDebugMode(cfg *configs.AppConfig, osArgs []string) {
	if len(osArgs) > 1 {
		cfg.DBHost = "127.0.0.1"
		cfg.WebServerPort = os.Args[1]
	}
}
