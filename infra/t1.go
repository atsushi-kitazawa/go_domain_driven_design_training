package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type t1 struct {
    id int
    name string
}

func Select() {
    db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=mysecretpassword dbname=testdb sslmode=disable")
    if err != nil {
	panic(err)
    }
    defer db.Close()

    rows, err := db.Query("select id, name from t1")
    if err != nil {
	panic(err)
    }

    for rows.Next() {
	var t t1
	rows.Scan(&t.id, &t.name)
	fmt.Println(t.id)
	fmt.Println(t.name)
    }
}
