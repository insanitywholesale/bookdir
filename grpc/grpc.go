package grpc

import (
	"context"
	"errors"
	"gitlab.com/insanitywholesale/bookdir/models"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	//"gitlab.com/insanitywholesale/bookdir/repo/mock"
	//"gitlab.com/insanitywholesale/bookdir/repo/postgres"
	"gitlab.com/insanitywholesale/bookdir/repo/redis"
	"log"
	"os"
	"regexp"
)

type Server struct {
	// needed for forward compat for some reason
	pb.UnimplementedBookDirServer
}

var dbstore models.BookDirRepo

func init() {
/*
	if os.Getenv("PG_URL") != "" {
		pgURL := os.Getenv("PG_URL")
		repo, err := postgres.NewPostgresRepo(pgURL)
		if err != nil {
			log.Fatal(err)
		}
		dbstore = repo
		return
	}
*/
	if os.Getenv("REDIS_URL") != "" {
		redisURL := os.Getenv("REDIS_URL")
		repo, err := redis.NewRedisRepo(redisURL)
		if err != nil {
			log.Fatal(err)
			return
		}
		dbstore = repo
		return
	}
	//dbstore, _ = mock.NewMockRepo()
	//return
}

//TODO: implement
func (Server) GetBooksByAuthor(_ context.Context, author *pb.Author) (*pb.BookList, error) {
	return &pb.BookList{}, nil
}

//TODO: implement
func (Server) GetAllAuthors(_ context.Context, _ *pb.Empty) (*pb.AuthorList, error) {
	allauthors, err := dbstore.RetrieveAllAuthors()
	if err != nil {
		return nil, err
	}
	return &pb.AuthorList{Authors: allauthors}, nil
}

func (Server) GetBookByISBN(_ context.Context, pbisbn *pb.ISBN) (*pb.Book, error) {
	book, err := dbstore.Retrieve(pbisbn.ISBN)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (Server) GetAllBooks(_ context.Context, _ *pb.Empty) (*pb.BookList, error) {
	allbooks, err := dbstore.RetrieveAll()
	if err != nil {
		return nil, err
	}
	return &pb.BookList{Books: allbooks}, nil
}

func (Server) AddBook(_ context.Context, book *pb.Book) (*pb.Empty, error) {
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
			err := dbstore.Save(book)
			if err != nil {
				return nil, err
			}
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
			err := dbstore.Save(book)
			if err != nil {
				return nil, err
			}
			return &pb.Empty{}, nil
		} else {
			return nil, errors.New("invalid ISBN")
		}
	}
}
