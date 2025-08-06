package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tabrizihamid84/library-application/domain"
)

type BookController struct {
	BookUsecase domain.BookUsecase
}

func (c *BookController) GetAll(ec echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	books, err := c.BookUsecase.GetAll(ctx)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return ec.JSON(http.StatusOK, domain.SuccessResponse{Message: "Books fetched successfully.", Data: books})
}
