package postgres

var createBookTableQuery = `CREATE TABLE if not exists Book (
	ISBN VARCHAR(13) PRIMARY KEY,
	Title VARCHAR(100),
	Author SERIAL REFERENCES Author(AuthorID),
	Year INTEGER,
	Edition INTEGER,
	Publisher SERIAL REFERENCES Publisher(PublisherID),
	Pages INTEGER,
	Category VARCHAR(20),
	PDF BOOLEAN,
	Owned BOOLEAN
);`

var createAuthorTableQuery = `CREATE TABLE if not exists Author (
	AuthorID SERIAL PRIMARY KEY,
	FirstName VARCHAR(30),
	MiddleName VARCHAR(30),
	LastName VARCHAR(30),
	YearBorn INTEGER,
	YearDied INTEGER,
	BooksWritten INTEGER
);`

var createPublisherTableQuery = `CREATE TABLE if not exists Publisher (
	PublisherID SERIAL PRIMARY KEY,
	Name VARCHAR(30),
	YearStarted INTEGER,
	YearEnded INTEGER,
	BooksPublished INTEGER
);`

var bookRetrievalQuery = `SELECT * FROM Book WHERE ISBN=$1`
var bookRetrieveAllQuery = `SELECT * FROM Book`

var bookInsertQuery = `INSERT INTO Book (
	ISBN,
	Title,
	Author,
	Year,
	Edition,
	Pages,
	Category,
	PDF,
	Owned
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
