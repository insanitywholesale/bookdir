package grpc

import (
	_ "log"
	"context"
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	_ "gitlab.com/insanitywholesale/bookdir/repo/postgres"
	"gitlab.com/insanitywholesale/bookdir/repo/mock"
	repointerface "gitlab.com/insanitywholesale/bookdir/repo/interface"
)

type Server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

var dbstore repointerface.BookDirRepo

func init() {
	//pgURL := "postgres://tester:Apasswd@localhost?sslmode=disable"
	//repo, err := postgres.NewPostgresRepo(pgURL)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	repo, _ := mock.NewMockRepo()
	dbstore = repo
}

func (Server) GetBookByISBN(_ context.Context, pbisbn *pb.ISBN) (*pb.Book, error) {
	fmt.Println("getbookbyisbn")
	book, err := dbstore.Retrieve(pbisbn.ISBN)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (Server) GetAllBooks(_ context.Context, _ *pb.Empty) (*pb.BookList, error) {
	fmt.Println("getallbooks")
	allbooks, err := dbstore.RetrieveAll()
	if err != nil {
		return nil, err
	}
	return &pb.BookList{Books: allbooks}, nil
}

func (Server) AddBook(_ context.Context, book *pb.Book) (*pb.Empty, error) {
	fmt.Println("addbook")
	return &pb.Empty{}, nil
}
