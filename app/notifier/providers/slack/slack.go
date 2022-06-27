package slack

import (
	"errors"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

type Slack struct {
	URL       string `json:"url"`
	APIKey    string `json:"apikey"`
	Signature string `json:"name"`

	logging providers.Logger
}

var _ providers.SlackProvider = (*Slack)(nil)

func NewSlack(logging providers.Logger) *Slack {
	return &Slack{logging: logging}
}

func (e *Slack) Send(from, to, message string) error {
	e.logging.Infof("Sending a slack message with content: <%s> from: <%s> to channel: <%s>", message, from, to)
	return nil
}

func (e *Slack) Name() string {
	return e.Signature
}

func (e *Slack) Init(slackConf interface{}) (interface{}, error) {
	slack, ok := slackConf.(config.SlackConfig)
	if !ok {
		return nil, errors.New("slack config doesn't exists")
	}

	slackCfg := &Slack{
		URL:       slack.URL,
		APIKey:    slack.APIKey,
		Signature: slack.Name,
		logging:   e.logging,
	}

	if slack.URL == "" && slack.APIKey == "" {
		return nil, errors.New("invalid configuration")
	}

	return slackCfg, nil
}

func (e *Slack) Connect() error {
	e.logging.Infof("Connect to Slack server %s", e.URL)

	return nil
}

func (e *Slack) Process(options interface{}) error {
	opt, ok := options.(*Options)
	if !ok {
		return errors.New("options doesn't exists")
	}

	if errConnect := e.Connect(); errConnect != nil {
		return errConnect
	}

	return e.Send("System", opt.Channel, opt.Message)
}
