package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBook = "books"
)

type Book struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title" validate:"required"`
	Author         string             `bson:"author" json:"author" validate:"required"`
	PublishYear    int                `bson:"publishYear" json:"publishYear" validate:"required"`
	CreatedAt      primitive.DateTime `bson:"createdAt" json:"-"`
	LastModifiedAt primitive.DateTime `bson:"lastModifiedAt" json:"-"`
}
