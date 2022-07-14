package config

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeshaw/envdecode"
	"gopkg.in/yaml.v3"

	"broadcaster/internal/config/fields"
)

const (
	KindEVM    = "evm"
	KindCosmos = "cosmos"
)

type ChainsConfig []*ChainConfig

func (chainsConfig *ChainsConfig) Get(chainId int64) (*ChainConfig, error) {
	for _, chainConfig := range *chainsConfig {
		if chainConfig.ID == chainId {
			return chainConfig, nil
		}
	}
	return nil, fmt.Errorf("chain %d no found chain", chainId)
}

func (chainsConfig *ChainsConfig) UnmarshalYAML(value *yaml.Node) error {
	var err error
	for _, chainValue := range value.Content {
		chainConfig := ChainConfig{}
		err = defaults.Set(&chainConfig)
		if err != nil {
			return err
		}
		err = chainValue.Decode(&chainConfig)
		if err != nil {
			return err
		}
		key := os.Getenv(fmt.Sprintf("BROADCASTER_CHAIN_%v_KEY", chainConfig.ID))
		if len(key) > 0 {
			if err := chainConfig.Key.Decode(key); err != nil {
				return err
			}
		}
		if whitelist := os.Getenv(fmt.Sprintf("BROADCASTER_CHAIN_%v_WHITELIST", chainConfig.ID)); len(whitelist) > 0 {
			for _, address := range strings.Split(whitelist, ";") {
				chainConfig.Whitelist = append(chainConfig.Whitelist, common.HexToAddress(address))
			}
		}
		*chainsConfig = append(*chainsConfig, &chainConfig)
	}
	return nil
}

type LoggerConfig struct {
	Level  fields.LoggerLevel  `default:"info" yaml:"level" env:"BROADCASTER_LOGGER_LEVEL"`
	Format fields.LoggerFormat `default:"text" yaml:"format" env:"BROADCASTER_LOGGER_FORMAT"`
}

type AdvisorMockConfig struct {
	Result           string `yaml:"result"`
	GasPrice         int64  `yaml:"gas_price"`
	GasPriceE1559    bool   `yaml:"gas_price_e1559"`
	GasPricePriority int64  `yaml:"gas_price_priority"`
}

type AdvisorConfig struct {
	Disable bool               `yaml:"disable"`
	Url     string             `yaml:"url" env:"BROADCASTER_ADVISOR_URL"`
	Mock    *AdvisorMockConfig `yaml:"mock"`
}

type EstimateConfig struct {
	Disable bool `yaml:"disable"`
}

type Config struct {
	ConfigFile  string                 `yaml:"-"`
	Chains      ChainsConfig           `yaml:"chains"`
	Server      Server                 `yaml:"server" `
	Advisor     AdvisorConfig          `yaml:"advisor"`
	Estimate    EstimateConfig         `yaml:"estimate"`
	Store       Store                  `yaml:"store" `
	Recovery    Recovery               `yaml:"recovery"`
	Logger      LoggerConfig           `yaml:"logger"`
	Queue       QueueConfig            `yaml:"queue"`
	Environment fields.EnvironmentName `default:"dev" env:"BROADCASTER_ENV"`
}

type QueueConfig struct {
	Size uint `yaml:"size" env:"BROADCASTER_QUEUE_SIZE" default:"100"`
}

type WhitelistAddress []common.Address

func (wa WhitelistAddress) Accepted(target common.Address) bool {
	if len(wa) == 0 {
		return true
	}
	for _, address := range wa {
		if bytes.Equal(address.Bytes(), target.Bytes()) {
			return true
		}
	}

	return false
}

type ChainConfig struct {
	Whitelist WhitelistAddress  `yaml:"whitelist"`
	ID        int64             `yaml:"id"`
	Terra     TerraConfig       `yaml:"terra"`
	Kind      string            `yaml:"kind"`
	GasLimit  uint64            `yaml:"gas_limit"`
	Endpoints []string          `yaml:"endpoints"`
	Key       fields.PrivateKey `yaml:"key"`
	Retry     RetryConfig       `yaml:"retry"`
	Confirm   bool              `yaml:"confirm" default:"false"`
}

type TerraConfig struct {
	ID   string `yaml:"id"`
	Auth string `yaml:"auth" env:"BROADCASTER_TERRA_AUTH"`
}

type RetryConfig struct {
	Attempts uint                  `yaml:"attempts"`
	Delay    fields.DurationConfig `yaml:"delay"`
}

func (rc *RetryConfig) SetDefaults() {
	rc.Attempts = 20
	rc.Delay.Set(15 * time.Second)
}

type Server struct {
	Address fields.HostPortConfig `yaml:"address" env:"BROADCASTER_SERVER_ADDRESS" default:"127.0.0.1:4040"`
}

type ServerAuth struct {
	User string `yaml:"user" env:"BROADCASTER_SERVER_AUTH_USER"`
	Pass string `yaml:"pass" env:"BROADCASTER_SERVER_AUTH_PASS"`
}

type Store struct {
	DSN string `yaml:"dsn" env:"BROADCASTER_STORE_DSN" required:"true"`
}

type Recovery struct {
	Enabled bool `yaml:"enabled" default:"true" env:"BROADCASTER_RECOVERY"`
}

func Load(path string) (*Config, error) {
	config := &Config{
		ConfigFile: path,
	}

	file, err := os.Open(config.ConfigFile)
	if err == nil {
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		if err = yaml.NewDecoder(file).Decode(&config); err != nil && err != io.EOF {
			return nil, err
		}
	} else {
		return nil, err
	}

	if err = envdecode.Decode(config); err != nil && err != envdecode.ErrNoTargetFieldsAreSet {
		return nil, err
	}

	return config, nil
}

type Flags struct {
	ValueConfig string `docopt:"--config"`
}
