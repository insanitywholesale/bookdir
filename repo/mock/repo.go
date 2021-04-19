package mock

import (
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
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

var testbooks = []*pb.Book{
	testbook,
}

var testbooklist = &pb.BookList{
	Books: testbooks,
}

//TODO: Implement same methods as postgres
