package repository

import (
	"context"

	"github.com/tabrizihamid84/library-application/domain"
	"github.com/tabrizihamid84/library-application/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bookRepository struct {
	database   mongo.Database
	collection string
}

func NewBookRepository(db mongo.Database, collection string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: collection,
	}
}

// Create implements domain.BookRepository.
func (b *bookRepository) Create(ctx context.Context, book *domain.Book) error {
	collection := b.database.Collection(b.collection)

	_, err := collection.InsertOne(ctx, book)

	return err
}

// Delete implements domain.BookRepository.
func (b *bookRepository) Delete(ctx context.Context, id int) error {
	collection := b.database.Collection(b.collection)

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})

	return err
}

// GetAll implements domain.BookRepository.
func (r *bookRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	collection := r.database.Collection(r.collection)
	var books []domain.Book

	// Pagination parameters
	page := int64(1)      // Current page (1-based)
	pageSize := int64(10) // Number of documents per page

	// Calculate skip value
	skip := (page - 1) * pageSize

	// Find options for pagination
	findOptions := options.Find()
	findOptions.SetLimit(pageSize)
	findOptions.SetSkip(skip)

	results, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var book domain.Book
		if err = results.Decode(&book); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

// GetById implements domain.BookRepository.
func (r *bookRepository) GetById(ctx context.Context, id int) (domain.Book, error) {
	collection := r.database.Collection(r.collection)
	var book domain.Book
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&book)
	return book, err
}

// Update implements domain.BookRepository.
func (r *bookRepository) Update(ctx context.Context, id int, book *domain.Book) error {
	collection := r.database.Collection(r.collection)
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": book},
	)
	return err
}
