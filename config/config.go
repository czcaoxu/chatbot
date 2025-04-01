package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	AIModelConfig AIModelConfig `json:"ai_model_config"`
}

type AIModelConfig struct {
	OpenAIKey string `json:"open_ai_key"`
	QwenKey   string `json:"qwen_key"`
	LlamaHost string `json:"llama_host"`
	LlamaPort string `json:"llama_port"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
