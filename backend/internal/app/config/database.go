package config

import (
	"context"
	"fmt"
	"log"
	"music-app-backend/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func InitDB(dns string) (DB *sqlc.SQLStore) {

	ctx := context.Background()
	Conn, err := pgxpool.New(ctx, dns)
	if err != nil {
		log.Fatal(err)
	}
	DB = sqlc.NewStore(Conn)
	if err := Conn.Ping(ctx); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	fmt.Println("Connected to database successfully.")
	return DB
}

// func CloseDB() {
// 	ctx := context.Background()
// 	if Conn != nil {
// 		if err := Conn.Close(ctx); err != nil {
// 			log.Println("Failed to close database connection:", err)
// 		}
// 	}
// }
