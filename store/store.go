package store

import (
	"context"
	"math/rand"
	"time"
	"fmt"

	"github.com/NickMoorman123/receipt-processor/objects"
)

type IReceiptStore interface {
	Get(ctx context.Context, in *objects.GetRequest) (*objects.Receipt, error)
	Process(ctx context.Context, in *objects.ProcessRequest) error
}

func init() {
	rand.Seed(time.Now().UTC().Unix())
}

func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v%s", now.Nanosecond(), string(word))
}