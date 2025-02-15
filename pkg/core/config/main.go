package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv            string `mapstructure:"APP_ENV"`
	AppPort           string `mapstructure:"APP_PORT"`
	DBType            string `mapstructure:"DB_TYPE"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            int    `mapstructure:"DB_PORT"`
	DBName            string `mapstructure:"DB_NAME"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBSSLMode         string `mapstructure:"DB_SSLMODE"`
	Migrate           bool   `mapstructure:"MIGRATE"`
	Seed              bool   `mapstructure:"SEED"`
	RedisExp          int    `mapstructure:"REDIS_EXP"`
	RedisURI          string `mapstructure:"REDIS_URI"`
	RedisPassword     string `mapstructure:"REDIS_PASSWORD"`
	RedisDB           int    `mapstructure:"REDIS_DB"`
	IsRedis           bool   `mapstructure:"IS_REDIS"`
	RateLimitEnabled  bool   `mapstructure:"RATE_LIMIT_ENABLED"`
	RateLimit         int    `mapstructure:"RATE_LIMIT"`
	RateLimitDuration string `mapstructure:"RATE_LIMIT_DURATION"`
	JwtSecretKey      string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpiration     string `mapstructure:"JWT_EXPIRATION"`
}

var (
	GlobalConfig *Config
	configMutex  sync.Mutex
)

func LoadConfig() error {
	// Lock the mutex to ensure thread safety during configuration loading
	configMutex.Lock()
	defer configMutex.Unlock()
	configFilePath := ".env"
	viper.SetConfigFile(configFilePath)
	// Initialize viper
	// viper.SetConfigName("prod-config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".") // Look for the config.yaml file in the current directory
	// viper.AutomaticEnv()
	// Enable automatic configuration watching
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	viper.WatchConfig()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the configuration into a Config struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Set default values for configuration fields
	setDefaultValues(&cfg)

	// Set the global configuration variable
	GlobalConfig = &cfg

	return nil
}

// setDefaultValues sets default values for configuration fields
func setDefaultValues(cfg *Config) {
	if cfg.RedisDB == 0 {
		cfg.RedisDB = 0 // Default Redis database
	}
	// Add default values for other configuration fields as needed
}
