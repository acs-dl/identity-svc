package receiver

import (
	"context"

	"github.com/acs-dl/identity-svc/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewReceiver(cfg).Run(ctx)
}
