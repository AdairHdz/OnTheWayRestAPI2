package tokenBlacklist

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)


var (
	rdb *redis.Client
)

type tokenBlackList struct { }

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "host.docker.internal:6379",
		Password: "",
		DB: 0,
	})	
}

func GetInstance() *tokenBlackList {	
	return &tokenBlackList{}
}

func (tokenBlackList) Save(key string, value interface{}, expiration time.Duration) error {
	redisStatus := rdb.Set(context.Background(), key, value, expiration)
	_, err := redisStatus.Result()
	if err != nil {
		return err
	}	
	return nil
}

func (tokenBlackList) Get(key string) ([]byte, error) {
	redisStatus := rdb.Get(context.Background(), key)	
	result, err := redisStatus.Result()
	resultParsedInBytes := []byte(result)
	
	return resultParsedInBytes, err
}


type UserBlackList struct {
	EmailAddress string
	AssociatedTokens []string
}

func (userBlackList UserBlackList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(userBlackList)
}

func (userBlackList *UserBlackList) UnmarshalBinary(data []byte) error{
	return json.Unmarshal(data, &userBlackList)
}