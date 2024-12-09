package mongodb

import (
	_ "compress/zlib"
	"context"
	_ "github.com/golang/snappy"
	_ "github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const (
	uri      = "mongodb://localhost:27017"
	username = "username"
	password = "password"
)

func getClient(ctx context.Context, t *testing.T) *mongo.Client {
	opts := options.Client()
	opts = opts.ApplyURI(uri)
	opts = opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})
	opts = opts.SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	opts = opts.SetCompressors([]string{"snappy", "zstd", "zlib"})
	client, err := mongo.Connect(ctx, opts)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, client.Disconnect(context.TODO()))
	})
	return client
}

func TestConnection(t *testing.T) {
	t.Parallel()
	t.Run("default", func(t *testing.T) {
		t.Parallel()

		client := getClient(context.TODO(), t)
		err := client.Ping(context.TODO(), nil)
		require.NoError(t, err)
	})
}
