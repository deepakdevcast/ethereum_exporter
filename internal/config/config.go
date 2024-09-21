package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MetricsPort string
var ClientInfo []struct {
	Client string `json:"client"`
	Url    string `json:"url"`
}

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalln("Environment not loaded!", err)
	}
	MetricsPort = os.Getenv("METRICS_PORT")
	if err := json.Unmarshal([]byte(os.Getenv("CLIENT_INFO")), &ClientInfo); err != nil {
		log.Fatalf("METRICS_INFO env has invalid json: %s", os.Getenv("METRICS_INFO"))
	}
}
