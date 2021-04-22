package grpc

import (
	"context"
	"errors"
	"fmt"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	repointerface "gitlab.com/insanitywholesale/bookdir/repo/interface"
	"gitlab.com/insanitywholesale/bookdir/repo/mock"
	"gitlab.com/insanitywholesale/bookdir/repo/postgres"
	"regexp"
	"log"
	"os"
)

type Server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

var dbstore repointerface.BookDirRepo

func init() {
	if os.Getenv("PG_URL") != "" {
		pgURL := "postgres://tester:Apasswd@localhost?sslmode=disable"
		repo, err := postgres.NewPostgresRepo(pgURL)
		if err != nil {
			log.Fatal(err)
		}
		dbstore = repo
		return
	}
	dbstore, _ = mock.NewMockRepo()
	return
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
	r := regexp.MustCompile("[^0-9]")
	isbn := r.ReplaceAllString(book.ISBN, "")
	isbnlen := len(isbn)
	if isbnlen != 13 && isbnlen != 10 {
		return nil, errors.New("invalid ISBN length")
	}

	//TODO: add check if book with isbn exists

	sum := 0
	if isbnlen == 10 {
		for i, v := range isbn {
			sum += int(v) * (10 - i)
		}
		if sum%11 == 0 {
			return &pb.Empty{}, nil
		} else {
			return nil, errors.New("invalid ISBN")
		}
	} else {
		for i, v := range isbn {
			tempadd := 0
			if i%2 == 0 {
				tempadd = 1
			} else {
				tempadd = 3
			}
			sum += int(v) * tempadd
		}
		if sum%10 == 0 {
			return &pb.Empty{}, nil
		} else {
			return nil, errors.New("invalid ISBN")
		}
	}
}
