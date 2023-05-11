package types

import (
	"context"

	"github.com/acs-dl/identity-svc/internal/config"
)

type Runner = func(context context.Context, config config.Config)
