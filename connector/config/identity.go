package config

import (
	"gitlab.com/distributed_lab/acs/identity-svc/connector"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type IdentityConfiger interface {
	IdentityConfig() *IdentityConfig
	IdentityConnector() connector.ConnectorI
}

type IdentityConfig struct {
	ServiceUrl string `fig:"service_url,required"`
}

func NewIdentityConfiger(getter kv.Getter) IdentityConfiger {
	return &identityConfig{
		getter: getter,
	}
}

type identityConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *identityConfig) IdentityConfig() *IdentityConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "identity")
		config := IdentityConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out identity"))
		}

		return &config
	}).(*IdentityConfig)
}

func (c *identityConfig) IdentityConnector() connector.ConnectorI {
	return connector.NewConnector(c.IdentityConfig().ServiceUrl)
}
