package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
	"log"
)

type postgresRepo struct {
	client *sql.DB
	pgURL  string
}

func newPostgresClient(url string) (*sql.DB, error) {
	client, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = client.Ping()
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createAuthorTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createPublisherTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createBookTableQuery)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewPostgresRepo(url string) (*postgresRepo, error) {
	pgclient, err := newPostgresClient(url)
	if err != nil {
		return nil, err
	}
	repo := &postgresRepo{
		pgURL:  url,
		client: pgclient,
	}
	return repo, nil
}

func (r *postgresRepo) RetrieveAllPublishers() ([]*pb.Publisher, error) {
	var publisher = &pb.Publisher{}
	var publisherList []*pb.Publisher

	rows, err := r.client.Query(`SELECT * FROM Publisher`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&publisher.PublisherID,
			&publisher.Name,
			&publisher.YearStarted,
			&publisher.YearEnded,
			&publisher.BooksPublished,
		)
		if err != nil {
			return nil, err
		}
		publisherList = append(publisherList, publisher)
	}
	return publisherList, nil

}

func (r *postgresRepo) RetrieveBooksByPublisher(publisherId uint32) ([]*pb.Book, error) {
	var author = &pb.Author{}
	var publisher = &pb.Publisher{}
	var authID int
	var pubID int

	var book = &pb.Book{}
	var bookList []*pb.Book

	rows, err := r.client.Query(`SELECT * FROM Book WHERE PublisherID=$1`, publisherId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&book.ISBN,
			&book.Title,
			&author.AuthorID,
			&book.Year,
			&book.Edition,
			&publisher.PublisherID,
			&book.Pages,
			&book.Category,
			&book.PDF,
			&book.Owned,
		)
		if err != nil {
			return nil, err
		}

		isbn := book.ISBN

		rowPublisherID := r.client.QueryRow(giveBookGetPublisherID, isbn)
		err := rowPublisherID.Scan(&pubID)
		if err != nil {
			return nil, err
		}
		rowAuthorID := r.client.QueryRow(giveBookGetAuthorID, isbn)
		err = rowAuthorID.Scan(&authID)
		if err != nil {
			return nil, err
		}

		rowAuthor := r.client.QueryRow(authorRetrievalQuery, authID)
		err = rowAuthor.Scan(
			&author.AuthorID,
			&author.FirstName,
			&author.MiddleName,
			&author.LastName,
			&author.YearBorn,
			&author.YearDied,
			&author.BooksWritten,
		)
		if err != nil {
			return nil, err
		}
		rowPublisher := r.client.QueryRow(publisherRetrievalQuery, pubID)
		err = rowPublisher.Scan(
			&publisher.PublisherID,
			&publisher.Name,
			&publisher.YearStarted,
			&publisher.YearEnded,
			&publisher.BooksPublished,
		)
		if err != nil {
			return nil, err
		}

		book.Author = author
		book.Publisher = publisher

		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (r *postgresRepo) RetrievePublisherById(publisherId uint32) (*pb.Publisher, error) {
	var publisher = &pb.Publisher{}
	rowPublisher, err := r.client.Query(`SELECT * FROM Publisher WHERE PublisherID=$1`, publisherId)
	if err != nil {
		return nil, err
	}
	err = rowPublisher.Scan(
		&publisher.PublisherID,
		&publisher.Name,
		&publisher.YearStarted,
		&publisher.YearEnded,
		&publisher.BooksPublished,
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *postgresRepo) RetrieveBooksByAuthor(authorId uint32) ([]*pb.Book, error) {
	var author = &pb.Author{}
	var publisher = &pb.Publisher{}
	var authID int
	var pubID int

	var book = &pb.Book{}
	var bookList []*pb.Book

	rows, err := r.client.Query(`SELECT * FROM Book WHERE AuthorID=$1`, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&book.ISBN,
			&book.Title,
			&author.AuthorID,
			&book.Year,
			&book.Edition,
			&publisher.PublisherID,
			&book.Pages,
			&book.Category,
			&book.PDF,
			&book.Owned,
		)
		if err != nil {
			return nil, err
		}

		isbn := book.ISBN

		rowPublisherID := r.client.QueryRow(giveBookGetPublisherID, isbn)
		err := rowPublisherID.Scan(&pubID)
		if err != nil {
			return nil, err
		}
		rowAuthorID := r.client.QueryRow(giveBookGetAuthorID, isbn)
		err = rowAuthorID.Scan(&authID)
		if err != nil {
			return nil, err
		}

		rowAuthor := r.client.QueryRow(authorRetrievalQuery, authID)
		err = rowAuthor.Scan(
			&author.AuthorID,
			&author.FirstName,
			&author.MiddleName,
			&author.LastName,
			&author.YearBorn,
			&author.YearDied,
			&author.BooksWritten,
		)
		if err != nil {
			return nil, err
		}
		rowPublisher := r.client.QueryRow(publisherRetrievalQuery, pubID)
		err = rowPublisher.Scan(
			&publisher.PublisherID,
			&publisher.Name,
			&publisher.YearStarted,
			&publisher.YearEnded,
			&publisher.BooksPublished,
		)
		if err != nil {
			return nil, err
		}

		book.Author = author
		book.Publisher = publisher

		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (r *postgresRepo) RetrieveAuthorById(authorId uint32) (*pb.Author, error) {
	var author = &pb.Author{}
	rowAuthor, err := r.client.Query(`SELECT * FROM Author WHERE AuthorID=$1`, authorId)
	if err != nil {
		return nil, err
	}
	err = rowAuthor.Scan(
		&author.AuthorID,
		&author.FirstName,
		&author.MiddleName,
		&author.LastName,
		&author.YearBorn,
		&author.YearDied,
		&author.BooksWritten,
	)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (r *postgresRepo) RetrieveAllAuthors() ([]*pb.Author, error) {
	var author = &pb.Author{}
	var authorList []*pb.Author

	rows, err := r.client.Query(`SELECT * FROM Author`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&author.AuthorID,
			&author.FirstName,
			&author.MiddleName,
			&author.LastName,
			&author.YearBorn,
			&author.YearDied,
			&author.BooksWritten,
		)
		if err != nil {
			return nil, err
		}
		authorList = append(authorList, author)
	}
	return authorList, nil
}

func (r *postgresRepo) RetrieveAll() ([]*pb.Book, error) {
	var author = &pb.Author{}
	var publisher = &pb.Publisher{}
	var authID int
	var pubID int

	var book = &pb.Book{}
	var bookList []*pb.Book

	rows, err := r.client.Query(bookRetrieveAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&book.ISBN,
			&book.Title,
			&author.AuthorID,
			&book.Year,
			&book.Edition,
			&publisher.PublisherID,
			&book.Pages,
			&book.Category,
			&book.PDF,
			&book.Owned,
		)
		if err != nil {
			return nil, err
		}

		isbn := book.ISBN

		rowPublisherID := r.client.QueryRow(giveBookGetPublisherID, isbn)
		err := rowPublisherID.Scan(&pubID)
		if err != nil {
			return nil, err
		}
		rowAuthorID := r.client.QueryRow(giveBookGetAuthorID, isbn)
		err = rowAuthorID.Scan(&authID)
		if err != nil {
			return nil, err
		}

		rowAuthor := r.client.QueryRow(authorRetrievalQuery, authID)
		err = rowAuthor.Scan(
			&author.AuthorID,
			&author.FirstName,
			&author.MiddleName,
			&author.LastName,
			&author.YearBorn,
			&author.YearDied,
			&author.BooksWritten,
		)
		if err != nil {
			return nil, err
		}
		rowPublisher := r.client.QueryRow(publisherRetrievalQuery, pubID)
		err = rowPublisher.Scan(
			&publisher.PublisherID,
			&publisher.Name,
			&publisher.YearStarted,
			&publisher.YearEnded,
			&publisher.BooksPublished,
		)
		if err != nil {
			return nil, err
		}

		book.Author = author
		book.Publisher = publisher

		bookList = append(bookList, book)
	}
	return bookList, nil
}

func (r *postgresRepo) Retrieve(isbn string) (*pb.Book, error) {
	var author = &pb.Author{}
	var publisher = &pb.Publisher{}
	var authID int
	var pubID int

	rowPublisherID := r.client.QueryRow(giveBookGetPublisherID, isbn)
	err := rowPublisherID.Scan(&pubID)
	if err != nil {
		return nil, err
	}
	rowAuthorID := r.client.QueryRow(giveBookGetAuthorID, isbn)
	err = rowAuthorID.Scan(&authID)
	if err != nil {
		return nil, err
	}

	rowAuthor := r.client.QueryRow(authorRetrievalQuery, authID)
	err = rowAuthor.Scan(
		&author.AuthorID,
		&author.FirstName,
		&author.MiddleName,
		&author.LastName,
		&author.YearBorn,
		&author.YearDied,
		&author.BooksWritten,
	)
	if err != nil {
		return nil, err
	}
	rowPublisher := r.client.QueryRow(publisherRetrievalQuery, pubID)
	err = rowPublisher.Scan(
		&publisher.PublisherID,
		&publisher.Name,
		&publisher.YearStarted,
		&publisher.YearEnded,
		&publisher.BooksPublished,
	)
	if err != nil {
		return nil, err
	}

	var book = &pb.Book{Author: author, Publisher: publisher}
	row := r.client.QueryRow(bookRetrievalQuery, isbn)
	err = row.Scan(
		&book.ISBN,
		&book.Title,
		&author.AuthorID,
		&book.Year,
		&book.Edition,
		&publisher.PublisherID,
		&book.Pages,
		&book.Category,
		&book.PDF,
		&book.Owned,
	)
	if err != nil {
		return nil, err
	}
	log.Println("retrieved book:", book)
	return book, nil
}

func (r *postgresRepo) Save(book *pb.Book) error {
	log.Println("book", book)

	var authorId int
	errA := r.client.QueryRow(authorInsertQuery,
		book.Author.FirstName,
		book.Author.MiddleName,
		book.Author.LastName,
		book.Author.YearBorn,
		book.Author.YearDied,
		book.Author.BooksWritten,
	).Scan(&authorId)
	if errA != nil {
		log.Println("A1err", errA)
		return errA
	}

	var publisherId int
	errP := r.client.QueryRow(publisherInsertQuery,
		book.Publisher.Name,
		book.Publisher.YearStarted,
		book.Publisher.YearEnded,
		book.Publisher.BooksPublished,
	).Scan(&publisherId)
	if errP != nil {
		log.Println("P1err", errP)
		return errP
	}

	resBook, err := r.client.Exec(bookInsertQuery,
		book.ISBN,
		book.Title,
		authorId,
		book.Year,
		book.Edition,
		publisherId,
		book.Pages,
		book.Category,
		book.PDF,
		book.Owned,
	)
	log.Println(resBook)
	if err != nil {
		log.Println("Berr", err)
		return err
	}
	return nil
}
