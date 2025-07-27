package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string , error) {
	home,err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(home, configFileName)

	return filePath, nil
}

func Read() (Config , error) {
	var c Config

	filePath,err := getConfigFilePath()
	if err != nil {
		return c , err
	}

	content,err := os.ReadFile(filePath)
	if err != nil {
		return c , err
	}

	err = json.Unmarshal(content, &c)
	if err !=  nil {
		return c , err
	}

	return c , nil
}

func write(c *Config) error {
	jsonData,err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath,err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath , jsonData , 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}