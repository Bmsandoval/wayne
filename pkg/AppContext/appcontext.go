package AppContext

import (
	"context"
	"github.com/bmsandoval/wayne/configs"
	"github.com/bmsandoval/wayne/pkg/DbContext"
)

type Context struct {
	DB        DbContext.Connection
	Config    configs.Configuration
	GoContext context.Context
}