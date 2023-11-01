package session

import (
	"math/rand"
	"errors"
	"time"
	"context"

	"github.com/redis/go-redis/v9"
)

type manager struct {
	redis *redis.Client
	sessionIDLength int
	lifeTime time.Duration
	ctx *context.Context
}

var sessionManager *manager = nil

func Init(ctx *context.Context) error {
	if sessionManager != nil {
		return errors.New("sessionManager is exist")
	}

	sessionManager = &manager{
		sessionIDLength:	16,
		lifeTime:	30*time.Second,
		ctx:	ctx,
	}
	sessionManager.redis = redis.NewClient(&redis.Options{
		Addr:	"localhost:6379",
		Password:	"",
		DB:	0,
	})

	return nil
}

func Create() (string, error) {
	if sessionManager == nil {
		return "", errors.New("sessionManager is not exist");
	}

	sessionID := CreateSessionID(sessionManager.sessionIDLength)
	err := sessionManager.redis.Set(*sessionManager.ctx, sessionID, time.Now(), sessionManager.lifeTime).Err()
	if err != nil {
		return "", err
	}

	return sessionID, nil
}

func CreateSessionID(sessionIDLength int) string {
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	length := len(str)
	sessionid := ""
	for i := 0; i < sessionIDLength; i++ {
		num := rand.Intn(length)
		sessionid = sessionid + str[num:num+1]
	}
	return sessionid
}

func Get(key string) (string, error) {
	if sessionManager == nil {
		return "", errors.New("sessionManager is not exist");
	}
	session, err := sessionManager.redis.Get(*sessionManager.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return session, nil

}