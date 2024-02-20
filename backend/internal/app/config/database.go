package config

import (
	"context"
	"fmt"
	"log"
	"music-app-backend/sqlc"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func InitDB(dns string) (DB *sqlc.Queries) {

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dns)
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
