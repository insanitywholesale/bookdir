package mock

import (
	"errors"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

var testbook = &pb.Book{
	ISBN:  "9605122839",
	Title: "hey",
	Author: &pb.Author{
		AuthorID:     1,
		FirstName:    "anguish",
		MiddleName:   "none",
		LastName:     "mental",
		YearBorn:     404,
		YearDied:     201,
		BooksWritten: 5,
	},
	Year:    1999,
	Edition: 1,
	Publisher: &pb.Publisher{
		PublisherID:    1,
		Name:           "alsome",
		YearStarted:    101,
		YearEnded:      2177,
		BooksPublished: 9001,
	},
	Pages:    169,
	Category: "blog",
	PDF:      true,
	Owned:    true,
}

var testbook2 = &pb.Book{
	ISBN:  "0605122839",
	Title: "xey",
	Author: &pb.Author{
		AuthorID:     7,
		FirstName:    "storm",
		MiddleName:   "u",
		LastName:     "cloud",
		YearBorn:     1955,
		YearDied:     2003,
		BooksWritten: 9,
	},
	Year:    1993,
	Edition: 3,
	Publisher: &pb.Publisher{
		PublisherID:    2,
		Name:           "poobleesh",
		YearStarted:    1820,
		YearEnded:      2177,
		BooksPublished: 69,
	},
	Pages:    420,
	Category: "test",
	PDF:      false,
	Owned:    true,
}

var authorCount uint32 = 1
var publisherCount uint32 = 1

type bookrepo []*pb.Book

var testbooks bookrepo = []*pb.Book{
	testbook,
	testbook2,
}

var testbooklist = &pb.BookList{
	Books: testbooks,
}

func NewMockRepo() (bookrepo, error) {
	return testbooks, nil
}

func (bookrepo) RetrieveAllAuthors() ([]*pb.Author, error) {
	authorlist := []*pb.Author{}
	for _, b := range testbooks {
		authorlist = append(authorlist, b.Author)
	}
	return authorlist, nil
}

func (bookrepo) RetrieveBooksByAuthor(authorId uint32) ([]*pb.Book, error) {
	for _, b := range testbooks {
		if authorId == b.Author.AuthorID {
			testbooks = append(testbooks, b)
		}
	}
	return testbooks, nil
}

func (bookrepo) RetrieveAuthorById(authorId uint32) (*pb.Author, error) {
	for _, b := range testbooks {
		if authorId == b.Author.AuthorID {
			return b.Author, nil
		}
	}
	return nil, errors.New("author not found")
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
	authorCount = authorCount + 1
	book.Author.AuthorID = authorCount
	publisherCount = publisherCount + 1
	book.Publisher.PublisherID = publisherCount
	testbooks = append(testbooks, book)
	return nil
}
