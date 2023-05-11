package processor

import (
	"github.com/acs-dl/identity-svc/internal/data"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateUpdateUserTelegram(msg data.ModulePayload) error {
	phoneValidationCase := validation.When(msg.Username == nil, validation.Required.Error("phone is required if username is not set"))
	usernameValidationCase := validation.When(msg.Phone == nil, validation.Required.Error("username is required if phone is not set"))

	return validation.Errors{
		"user_id":  validation.Validate(msg.UserId, validation.Required),
		"username": validation.Validate(msg.Username, usernameValidationCase),
		"phone":    validation.Validate(msg.Phone, phoneValidationCase),
	}.Filter()
}

func (p *processor) handleUpdateTelegramAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message with id `%s`", msg.RequestId)

	err := p.validateUpdateUserTelegram(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate message with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate message")
	}

	var telegramInfo string
	if msg.Username != nil {
		telegramInfo = *msg.Username
	} else if msg.Phone != nil {
		telegramInfo = *msg.Phone
	} else {
		p.log.Errorf("no telegram info was specified")
		return errors.Errorf("no telegram info was specified")
	}

	_, userId, err := p.parseIdAndGetUser(msg.UserId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse id and get user for message with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to parse id and get user")
	}

	err = p.usersQ.UpdateTelegram(data.User{
		Id:       *userId,
		Telegram: &telegramInfo,
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to update telegram info `%s` for user id `%s` for message with id `%s`", telegramInfo, msg.UserId, msg.RequestId)
		return errors.Wrap(err, "failed to update telegram info")
	}

	p.log.Infof("finish handle message with id `%s`", msg.RequestId)
	return nil
}
