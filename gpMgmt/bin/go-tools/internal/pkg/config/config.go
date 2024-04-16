package config

import (
	"github.com/spf13/viper"
)

type Config interface {
	SetName(string)
	Load(...string) error

	GetDatabaseConfig() Database
	GetArtefactConfig() Artefact
	GetInfraConfig() Infra
}

func New() Config {
	parser := viper.New()
	parser.SetConfigName("gp.conf") // default config name
	parser.SetConfigType("json")
	return &appConfig{
		parser: parser,
	}
}

type appConfig struct {
	parser *viper.Viper

	// TODO : embed hub Config

	Database Database `json:"database"`
	Artefact Artefact `json:"artefact"`
	Infra    Infra    `json:"infra"`
}

func (conf *appConfig) SetName(confgiName string) {
	conf.parser.SetConfigName(confgiName)
}

func (conf *appConfig) Load(configPaths ...string) error {
	for _, configPath := range configPaths {
		conf.parser.AddConfigPath(configPath)
	}
	conf.setDefaults()
	err := conf.parser.ReadInConfig()
	if err != nil {
		return err
	}
	err = conf.parser.Unmarshal(conf)
	if err != nil {
		return err
	}
	return nil
}

func (conf *appConfig) GetDatabaseConfig() Database {
	return conf.Database
}

func (conf *appConfig) GetArtefactConfig() Artefact {
	return conf.Artefact
}

func (conf *appConfig) GetInfraConfig() Infra {
	return conf.Infra
}

func (conf *appConfig) setDefaults() {
	conf.parser.SetDefault("Infra.RequestPort", 4506)
	conf.parser.SetDefault("Infra.PublishPort", 4505)
}
