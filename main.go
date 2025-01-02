package main

import (
	"context"
	"fmt"

	"github.com/madhav23bansal/sqlc-poc/internal/config"
	initializers "github.com/madhav23bansal/sqlc-poc/internal/db"
)

func init() {
	fmt.Println("Starting the init file")
	config.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Starting the main file")
	dbRef := initializers.GetDB()

	// user, err := dbRef.CreateUser(context.Background(), db.CreateUserParams{
	// 	Name:  "Madhav Bansal",
	// 	Email: "madhavb.dev@gmail.com",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(`User created with ID:`, user.ID, " ", user.Name)

	users, err := dbRef.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)

}
