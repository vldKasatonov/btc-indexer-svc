package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const signerConfigKey = "signer"

type Signer interface {
	SignerConfig() *SignerConfig
}

type SignerConfig struct {
	JwtSecret string `fig:"jwt_secret"`
}

func NewSigner(getter kv.Getter) Signer {
	return &signer{
		getter: getter,
	}
}

type signer struct {
	getter kv.Getter
	once   comfig.Once
}

func (s *signer) SignerConfig() *SignerConfig {
	return s.once.Do(func() interface{} {
		var cfg SignerConfig
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(s.getter, signerConfigKey)).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out signer config"))
		}
		return &cfg
	}).(*SignerConfig)
}
