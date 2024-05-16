package main

import (
	"context"
	"fmt"

	"barnasipiczki.com/cli-db/db"
	"github.com/jackc/pgx/v5"
)

func main() {
	connString := "user=user password=password dbname=postgres host=192.168.1.163 port=5432 sslmode=disable"
	fmt.Println("Hello World!")
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Println("Error")
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	selectedEmp, err := queries.SelectEmployee(ctx, 1)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Printf("%+v\n", selectedEmp)

}
