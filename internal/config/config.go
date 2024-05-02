package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

const DefaultConfigPath = "\\config\\config.json"

type Config struct {
	OSType            int
	CollectorInterval time.Duration `json:"interval"`
	CollectCPU        bool          `json:"cpu"`
	CollectSpace      bool          `json:"space"`
	GRPCConfig        `json:"grpc"`
}

type GRPCConfig struct {
	Port    string        `json:"port"`
	Timeout time.Duration `json:"timeout"`
}

func LoadConfig() (*Config, error) {
	configPath := fetchConfigPath()
	if configPath == "" {
		return nil, fmt.Errorf("path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", configPath)
	}

	var cfg Config
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("reading %s error %e", configPath, err)
	}

	err = json.Unmarshal(file, cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshall error %e", err)
	}
	return &cfg, nil
}

func fetchConfigPath() string {
	var res string
	res = DefaultConfigPath
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("SYSTEM_STATS_CONFIG_PATH")
	}

	return res
}
