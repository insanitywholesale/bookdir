package models

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type BookDirRepo interface {
	RetrieveAll() ([]*pb.Book, error)
	Retrieve(isbn string) (*pb.Book, error)
	Save(*pb.Book) error
}
