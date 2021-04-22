package grpc

import (
	"log"
	"context"
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"gitlab.com/insanitywholesale/bookdir/repo/postgres"
	repointerface "gitlab.com/insanitywholesale/bookdir/repo/interface"
)

type Server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

var dbstore repointerface.BookDirRepo

func init() {
	pgURL := "postgres://tester:Apasswd@localhost?sslmode=disable"
	repo, err := postgres.NewPostgresRepo(pgURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	dbstore = repo
}

func (Server) GetAllBooks(_ context.Context, _ *pb.Empty) (*pb.BookList, error) {
	fmt.Println("getallbooks")
	allbooks, err := dbstore.RetrieveAll()
	if err != nil {
		return nil, err
	}
	return &pb.BookList{Books: allbooks}, nil
}

func (Server) AddBook(_ context.Context, book *pb.Book) (*pb.BookList, error) {
	fmt.Println("addbook")
	/*
	fmt.Println(book)
	testbooks = append(testbooks, book)
	*/
	return &pb.BookList{}, nil
}
