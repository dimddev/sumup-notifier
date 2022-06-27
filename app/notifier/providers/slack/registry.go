package slack

import (
	"encoding/json"
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type Providers map[string]providers.SlackProvider

type Registry struct {
	providers Providers
}

var _ providers.SlackRegister = (*Registry)(nil)

func NewRegistry(cfg providers.SlackConfiger, provider providers.SlackProvider) (*Registry, error) {
	r := new(Registry)
	r.providers = make(Providers)

	conf, errCfg := cfg.Decoder(cfg.SlackConfig())
	if errCfg != nil {
		return nil, errCfg
	}

	var slackConfigs config.SlackConfigs

	errDecode := json.Unmarshal(conf, &slackConfigs)
	if errDecode != nil {
		return nil, errDecode
	}

	for _, slackConfig := range slackConfigs {
		_init, errInit := provider.Init(slackConfig)
		if errInit != nil {
			return nil, errInit
		}

		init, ok := _init.(*Slack)
		if !ok {
			return nil, errors.New("slack doesn't exists")
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
