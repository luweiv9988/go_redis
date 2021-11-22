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

	_ = s.Init()
	_, err := s.Client.Set(key, value, expire).Result()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return err

}

func (s *Storage) Getkeys(key ...string) interface{} {

	var dataList []string

	_ = s.Init()

	number := len(key)

	for item := 0; item < number; item++ {

		result, err := s.Client.Get(key[item]).Result()
		if err != nil {
			log.Print(err)
			os.Exit(-1)
		}
		dataList = append(dataList, result)
	}

	return dataList
}

func (s *Storage) GetAllKeys() interface{} {

	_ = s.Init()

	result, err := s.Client.Do("keys", "*").Result()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	return result

}

func (s *Storage) Exists(key string) int64 {

	_ = s.Init()

	result, err := s.Client.Exists(key).Result()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return result
}

// TTL: ns
func (s *Storage) Expirekey(key string, ttl time.Duration) error {

	_ = s.Init()

	_, err := s.Client.Expire(key, ttl*10000000000).Result()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return err

}

func (s *Storage) Close() {
	s.Client.Close()
}
