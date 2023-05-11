package processor

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/acs-dl/identity-svc/internal/config"
	"github.com/acs-dl/identity-svc/internal/data"
	"github.com/acs-dl/identity-svc/internal/data/postgres"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	serviceName = data.ModuleName + "-processor"

	//add needed actions for module
	UpdateTelegramAction = "update_telegram"
	RemoveTelegramAction = "remove_telegram"
)

type Processor interface {
	HandleNewMessage(msg data.ModulePayload) error
}

type processor struct {
	log    *logan.Entry
	usersQ data.UsersQ
}

var handleActions = map[string]func(proc *processor, msg data.ModulePayload) error{
	UpdateTelegramAction: (*processor).handleUpdateTelegramAction,
	RemoveTelegramAction: (*processor).handleRemoveTelegramAction,
}

func NewProcessor(cfg config.Config) Processor {
	return &processor{
		log:    cfg.Log().WithField("service", serviceName),
		usersQ: postgres.NewUsersQ(cfg.DB()),
	}
}

func (p *processor) HandleNewMessage(msg data.ModulePayload) error {
	p.log.Infof("handling message with id `%s`", msg.RequestId)

	err := validation.Errors{
		"action": validation.Validate(msg.Action, validation.Required, validation.In(UpdateTelegramAction)),
	}.Filter()
	if err != nil {
		p.log.WithError(err).Errorf("no such action `%s` to handle for message with id `%s`", msg.Action, msg.RequestId)
		return errors.Wrap(err, fmt.Sprintf("no such action `%s` to handle for message with id `%s`", msg.Action, msg.RequestId))
	}

	requestHandler := handleActions[msg.Action]
	if err = requestHandler(p, msg); err != nil {
		p.log.WithError(err).Errorf("failed to handle message with id `%s`", msg.RequestId)
		return err
	}

	p.log.Infof("finish handling message with id `%s`", msg.RequestId)
	return nil
}
