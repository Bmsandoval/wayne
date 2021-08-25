package appcontext

import (
	"context"
	"github.com/bmsandoval/wayne/configs"
	"github.com/bmsandoval/wayne/db"
)

type Context struct {
	DB db.Connection
	Config configs.Configuration
	GoContext context.Context
}