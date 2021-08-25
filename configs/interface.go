package configs

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
)

type Secrets struct {
	DatabaseSecrets
}

func Configure() (*Configuration, error) {
	viperConfig := viper.GetViper()
	viperConfig.AutomaticEnv()

	config := Configuration{}
	config.GetConfiguration(*viperConfig)
	if err := config.GetSecrets(); err != nil {
		return nil, err
	}
	return &config, nil
}

type Configuration struct {
	ServerConfiguration
	DatabaseConfiguration
	Secrets
}

func (c *Configuration) GetConfiguration(v viper.Viper) {
	c.DatabaseConfiguration = GetDatabaseConfig(v)
	c.ServerConfiguration = GetServerConfig(v)
}

func (c *Configuration) GetSecrets() error {
	var secretsMap Secrets

	secrets, err := ioutil.ReadFile(c.SrvSecretsFile)
	if err != nil { return err }

	if err := json.Unmarshal(secrets, &secretsMap); err != nil {
		return err
	}

	c.Secrets = secretsMap
	return nil
}
