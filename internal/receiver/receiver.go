package receiver

import (
	"context"

	"gitlab.com/distributed_lab/acs/identity-svc/internal/config"
)

func Run(ctx context.Context, cfg config.Config) {
	NewReceiver(cfg).Run(ctx)
}
