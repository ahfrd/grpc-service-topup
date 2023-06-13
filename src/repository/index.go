package repository

import (
	"database/sql"

	"github.com/ahfrd/grpc/micro-topup/src/db"
)

type NullString struct {
	sql.NullString
}

type NullInt struct {
	sql.NullInt64
}

type TopUpRepository struct {
	db.Database
}
