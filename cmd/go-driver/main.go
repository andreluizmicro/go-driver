package main

import (
	"fmt"

	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/repository"
	"github.com/andreluizmicro/go-driver-api/internal/infrastructure/web/handlers"
	user "github.com/andreluizmicro/go-driver-api/internal/usecase/user/create"
	"github.com/andreluizmicro/go-driver-api/pkg/infrastructure/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewConnection()
	if err != nil {
		fmt.Println(err)
	}

	userRepository := repository.NewUserRepository(db)
	createUser := user.NewCreateUser(userRepository)
	userHandler := handlers.NewUserHandler(createUser)

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
