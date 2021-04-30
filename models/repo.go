package models

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type BookDirRepo interface {
	RetrieveAllAuthors() ([]*pb.Author, error)
	RetrieveBooksByAuthor(authorId) ([]*pb.Book, error)
	RetrieveAuthorById() (*pb.Author, error)
	RetrieveAll() ([]*pb.Book, error)
	Retrieve(isbn string) (*pb.Book, error)
	Save(*pb.Book) error
}
