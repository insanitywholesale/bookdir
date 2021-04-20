package postgres

var createBookTableQuery = `CREATE TABLE if not exists Book (
	ISBN VARCHAR(13),
	PRIMARY KEY (ISBN),
	Title VARCHAR(100),
	--CONSTRAINT FK_Author FOREIGN KEY (FirstName, MiddleName, LastName) REFERENCES Author(FirstName, MiddleName, LastName),
	Year INTEGER,
	Edition INTEGER,
	Pages INTEGER,
	Category VARCHAR(20),
	PDF BOOLEAN,
	Owned BOOLEAN
);`

var createAuthorTableQuery = `CREATE TABLE if not exists Author (
	FirstName VARCHAR(30),
	MiddleName VARCHAR(30),
	LastName VARCHAR(30),
	--CONSTRAINT PK_Author PRIMARY KEY (FirstName, MiddleName, LastName),
	YearBorn INTEGER,
	YearDied INTEGER,
	BooksWritten INTEGER
);`

var createPublisherTableQuery = `CREATE TABLE if not exists Publisher (
	Name VARCHAR(30),
	YearStarted INTEGER,
	YearEnded INTEGER,
	BooksPublished INTEGER
);`

var bookRetrievalQuery = `SELECT * FROM Book WHERE ISBN=$1`

var bookInsertQuery = `INSERT INTO books (
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
