package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	CORS     CORSConfig     `mapstructure:"cors"`
	Upload   UploadConfig   `mapstructure:"upload"`
	SMTP     SMTPConfig     `mapstructure:"smtp"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `mapstructure:"postgres"`
	Mongo    MongoConfig    `mapstructure:"mongo"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type PostgresConfig struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DBName         string `mapstructure:"dbname"`
	SSLMode        string `mapstructure:"sslmode"`
	MaxConnections int    `mapstructure:"max_connections"`
	MaxIdleConns   int    `mapstructure:"max_idle_connections"`
}

func (p PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode,
	)
}

type MongoConfig struct {
	URI    string `mapstructure:"uri"`
	DBName string `mapstructure:"dbname"`
}

type RedisConfig struct {
	URL      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret        string `mapstructure:"secret"`
	AccessExpiry  string `mapstructure:"access_expiry"`
	RefreshExpiry string `mapstructure:"refresh_expiry"`
}

type CORSConfig struct {
	AllowOrigins string `mapstructure:"allow_origins"`
}

type UploadConfig struct {
	MaxSize           int      `mapstructure:"max_size"`
	Path              string   `mapstructure:"path"`
	AllowedExtensions []string `mapstructure:"allowed_extensions"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Environment variable overrides
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Defaults
	viper.SetDefault("app.name", "zplus-calendar")
	viper.SetDefault("app.env", "development")
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("database.postgres.host", "localhost")
	viper.SetDefault("database.postgres.port", "5432")
	viper.SetDefault("database.postgres.sslmode", "disable")
	viper.SetDefault("database.postgres.max_connections", 25)
	viper.SetDefault("database.postgres.max_idle_connections", 5)
	viper.SetDefault("database.redis.db", 0)
	viper.SetDefault("jwt.access_expiry", "15m")
	viper.SetDefault("jwt.refresh_expiry", "168h")
	viper.SetDefault("cors.allow_origins", "http://localhost:3000")
	viper.SetDefault("upload.max_size", 10485760)
	viper.SetDefault("upload.path", "./uploads")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config: %w", err)
		}
		// Config file not found, using defaults + env vars
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &cfg, nil
}
