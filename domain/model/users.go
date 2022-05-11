package model

import (
	"time"
)

type InsertUsers struct {
	ID        string    `bson:"_id"`
	UserID    int       `bson:"UserID"`
	UserName  string    `bson:"UserName"`
	IsDeleted bool      `bson:"isDeleted"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	IsActive  bool      `bson:"isActive"`
}

type User struct {
	ID        string    `bson:"_id" json:"id"`
	UserID    string    `bson:"userID" json:"user_id"`
	Type      string    `bson:"type" json:"type"`
	IsDeleted bool      `bson:"isDeleted" json:"is_deleted"`
	CreatedBy string    `bson:"createdBy" json:"created_by"`
	UpdatedBy string    `bson:"updatedBy" json:"updated_by"`
	CreatedAt time.Time `bson:"createdAt" json:"created_at"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updated_at"`
	IsActive  bool      `bson:"isActive" json:"is_active"`
}
