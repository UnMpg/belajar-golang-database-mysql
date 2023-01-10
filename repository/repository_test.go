package repository

import (
	"context"
	"fmt"
	"golang-database-mysql/entity"
	"testing"

	golangdatabasemysql "golang-database-mysql"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabasemysql.GetConnection())
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "fendyasnanda@gmail.com",
		Comment: "komentar dari fendy",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabasemysql.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 31)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabasemysql.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
