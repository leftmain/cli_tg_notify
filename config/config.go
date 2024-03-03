package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type SendConfig struct {
	TelegramApiToken string `json:"token"`
	ChatId           int64  `json:"chat_id"`
}

const DEFAULT_CONFIG_NAME = ".send2telegram_config"

func getDefaultConfigPath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return DEFAULT_CONFIG_NAME
	}

	return filepath.Join(homeDir, DEFAULT_CONFIG_NAME)
}

func LoadConfig(configPath string) (SendConfig, error) {
	var config SendConfig

	if configPath == "" {
		configPath = getDefaultConfigPath()
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Cannot open config file: %s\n", configPath)
		return config, err
	}
	defer configFile.Close()

	configBytes, _ := io.ReadAll(configFile)
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Printf("Cannot parse config file: %s\n", configPath)
		return config, err
	}

	return config, nil
}

func SaveConfig(config SendConfig, configPath string) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
