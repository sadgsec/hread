package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//CREATE TABLE testuno ( id SERIAL, title TEXT, body TEXT);
//export postgresql://user:secret@localhost

type testuno struct {
	id    int
	title string
	body  string
}

type hreadapp struct {
	Dbpool *pgxpool.Pool
}

func dbHandler(dbpool *pgxpool.Pool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to ", r.RequestURI)
		fmt.Fprintf(w, "henlo")
		if r.RequestURI == "/list" {
			list(w, r, dbpool)
		}
	}
	return http.HandlerFunc(fn)
}

func list(w http.ResponseWriter, req *http.Request, dbpool *pgxpool.Pool) {
	var id int
	var title, body string
	err := dbpool.QueryRow(context.Background(), "select * from testuno").Scan(&id, &title, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	w.Write([]byte(title))
	w.Write([]byte("<br>"))
	w.Write([]byte(body))
}

func main() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DBURL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	mux := http.NewServeMux()
	h := dbHandler(dbpool)

	mux.Handle("/", h)

	fmt.Println("testing testing, starting server on :8000")
	http.ListenAndServe(":8000", mux)
}
