package mongodb

import (
	"context"
	fake "github.com/brianvoe/gofakeit/v7"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *MongoDBTestSuite) TestInsert() {
	s.T().Parallel()

	s.Run("insert user", func() {
		user := &User{
			ID:       primitive.NewObjectID(),
			Username: fake.Username(),
		}
		s.create(s.db.Collection(usersCol), user)
		s.Require().NotEmpty(user.ID)
	})

	s.Run("insert post", func() {
		users := s.db.Collection(usersCol)
		posts := s.db.Collection(postsCol)
		locations := s.db.Collection(locationsCol)

		author := User{
			ID:       primitive.NewObjectID(),
			Username: fake.Username(),
		}
		s.create(users, author)

		location := Location{
			ID:   primitive.NewObjectID(),
			Name: fake.City(),
			Lat:  fake.Latitude(),
			Lon:  fake.Longitude(),
		}

		s.create(locations, location)

		post := Post{
			UserID:     author.ID,
			Text:       fake.Phrase(),
			CreatedAt:  fake.PastDate(),
			Tags:       []string{fake.Word(), fake.Word()},
			LocationID: location.ID,
		}
		s.create(posts, post)
	})
	s.Run("insert many", func() {
		users := s.db.Collection(usersCol)

		data := lo.Times(3, func(i int) interface{} {
			return &User{
				ID:       primitive.NewObjectID(),
				Username: fake.Username(),
			}
		})
		result, err := users.InsertMany(context.TODO(), data, options.InsertMany().SetBypassDocumentValidation(true))
		s.Require().NoError(err)
		s.T().Cleanup(func() {
			r, err := users.DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", result.InsertedIDs}}}})
			s.Require().NoError(err)
			s.Require().Equal(int64(len(result.InsertedIDs)), r.DeletedCount)
		})
	})

	s.Run("duplicate key error", func() {
		const duplicateKeyError = 11000

		users := s.db.Collection(usersCol)

		id := primitive.NewObjectID()

		_, err := users.InsertMany(
			context.TODO(),
			[]interface{}{
				&User{ID: id, Username: fake.Username()},
				&User{ID: id, Username: fake.Username()},
			},
		)
		s.T().Cleanup(func() {
			r, err := users.DeleteOne(context.TODO(), bson.D{{"_id", id}})
			s.Require().NoError(err)
			s.Require().Equal(int64(1), r.DeletedCount)
		})

		s.Require().Error(err)
		var asError mongo.BulkWriteException
		s.Require().ErrorAs(err, &asError)
		s.Require().True(asError.HasErrorCode(duplicateKeyError))
	})
}
