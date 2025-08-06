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
func (u *bookUsecase) Create(c context.Context, book *domain.Book) error {
	return u.bookRepository.Create(c, book)
}

// Delete implements domain.BookUsecase.
func (*bookUsecase) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetAll implements domain.BookUsecase.
func (u *bookUsecase) GetAll(c context.Context) ([]domain.Book, error) {
	return u.bookRepository.GetAll(c)
}

// GetById implements domain.BookUsecase.
func (*bookUsecase) GetById(ctx context.Context, id int) (domain.Book, error) {
	panic("unimplemented")
}

// Update implements domain.BookUsecase.
func (*bookUsecase) Update(ctx context.Context, id int, book *domain.Book) error {
	panic("unimplemented")
}
