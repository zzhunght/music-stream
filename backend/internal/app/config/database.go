package config

import (
	"context"
	"fmt"
	"log"
	"music-app-backend/sqlc"

	"github.com/jackc/pgx/v5"
)

var DB *sqlc.Queries
var conn *pgx.Conn

func InitDB() {

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
}

func CloseDB() {
	ctx := context.Background()
	if conn != nil {
		if err := conn.Close(ctx); err != nil {
			log.Println("Failed to close database connection:", err)
		}
	}
}
