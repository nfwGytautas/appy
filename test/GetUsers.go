package test_api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/nfwGytautas/appy-go/driver"
	"github.com/nfwGytautas/appy-go/http/extensions/utility"
	appy_tracker "github.com/nfwGytautas/appy-go/tracker"
)

type Users struct{}

type Private struct {
	ctx         context.Context
	scope       appy_tracker.Scope
	transaction appy_tracker.Transaction
	tx          driver.Tx
}

type User struct {
	userId   uint64
	username string
}

type Page struct {
	page utility.PagingSettings
}

func SetupPrivate(c *gin.Context) (*Private, error) {
	currentFunctionName := "SetupPrivate"
	ctx, scope, transaction := appy_tracker.Begin(c.Request.Context(), currentFunctionName)

	return &Private{
		ctx:         ctx,
		scope:       scope,
		transaction: transaction,
	}, nil
}

...

func GetAllUsers(private Private, user User, page Page) ([]Users, error) {

	return nil, nil
}
