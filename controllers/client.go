package clientController

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	token "github.com/sahildotexe/go-rate-limit/utils"
)

type Client struct {
	Key       string `json:"key"`
	MaxTokens int64  `json:"maxTokens"`
	FillRate  int64  `json:"fillRate"`
}

var clientBuckets = make(map[string]*token.TokenBucket)

type Rule struct {
	MaxTokens int64
	Rate      int64
}

func GenerateClientKey(c *gin.Context) {
	key := uuid.New().String()
	clientBuckets[key] = token.NewTokenBucket(1, 10)
	c.JSON(200, gin.H{"key": key})
}

func GetBucket(key string) (*token.TokenBucket, error) {
	fmt.Println(clientBuckets)
	clientBucket := clientBuckets[key]

	if clientBucket == nil {
		return nil, errors.New("Client not found")
	}

	return clientBucket, nil
}
