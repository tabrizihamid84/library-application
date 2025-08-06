package controller

import (
	"context"
	"net/http"
	"time"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/tabrizihamid84/library-application/domain"
)

var validate = validator.New()

type BookController struct {
	BookUsecase domain.BookUsecase
}

func (bc *BookController) Create(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book domain.Book
	defer cancel()

	//validate the request body
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&book); validationErr != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: validationErr.Error()})
	}

	err := bc.BookUsecase.Create(ctx, &book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Book created successfully."})
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

func (bc *BookController) Delete(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Param("id")
	if idParam == "" {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Missing book ID"})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid book ID"})
	}

	err = bc.BookUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Book deleted successfully."})
}
