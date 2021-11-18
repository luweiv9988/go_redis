package redisStorage

import (
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Storage struct {
	Address  string
	Password string
	DB       int
	Client   *redis.Client
	Expires  time.Duration
}

func (s *Storage) Init() error {
	if s.Client == nil {
		s.Client = redis.NewClient(&redis.Options{
			Addr:     s.Address,
			Password: s.Password,
			DB:       s.DB,
		})

	}
	_, err := s.Client.Ping().Result()
	if err != nil {
		log.Println("Redis Connection Error: ", err)
		os.Exit(-1)
	}

	return err
}

func (s *Storage) Insert(key string, value interface{}, expire time.Duration) error {
	_=s.Init()
	_, err := s.Client.Set(key, value, expire).Result()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return err

}
