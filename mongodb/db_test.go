//go:build mongodb

package mongodb

import (
	_ "compress/zlib"
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/golang/snappy"
	_ "github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName       = "db"
	usersCol     = "users"
	locationsCol = "locations"
	postsCol     = "posts"
)

var host string
var port string
var username string
var password string

func init() {
	host = os.Getenv("MONGO_HOST")
	if host == "" {
		log.Fatal("MONGO_HOST is not set")
	}
	port = os.Getenv("MONGO_PORT")
	if port == "" {
		log.Fatal("MONGO_PORT is not set")
	}
	username = os.Getenv("MONGO_USER")
	if username == "" {
		log.Fatal("MONGO_USER is not set")
	}
	password = os.Getenv("MONGO_PASSWORD")
	if password == "" {
		log.Fatal("MONGO_PASSWORD is not set")
	}
}

type MongoDBTestSuite struct {
	suite.Suite
	db *mongo.Database
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MongoDBTestSuite))
}

func (s *MongoDBTestSuite) SetupSuite() {
	opts := options.Client()
	opts = opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s", host, port))
	opts = opts.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})
	opts = opts.SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
	opts = opts.SetCompressors([]string{"snappy", "zstd", "zlib"})
	client, err := mongo.Connect(context.TODO(), opts)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		s.Require().NoError(client.Disconnect(context.Background()))
	})
	s.db = client.Database(dbName)
}

func (s *MongoDBTestSuite) create(collection *mongo.Collection, item interface{}) {
	result, err := collection.InsertOne(context.TODO(), item)
	s.Require().NoError(err)
	s.T().Cleanup(func() {
		r, err := collection.DeleteOne(
			context.TODO(),
			bson.D{{
				Key:   "_id",
				Value: result.InsertedID,
			}},
		)
		s.Require().NoError(err)
		s.Require().Equal(int64(1), r.DeletedCount)
	})
}
