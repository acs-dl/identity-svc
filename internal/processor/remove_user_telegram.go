package processor

import (
	"strconv"

	"github.com/acs-dl/identity-svc/internal/data"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (p *processor) validateRemoveUserTelegram(msg data.ModulePayload) error {
	return validation.Errors{
		"user_id": validation.Validate(msg.UserId, validation.Required),
	}.Filter()
}

func (p *processor) handleRemoveTelegramAction(msg data.ModulePayload) error {
	p.log.Infof("start handle message with id `%s`", msg.RequestId)

	err := p.validateRemoveUserTelegram(msg)
	if err != nil {
		p.log.WithError(err).Errorf("failed to validate message with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to validate message")
	}

	_, userId, err := p.parseIdAndGetUser(msg.UserId)
	if err != nil {
		p.log.WithError(err).Errorf("failed to parse id and get user for message with id `%s`", msg.RequestId)
		return errors.Wrap(err, "failed to parse id and get user")
	}

	err = p.usersQ.UpdateTelegram(data.User{
		Id:       *userId,
		Telegram: nil,
	})
	if err != nil {
		p.log.WithError(err).Errorf("failed to update telegram info for user id `%s` for message with id `%s`", msg.UserId, msg.RequestId)
		return errors.Wrap(err, "failed to update telegram info")
	}

	p.log.Infof("finish handle message with id `%s`", msg.RequestId)
	return nil
}

func (p *processor) parseIdAndGetUser(id string) (*data.User, *int64, error) {
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse user id")
	}

	user, err := p.usersQ.GetById(userId)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get user by id")
	}
	if user == nil {
		return nil, nil, errors.Wrap(err, "no user with such id")
	}

	return user, &userId, err
}
