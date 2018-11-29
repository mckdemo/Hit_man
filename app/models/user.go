package models

import (
	"time"
)

type User struct {
	Username          string `json:"username" bson:"username"`
	Team              string `json:"team" bson:"team"`
	Score             int32  `json:"score" bson:"score"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
}

func NewUser(username string) *User {

	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{Username: username, Score: 0, Team: "A", TimestampCreated: unixTimestamp}
	return &u
}
