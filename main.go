package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	Token string
}

func GetConfig() func() *Config {
	var cfg *Config

	return func() *Config {
		if cfg != nil {
			return cfg
		}

		if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
			slog.Error("не удалось загрузить .env", "error", err)
			return nil
		}

		cfg = &Config{
			Port:  os.Getenv("PORT"),
			Token: os.Getenv("TOKEN"),
		}

		if cfg.Port == "" {
			panic("PORT не задан")
		}

		if cfg.Token == "" {
			panic("PORT не задан")
		}

		return cfg
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	config := GetConfig()
	_, err := w.Write([]byte("Hello world v2, token: " + config().Token))

	if err != nil {
		slog.Error("Error ocurred while writing a response")
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ready"))
}

func main() {
	config := GetConfig()
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(helloWorldHandler))
	mux.Handle("/health", http.HandlerFunc(healthHandler))
	mux.Handle("/ready", http.HandlerFunc(readyHandler))

	slog.Info("Server start at http://localhost:" + config().Port)

	err := http.ListenAndServe(":"+config().Port, mux)

	if err != nil {
		panic(err)
	}
}
