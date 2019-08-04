package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Environment ... 環境変数
type Environment struct {
	Deploy    string `envconfig:"DEPLOY"                         required:"true"`
	ProjectID string `envconfig:"PROJECT_ID"                     required:"true"`
	// LocationID      string `envconfig:"LOCATION_ID"                    default:"asia-northeast1"`
	// ServiceID       string `envconfig:"SERVICE_ID"                     required:"true"`
	CredentialsPath string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS" required:"true"`
	MinLogSeverity  string `envconfig:"MIN_LOG_SEVERITY"               required:"true"`
}

// Get ... 環境変数を取得する
func (e *Environment) Get() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("", e)
	if err != nil {
		panic(err)
	}
}
