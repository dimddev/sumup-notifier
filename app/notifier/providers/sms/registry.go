package sms

import (
	"encoding/json"
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type Providers map[string]providers.SMSProvider

type Registry struct {
	providers Providers
}

var _ providers.SMSRegister = (*Registry)(nil)

func NewRegistry(cfg providers.SMSConfiger, provider providers.SMSProvider) (*Registry, error) {
	r := new(Registry)
	r.providers = make(Providers)

	conf, errCfg := cfg.Decoder(cfg.SMSConfig())
	if errCfg != nil {
		return nil, errCfg
	}

	var smsConfigs config.SMSConfigs

	errDecode := json.Unmarshal(conf, &smsConfigs)
	if errDecode != nil {
		return nil, errDecode
	}

	for _, smsConfig := range smsConfigs {
		_init, errInit := provider.Init(smsConfig)
		if errInit != nil {
			return nil, errInit
		}

		init, ok := _init.(*SMS)
		if !ok {
			return nil, errors.New("sms doesn't exists")
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
