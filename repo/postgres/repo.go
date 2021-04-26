package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/jackc/pgx/v4/stdlib"
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
		pgURL: url,
		client: pgclient,
	}
	return repo, nil
}

func (r *postgresRepo) RetrieveAll() ([]*pb.Book, error) {
	var book *pb.Book
	var bookList []*pb.Book
	rows, err := r.client.Query(bookRetrieveAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			book.ISBN,
			book.Title,
			book.Author,//TODO: fix
			book.Year,
			book.Edition,
			book.Publisher,//TODO: fix
			book.Pages,
			book.Category,
			book.PDF,
			book.Owned,
		)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, book)
	}
	if err != nil {
		return nil, err
	}
	return bookList, nil
}

func (r *postgresRepo) Retrieve(isbn string) (*pb.Book, error) {
	row := r.client.QueryRow(bookRetrievalQuery, isbn)
	var book *pb.Book
	err := row.Scan(
		book.ISBN,
		book.Title,
		book.Author,
		book.Year,
		book.Edition,
		book.Publisher,
		book.Pages,
		book.Category,
		book.PDF,
		book.Owned,
	)
	if err != nil {
		return nil, err
	}
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

