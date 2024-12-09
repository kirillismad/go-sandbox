package mongodb

import "time"

type User struct {
	ID       string `bson:"_id"`
	Username string `bson:"username"`
}

type Location struct {
	ID   string  `bson:"_id"`
	Name string  `bson:"name"`
	Lat  float64 `bson:"lat"`
	Lon  float64 `bson:"lon"`
}

type Commend struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	User      *User     `bson:"user"`
	PostID    string    `bson:"post_id"`
	Text      string    `bson:"text"`
	CreatedAt time.Time `bson:"created_at"`
}

type Post struct {
	ID         string    `bson:"_id"`
	UserID     string    `bson:"user_id"`
	Text       string    `bson:"text"`
	CreatedAt  time.Time `bson:"created_at"`
	Tags       []string  `bson:"tags"`
	LocationID string    `bson:"location_id"`
}
