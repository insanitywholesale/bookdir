package models

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type BookDirRepo interface {
	//Publisher methods
	RetrieveAllPublishers() ([]*pb.Publisher, error)
	RetrieveBooksByPublisher(publisherId uint32) ([]*pb.Book, error)
	RetrievePublisherById(publisherId uint32) (*pb.Publisher, error)

	//Author methods
	RetrieveAllAuthors() ([]*pb.Author, error)
	RetrieveBooksByAuthor(authorId uint32) ([]*pb.Book, error)
	RetrieveAuthorById(authorId uint32) (*pb.Author, error)

	//Book methods
	RetrieveAll() ([]*pb.Book, error)
	Retrieve(isbn string) (*pb.Book, error)
	Save(*pb.Book) error
}
