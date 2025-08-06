package domain

import (
	"context"

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

type BookRepository interface {
	Create(ctx context.Context, book *Book) error
	Update(ctx context.Context, id int, book *Book) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
}

type BookUsecase interface {
	Create(ctx context.Context, book *Book) error
	Update(ctx context.Context, id int, book *Book) error
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (Book, error)
	GetAll(ctx context.Context) ([]Book, error)
}
