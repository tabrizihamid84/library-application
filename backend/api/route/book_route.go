package route

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tabrizihamid84/library-application/api/controller"
	"github.com/tabrizihamid84/library-application/bootstrap"
	"github.com/tabrizihamid84/library-application/domain"
	"github.com/tabrizihamid84/library-application/mongo"
	"github.com/tabrizihamid84/library-application/repository"
	"github.com/tabrizihamid84/library-application/usecase"
)

func NewBookRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *echo.Group) {
	r := repository.NewBookRepository(db, domain.CollectionBook)
	c := &controller.BookController{
		BookUsecase: usecase.NewBookUsecase(r, timeout),
	}

	group.GET("", c.GetAll)
	group.POST("", c.Create)
}
