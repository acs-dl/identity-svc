package service

import (
	"context"
	"sync"

	"gitlab.com/distributed_lab/acs/identity-svc/internal/receiver"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/registrator"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/api"

	"gitlab.com/distributed_lab/acs/identity-svc/internal/config"
	"gitlab.com/distributed_lab/acs/identity-svc/internal/service/types"
)

var availableServices = map[string]types.Runner{
	"api":       api.Run,
	"receiver":  receiver.Run,
	"registrar": registrator.Run,
}

func Run(cfg config.Config) {
	logger := cfg.Log().WithField("service", "main")
	ctx := context.Background()
	wg := new(sync.WaitGroup)

	logger.Info("Starting all available services...")

	for serviceName, service := range availableServices {
		wg.Add(1)

		go func(name string, runner types.Runner) {
			defer wg.Done()

			runner(ctx, cfg)

		}(serviceName, service)

		logger.WithField("service", serviceName).Info("Service started")
	}

	wg.Wait()

}
