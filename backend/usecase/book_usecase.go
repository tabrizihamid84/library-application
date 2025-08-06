package usecase

import (
	"context"
	"time"

	"github.com/tabrizihamid84/library-application/domain"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	cotextTimeout  time.Duration
}

func NewBookUsecase(bookRepository domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepository,
		cotextTimeout:  timeout,
	}
}

// Create implements domain.BookUsecase.
func (u *bookUsecase) Create(ctx context.Context, book *domain.Book) error {
	return u.bookRepository.Create(ctx, book)
}

// Delete implements domain.BookUsecase.
func (u *bookUsecase) Delete(ctx context.Context, id int) error {
	return u.bookRepository.Delete(ctx, id)
}

// GetAll implements domain.BookUsecase.
func (u *bookUsecase) GetAll(ctx context.Context) ([]domain.Book, error) {
	return u.bookRepository.GetAll(ctx)
}

// GetById implements domain.BookUsecase.
func (u *bookUsecase) GetById(ctx context.Context, id int) (domain.Book, error) {
	return u.bookRepository.GetById(ctx, id)
}

// Update implements domain.BookUsecase.
func (u *bookUsecase) Update(ctx context.Context, id int, book *domain.Book) error {
	return u.bookRepository.Update(ctx, id, book)
}
