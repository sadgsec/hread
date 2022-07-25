package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postList struct {
	Posts []string
}

type boardPage struct {
	Posts []*BoardPost
}

type indexData struct {
	Boards []*BoardItem
}

func list(w http.ResponseWriter, req *http.Request, dbpool *pgxpool.Pool) (err error) {
	if req.Method == "POST" {
		content := req.FormValue("content")
		_, err := dbpool.Exec(context.Background(), "INSERT INTO Post (content) VALUES ($1)", content)
		if err != nil {
			return err
		}
	}

	var contentList []string
	rows, err := dbpool.Query(context.Background(), "select content from post")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var content string
		err = rows.Scan(&content)
		if err != nil {
			return err
		}
		contentList = append(contentList, content)
	}
	tmpl, err := template.ParseFiles("./templates/post_list.html")
	if err != nil {
		return err
	}
	tmpl.Execute(w, postList{contentList})

	return nil
}

func index(w http.ResponseWriter, req *http.Request, dbpool *pgxpool.Pool) (err error) {
	boards, err := grabBoards(dbpool)
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		return err
	}
	for i, b := range boards {
		fmt.Println(i, " ", b.Shortname)
	}
	err = tmpl.Execute(w, indexData{boards})
	if err != nil {
		return err
	}
	return nil
}

func boardView(w http.ResponseWriter, req *http.Request, dbpool *pgxpool.Pool, shortname string) (err error) {
	if req.Method == "POST" {
		content := req.FormValue("content")
		_, err := dbpool.Exec(context.Background(), "INSERT INTO Post (content) VALUES ($1)", content)
		if err != nil {
			return err
		}
	}

	contentList, err := grabBoardPosts(dbpool, shortname)
	if err != nil {
		return err
	}
	for _, c := range contentList {
		fmt.Println(c.Content)
	}
	tmpl, err := template.ParseFiles("./templates/board_view.html")
	if err != nil {
		return err
	}
	tmpl.Execute(w, boardPage{contentList})

	return nil
}
