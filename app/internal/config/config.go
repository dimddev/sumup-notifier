package config

import (
	"encoding/base64"
	"os"

	"github.com/joho/godotenv"
)

type Configer interface {
	Decoder(string) ([]byte, error)
	MailConfig() string
	SlackConfig() string
	SMSConfig() string
	DBConfig() string
}

type Config struct {
	MailServers    string `env:"MAIL_SERVERS"`
	SlackEndPoints string `env:"SLACK_ENDPOINTS"`
	SMSServices    string `env:"SMS_SERVICES"`
	DataBase       string `env:"DATABASE"`
}

var _ Configer = (*Config)(nil)

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env.local")
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	cfg.MailServers = os.Getenv("MAIL_SERVERS")
	cfg.SlackEndPoints = os.Getenv("SLACK_ENDPOINTS")
	cfg.SMSServices = os.Getenv("SMS_SERVICES")
	cfg.DataBase = os.Getenv("DATABASE")

	return cfg, nil
}

func (c *Config) Decoder(data string) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dst, []byte(data))
	if err != nil {
		return nil, err
	}

	dst = dst[:n]

	return dst, nil
}

func (c *Config) MailConfig() string {
	return c.MailServers
}

func (c *Config) SlackConfig() string {
	return c.SlackEndPoints
}

func (c *Config) SMSConfig() string {
	return c.SMSServices
}

func (c *Config) DBConfig() string {
	return c.DataBase
}

type DataBaseConfig struct {
	Host     string `json:"host"`
	User     string `json:"username"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	Port     string `json:"port"`
	SSL      string `json:"sslmode"`
}

type MailConfig struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
	SSL    bool   `json:"ssl"`
	Name   string `json:"name"`
}

type MailConfigs []MailConfig

type SlackConfig struct {
	URL    string `json:"url"`
	APIKey string `json:"apikey"`
	Name   string `json:"name"`
}

type SlackConfigs []SlackConfig

type SMSConfig struct {
	Server string `json:"server"`
	APIKey string `json:"apikey"`
	Name   string `json:"name"`
}

type SMSConfigs []SMSConfig
