package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) Positions() []string {
	return c.positionsOnce.Do(func() interface{} {
		config := struct {
			PositionsList []string `fig:"list"`
		}{}

		raw := kv.MustGetStringMap(c.getter, "positions")

		err := figure.
			Out(&config).
			From(raw).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out positions list"))
		}

		return config.PositionsList
	}).([]string)
}
