package email

import (
	"encoding/json"
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type Providers map[string]providers.EmailProvider

type Registry struct {
	providers Providers
}

var _ providers.EmailRegister = (*Registry)(nil)

func NewRegistry(cfg providers.EmailConfiger, provider providers.EmailProvider) (*Registry, error) {
	r := new(Registry)
	r.providers = make(Providers)

	conf, errCfg := cfg.Decoder(cfg.MailConfig())
	if errCfg != nil {
		return nil, errCfg
	}

	var mailConfigs config.MailConfigs

	errDecode := json.Unmarshal(conf, &mailConfigs)
	if errDecode != nil {
		return nil, errDecode
	}

	for _, mailConfig := range mailConfigs {
		_init, errInit := provider.Init(mailConfig)
		if errInit != nil {
			return nil, errInit
		}

		init, ok := _init.(*Email)
		if !ok {
			return nil, errors.New("email doesnt exists")
		}

		r.providers[init.Name()] = init
	}

	return r, nil
}

func (r *Registry) Get(key string) (providers.BaseProvider, error) {
	provider, ok := r.providers[key]
	if !ok {
		return nil, errors.New("invalid provider")
	}

	return provider, nil
}
