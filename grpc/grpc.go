package grpc

import (
	"context"
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
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
		BooksPublished: 9000,
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

func (Server) GetAllBooks(_ context.Context, _ *pb.Empty) (*pb.BookList, error) {
	fmt.Println("getallbooks")
	return testbooklist, nil
}

func (Server) AddBook(_ context.Context, book *pb.Book) (*pb.BookList, error) {
	fmt.Println("addbook")
	testbooks = append(testbooks, book)
	return testbooklist, nil
}
