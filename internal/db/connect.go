package initializers

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"

	"github.com/madhav23bansal/sqlc-poc/internal/config"
	db "github.com/madhav23bansal/sqlc-poc/internal/db/generated"
)

var database *db.Queries

func ConnectToDB() {
	var once sync.Once
	once.Do(func() {
		conn, err := pgx.Connect(context.Background(), config.GetEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/postgres"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Connected to database")

		q := db.New(conn)

		if q != nil {
			database = q
		}
	})
}

func GetDB() *db.Queries {
	if database == nil {
		ConnectToDB()
	}
	return database
}
