package postgres

var createBookTableQuery = `CREATE TABLE if not exists Book (
	VARCHAR(13) ISBN
	VARCHAR(100) Title
	FOREIGN KEY Author
	INTEGER	Year
	INTEGER	Edition
	INTEGER Pages
	VARCHAR(20) Category
	BOOLEAN PDF
	BOOLEAN Owned
);`

var createAuthorTableQuery = `CREATE TABLE if not exists Author (
	VARCHAR(30) FirstName
	VARCHAR(30) MiddleName
	VARCHAR(30) LastName
	INTEGER	YearBorn
	INTEGER	YearDied
	INTEGER BooksWritter
);`

var createPublisherTableQuery = `CREATE TABLE if not exists Publisher (
	VARCHAR(30) Name
	INTEGER	YearStarted
	INTEGER	YearEnded
	INTEGER BooksPublished
);`

var bookRetrievalQuery = `SELECT * FROM Book WHERE ISBN=$1`

var bookInsertQuery = `INSERT INTO books (
		ISBN, Title, Author, Year, Edition, Pages, Category, PDF, Owned
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
