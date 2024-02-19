package config

import (
	"context"
	"fmt"
	"log"
	"music-app-backend/sqlc"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func InitDB() (DB *sqlc.Queries) {

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:music123@localhost:5434/music_app")
	if err != nil {
		log.Fatal(err)
	}
	DB = sqlc.New(conn)
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("Connected to database successfully.")
	return DB
}

func CloseDB() {
	ctx := context.Background()
	if conn != nil {
		if err := conn.Close(ctx); err != nil {
			log.Println("Failed to close database connection:", err)
		}
	}
}
