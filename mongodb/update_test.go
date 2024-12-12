package mongodb

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (s *MongoDBTestSuite) TestUpdate() {
	s.T().Parallel()

	s.Run("update by id", func() {
		r := s.Require()

		user := User{
			ID:       primitive.NewObjectID(),
			Username: gofakeit.Username(),
		}
		s.create(s.db.Collection(usersCol), &user)

		user.Username = gofakeit.Username()

		update := bson.D{{"$set", bson.D{{"username", user.Username}}}}

		result, err := s.db.Collection(usersCol).UpdateByID(context.TODO(), user.ID, update)
		r.NoError(err)
		r.Equal(int64(1), result.ModifiedCount)
		r.Equal(int64(1), result.MatchedCount)
	})

	s.Run("update one", func() {
		r := s.Require()

		user := User{
			ID:       primitive.NewObjectID(),
			Username: gofakeit.Username(),
		}
		s.create(s.db.Collection(usersCol), &user)

		user.Username = gofakeit.Username()

		filter := bson.D{{"_id", user.ID}}
		update := bson.D{{"$set", bson.D{{"username", user.Username}}}}

		result, err := s.db.Collection(usersCol).UpdateOne(context.TODO(), filter, update)
		r.NoError(err)
		r.Equal(int64(1), result.ModifiedCount)
		r.Equal(int64(1), result.MatchedCount)
	})

	s.Run("update many", func() {
		r := s.Require()

		user := User{
			ID:       primitive.NewObjectID(),
			Username: gofakeit.Username(),
		}
		s.create(s.db.Collection(usersCol), &user)

		posts, err := s.db.Collection(postsCol).InsertMany(context.TODO(), []interface{}{
			Post{
				ID:        primitive.NewObjectID(),
				UserID:    user.ID,
				Text:      gofakeit.Phrase(),
				CreatedAt: time.Now(),
				Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
			},
			Post{
				ID:        primitive.NewObjectID(),
				UserID:    user.ID,
				Text:      gofakeit.Phrase(),
				CreatedAt: time.Now(),
				Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
			},
		})
		r.NoError(err)
		r.Len(posts.InsertedIDs, 2)
		s.T().Cleanup(func() {
			del, err := s.db.Collection(postsCol).DeleteMany(context.TODO(), bson.D{{"_id", bson.D{{"$in", posts.InsertedIDs}}}})
			r.NoError(err)
			r.Equal(int64(2), del.DeletedCount)
		})

		filter := bson.D{{"user_id", user.ID}}
		update := bson.D{{"$set", bson.D{{"text", gofakeit.Phrase()}}}}

		result, err := s.db.Collection(postsCol).UpdateMany(context.TODO(), filter, update)
		r.NoError(err)
		r.Equal(int64(2), result.ModifiedCount)
		r.Equal(int64(2), result.MatchedCount)
	})

	s.Run("replace one", func() {
		r := s.Require()

		user := User{
			ID:       primitive.NewObjectID(),
			Username: gofakeit.Username(),
		}
		s.create(s.db.Collection(usersCol), &user)

		post := Post{
			ID:        primitive.NewObjectID(),
			UserID:    user.ID,
			Text:      gofakeit.Phrase(),
			CreatedAt: time.Now(),
			Tags:      []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()},
		}
		s.create(s.db.Collection(postsCol), &post)

		post.Text = gofakeit.Phrase()
		post.CreatedAt = time.Now()
		post.Tags = []string{gofakeit.Word(), gofakeit.Word(), gofakeit.Word()}

		filter := bson.D{{"user_id", post.UserID}}
		result, err := s.db.Collection(postsCol).ReplaceOne(context.TODO(), filter, post)
		r.NoError(err)
		r.Equal(int64(1), result.ModifiedCount)
		r.Equal(int64(1), result.MatchedCount)
	})
}
