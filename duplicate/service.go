package duplicate

import (
	"log"

	"github.com/go-redis/redis/v7"
)

type Service struct {
	Client *redis.Client
}

func Set(client *redis.Client, url string) error {
	err := client.Set(url, 1, 0).Err()
	if err != nil {
		log.Printf("Error saving URL %s: %v.", url, err)
		return err
	}
	log.Printf("URL %s saved to redis service.", url)
	return nil
}

func (s *Service) Exists(url string, result *bool) error {
	val, err := s.Client.Get(url).Int()
	*result = false
	if err == redis.Nil {
		log.Printf("URL %s doesn't exist. Add to Redis.", url)
		err = Set(s.Client, url)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	log.Printf("URL %s exists.", url)
	*result = val == 1
	return nil
}
