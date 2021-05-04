package models

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type BookDirRepo interface {
/*
	//Publisher methods
	RetrieveAllPublishers() ([]*pb.Publisher, error)
	RetrieveBooksByPublisher(publisherId string) ([]*pb.Book, error)
	RetrievePublisherById() (*pb.Publisher, error)
*/

	//Author methods
	RetrieveAllAuthors() ([]*pb.Author, error)
	RetrieveBooksByAuthor(authorId string) ([]*pb.Book, error)
	RetrieveAuthorById(authorId string) (*pb.Author, error)

	//Book methods
	RetrieveAll() ([]*pb.Book, error)
	Retrieve(isbn string) (*pb.Book, error)
	Save(*pb.Book) error
}
