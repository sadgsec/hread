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

type hreadapp struct {
	Dbpool *pgxpool.Pool
}

func dbHandler(dbpool *pgxpool.Pool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to ", r.RequestURI)
		if r.RequestURI == "/list" {
			err := list(w, r, dbpool)
			fmt.Println(err)
		} else if r.RequestURI == "/" {
			err := index(w, r, dbpool)
			fmt.Println(err)
		} else {
			shortname := r.RequestURI[1:]
			fmt.Println("Request for: ", shortname)
			err := boardView(w, r, dbpool, shortname)
			fmt.Println(err)
		}
	}
	return http.HandlerFunc(fn)
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
