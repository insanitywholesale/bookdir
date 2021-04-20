package repointerface

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type BookDirRepo interface {
	Retrieve(isbn string) (*pb.Book, error)
	Save(*pb.Book) error
}
