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

var testbook = &pb.Book{
	ISBN:  "1234",
	Title: "hey",
	Author: &pb.Author{
		FirstName:    "me",
		MiddleName:   "notme",
		LastName:     "notyou",
		YearBorn:     1234,
		YearDied:     7890,
		BooksWritten: 2,
	},
	Year:    1999,
	Edition: 1,
	Publisher: &pb.Publisher{
		Name:           "urmom",
		YearStarted:    1237,
		YearEnded:      2077,
		BooksPublished: 9001,
	},
	Pages:    100,
	Category: "tech",
	PDF:      true,
	Owned:    false,
}

var testbooks = []*pb.Book{
	testbook,
}

var testbooklist = &pb.BookList{
	Books: testbooks,
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
	for _, b := range testbooks {
		fmt.Println(b)
	}
	return testbooklist, nil
}

func (Server) AddBook(_ context.Context, book *pb.Book) (*pb.BookList, error) {
	fmt.Println("addbook")
	fmt.Println(book)
	testbooks = append(testbooks, book)
	return testbooklist, nil
}
