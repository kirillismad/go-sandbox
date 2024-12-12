//go:build mongodb

package mongodb

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *MongoDBTestSuite) TestDelete() {
	s.T().Parallel()

	s.Run("delete one", func() {
		user := User{
			ID:       primitive.NewObjectID(),
			Username: gofakeit.Username(),
		}

		_, err := s.db.Collection(usersCol).InsertOne(context.TODO(), &user)
		s.Require().NoError(err)

		_, err = s.db.Collection(usersCol).DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: user.ID}})
		s.Require().NoError(err)

	})

	s.Run("delete many", func() {
		users := []interface{}{
			User{
				ID:       primitive.NewObjectID(),
				Username: gofakeit.Username(),
			},
			User{
				ID:       primitive.NewObjectID(),
				Username: gofakeit.Username(),
			},
		}

		r, err := s.db.Collection(usersCol).InsertMany(context.TODO(), users)
		s.Require().NoError(err)

		_, err = s.db.Collection(usersCol).DeleteMany(context.TODO(), bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: r.InsertedIDs}}}})
		s.Require().NoError(err)
	})
}
