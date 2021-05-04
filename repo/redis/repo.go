package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	pb "gitlab.com/insanitywholesale/bookdir/proto/v1"
)

type redisRepo struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: redisURL})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepo(redisURL string) (*redisRepo, error) {
	redisclient, err := newRedisClient(redisURL)
	if err != nil {
		return nil, err
	}
	return &redisRepo{client: redisclient}, nil
}

func (r *redisRepo) generateKey(code string) string {
	return fmt.Sprintf("book:%s", code)
}

func (r *redisRepo) RetrieveAllAuthors() ([]*pb.Author, error) {
	var authorlist []*pb.Author
	keys, err := r.client.Do("KEYS", "book:*").Result()
	if err != nil {
		return nil, err
	}

	for _, b := range keys.([]interface{}) {
		book := &pb.Book{}
		reply, err := r.client.Do("GET", b.(string)).Result()
		if err != nil {
			return nil, err
		}
		fmt.Println("reply", reply)
		err = json.Unmarshal([]byte(reply.(string)), book)
		if err != nil {
			return nil, err
		}
		authorlist = append(authorlist, book.Author)
	}
	return authorlist, nil
}

func (r *redisRepo) RetrieveBooksByAuthor(authorId string) ([]*pb.Book, error) {
	var booklist []*pb.Book
	keys, err := r.client.Do("KEYS", "book:*").Result()
	if err != nil {
		return nil, err
	}

	for _, b := range keys.([]interface{}) {
		book := &pb.Book{}
		reply, err := r.client.Do("GET", b.(string)).Result()
		if err != nil {
			return nil, err
		}
		fmt.Println("reply", reply)
		err = json.Unmarshal([]byte(reply.(string)), book)
		if err != nil {
			return nil, err
		}
		authoriduint, err := strconv.ParseUint(authorId, 10, 64)
		if err != nil {
			return nil, err
		}
		if book.Author.AuthorID == uint32(authoriduint) {
			booklist = append(booklist, book)
		}
	}
	if len(booklist) == 0 {
		return nil, errors.New("no books by this author were found")
	} else {
		return booklist, nil
	}
}

func (r *redisRepo) RetrieveAuthorById(authorId string) (*pb.Author, error) {
	var author *pb.Author
	return author, nil
}

func (r *redisRepo) RetrieveAll() ([]*pb.Book, error) {
	var booklist []*pb.Book
	keys, err := r.client.Do("KEYS", "book:*").Result()
	if err != nil {
		return nil, err
	}

	for _, b := range keys.([]interface{}) {
		book := &pb.Book{}
		reply, err := r.client.Do("GET", b.(string)).Result()
		if err != nil {
			return nil, err
		}
		fmt.Println("reply", reply)
		err = json.Unmarshal([]byte(reply.(string)), book)
		if err != nil {
			return nil, err
		}
		booklist = append(booklist, book)
	}
	return booklist, nil
}

func (r *redisRepo) Retrieve(isbn string) (*pb.Book, error) {
	book := &pb.Book{}
	key := r.generateKey(isbn)
	data, err := r.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("book not found")
	}
	err = json.Unmarshal([]byte(data), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *redisRepo) Save(book *pb.Book) error {
	key := r.generateKey(book.ISBN)
	authId, err := r.client.Incr("AuthorID").Result()
	book.Author.AuthorID = uint32(authId)
	pubId, err := r.client.Incr("PublisherID").Result()
	book.Publisher.PublisherID = uint32(pubId)
	data, err := json.Marshal(book)
	if err != nil {
		return err
	}
	err = r.client.Set(key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
