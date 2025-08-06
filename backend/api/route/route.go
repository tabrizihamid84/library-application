package route

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tabrizihamid84/library-application/bootstrap"
	"github.com/tabrizihamid84/library-application/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, echo *echo.Echo) {
	bookRoute := echo.Group("/books")
	NewBookRoute(env, timeout, db, bookRoute)

}
