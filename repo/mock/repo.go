package mock

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"errors"
)

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

type bookrepo []*pb.Book

var testbooks bookrepo = []*pb.Book{
	testbook,
}

var testbooklist = &pb.BookList{
	Books: testbooks,
}

func NewMockRepo() (bookrepo, error) {
	return testbooks, nil
}

func (bookrepo) RetrieveAll() ([]*pb.Book, error) {
	return testbooks, nil
}

func (bookrepo) Retrieve(isbn string) (*pb.Book, error) {
	for _, b := range testbooks {
		if isbn == b.ISBN {
			return b, nil
		}
	}
	return nil, errors.New("no book with ISBN" + isbn + "was found")
}

func (bookrepo) Save(book *pb.Book) error {
	testbooks = append(testbooks, book)
	return nil
}
