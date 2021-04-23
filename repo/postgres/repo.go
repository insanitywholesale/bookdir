package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
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
	_, err = client.Exec(createPublisherTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = client.Exec(createAuthorTableQuery)
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
	_, err := r.client.Exec(bookInsertQuery,
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
		return err
	}
	return nil
}
