package config

import (
	"encoding/json"
	"os"
)

type Config struct {
    Blockchain BlockchainConfig `json:"blockchain"`
    LogLevel   string           `json:"log_level"`
}

type BlockchainConfig struct {
    NodeURL string `json:"node_url"`
}

func LoadConfig() (*Config, error) {
    file, err := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := &Config{}
    if err := decoder.Decode(config); err != nil {
        return nil, err
    }

    return config, nil
}
