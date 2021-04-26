package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	pb "gitlab.com/insanitywholesale/proto/v1"
	"log"
)

type redisRepo struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
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
	return &repo{client: redisclient}, nil
}

func (r *redisRepo) generateKey(code string) string {
	return fmt.Sprintf("book:%s", code)
}

func (r *redisRepo) RetrieveAll() ([]*pb.Book, error) {
	return nil, nil
}
func (r *redisRepo) Retrieve(isbn string) (*pb.Book, error) {
	return nil, nil
}

func (r *redisRepo) Save(book *pb.Book) error {
	key := r.generateKey(book.ISBN)
	data := map[string]interface{}{
		"isbn":      book.ISBN,
		"title":     book.Title,
		"author":    book.Author,
		"year":      book.Year,
		"edition":   book.Edition,
		"publisher": book.Publisher,
		"pages":     book.Pages,
		"category":  book.Category,
		"pdf":       book.PDF,
		"owned":     book.Owned,
	}

	res, err := r.client.HMSet(key, data).Result()
	log.Println("redis res:", res)
	if err != nil {
		return err
	}
	return nil
}
