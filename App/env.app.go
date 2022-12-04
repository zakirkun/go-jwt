package app

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		L.Panic("Failed load env var")
	}
	L.Info("Load env avr")
}
