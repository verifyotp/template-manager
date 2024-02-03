package config

import (
	"fmt"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	conf *sync.Map
}

func New() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading .env file")
	}

	conf := new(sync.Map)
	return &Config{conf: conf}
}

func (c *Config) SetEnv(key string, value any) *Config {
	c.conf.Store(key, value)
	return c
}

func (c *Config) GetEnv(key string) any {
	if value, ok := c.conf.Load(key); ok {
		return value
	}
	return nil
}

func (c *Config) GetString(key string) string {
	return fmt.Sprint(c.GetEnv(key))
}

func (c *Config) GetInt(key string) int {
	value, ok := c.conf.Load(key)
	if !ok {
		return 0
	}
	if intValue, ok := value.(int); ok {
		return intValue
	}
	return 0
}

func (c *Config) GetFloat64(key string) float64 {
	value, ok := c.conf.Load(key)
	if !ok {
		return 0
	}
	if intValue, ok := value.(float64); ok {
		return intValue
	}
	return 0
}
