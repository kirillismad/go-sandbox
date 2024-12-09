package mongodb

import (
	"context"
	fake "github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func create(collection *mongo.Collection, item interface{}, t *testing.T) string {
	result, err := collection.InsertOne(context.TODO(), item)
	require.NoError(t, err)
	t.Cleanup(func() {
		result, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", result.InsertedID}})
		require.NoError(t, err)
		require.Equal(t, int64(1), result.DeletedCount)
	})
	return result.InsertedID.(string)
}

func TestInsert(t *testing.T) {
	t.Parallel()

	t.Run("add user", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		client := getClient(context.TODO(), t)

		t.Cleanup(func() {
			r.NoError(client.Disconnect(context.TODO()))
		})

		user := &User{
			Username: fake.Username(),
		}
		userID := create(client.Database("test").Collection("users"), user, t)
		user.ID = userID
		require.NotEmpty(t, user.ID)
	})

	t.Run("add post", func(t *testing.T) {
		t.Parallel()

		client := getClient(context.TODO(), t)

		users := client.Database("test").Collection("users")
		posts := client.Database("test").Collection("posts")
		locations := client.Database("test").Collection("locations")

		author := User{
			Username: fake.Username(),
		}
		authorID := create(users, author, t)
		author.ID = authorID

		location := Location{
			Name: fake.City(),
			Lat:  fake.Latitude(),
			Lon:  fake.Longitude(),
		}

		locationID := create(locations, location, t)
		location.ID = locationID

		post := Post{
			UserID:     authorID,
			Text:       fake.Phrase(),
			CreatedAt:  fake.PastDate(),
			Tags:       []string{fake.Word(), fake.Word()},
			LocationID: location.ID,
		}
		postID := create(posts, post, t)
		post.ID = postID
	})
}
