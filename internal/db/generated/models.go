// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID        int32
	UserID    int32
	Content   string
	CreatedAt pgtype.Timestamp
}

type User struct {
	ID        int32
	Name      string
	Email     string
	CreatedAt pgtype.Timestamp
	Password  pgtype.Text
}
