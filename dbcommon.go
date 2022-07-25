package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BoardItem struct {
	Longname  string
	Shortname string
}

type BoardPost struct {
	Id      int
	Content string
}

func grabBoardPosts(dbpool *pgxpool.Pool, shortname string) ([]*BoardPost, error) {
	var posts []*BoardPost = *new([]*BoardPost)
	rows, err := dbpool.Query(context.Background(), "select post.id, post.content from post inner join board on post.boardid = board.id and board.shortname = $1", shortname)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var current = new(BoardPost)
		var id int
		var content string
		err = rows.Scan(&id, &content)
		if err != nil {
			return nil, err
		}
		current.Id = id
		current.Content = content
		posts = append(posts, current)
	}
	return posts, nil
}

func grabBoards(dbpool *pgxpool.Pool) ([]*BoardItem, error) {
	var boards []*BoardItem = *new([]*BoardItem)

	rows, err := dbpool.Query(context.Background(), "SELECT longname, shortname FROM Board")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var current *BoardItem = new(BoardItem)
		var longname, shortname string
		err = rows.Scan(&longname, &shortname)
		if err != nil {
			return nil, err
		}
		fmt.Println(longname)
		current.Longname = longname
		current.Shortname = shortname
		boards = append(boards, current)
	}
	return boards, nil
}
