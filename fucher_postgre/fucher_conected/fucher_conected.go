package fucher_conected

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DataBase() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:Rjhjktdf_6126@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных успешно!")
}
