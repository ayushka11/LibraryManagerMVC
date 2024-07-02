package models

import (
	"os"
	"gopkg.in/yaml.v3"
)

type JWTKey struct {
	JWTKey string `yaml:"JWTSecretKey"`
}

func GetJWT() (string, error) {
	file, err := os.ReadFile("db.yaml")
	if err != nil {
		return "", err
	}
	var key JWTKey
	if err := yaml.Unmarshal(file, &key); err != nil {
		return "", err
	}
	return key.JWTKey, nil
}