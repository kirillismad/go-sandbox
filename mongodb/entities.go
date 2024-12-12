package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
}

type Location struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Lat  float64            `bson:"lat"`
	Lon  float64            `bson:"lon"`
}

type Commend struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	PostID    primitive.ObjectID `bson:"post_id"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"created_at"`
}

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Text       string             `bson:"text"`
	CreatedAt  time.Time          `bson:"created_at"`
	Tags       []string           `bson:"tags"`
	LocationID primitive.ObjectID `bson:"location_id"`
}
