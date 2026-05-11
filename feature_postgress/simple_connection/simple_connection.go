package simple_connection

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() {
	ctx := context.Background()

	connection, err := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := connection.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")
}
