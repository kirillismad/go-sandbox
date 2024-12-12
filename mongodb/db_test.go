package mongodb

import (
	_ "compress/zlib"
	"context"
	_ "github.com/golang/snappy"
	_ "github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const (
	uri          = "mongodb://localhost:27017"
	username     = "usr"
	password     = "pwd"
	dbName       = "test"
	usersCol     = "users"
	locationsCol = "locations"
	postsCol     = "posts"
)

type MongoDBTestSuite struct {
	suite.Suite
	db *mongo.Database
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MongoDBTestSuite))
}

func (s *MongoDBTestSuite) SetupSuite() {
	opts := options.Client()
	opts = opts.ApplyURI(uri)
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
		r, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", result.InsertedID}})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), r.DeletedCount)
	})
}
