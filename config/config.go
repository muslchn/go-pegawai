package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for our application
type Config struct {
	Database DatabaseConfig
	App      AppConfig
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Name            string
	Charset         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
}

// AppConfig holds application-level configuration
type AppConfig struct {
	Environment string
	Port        int
}

// Global config instance
var GlobalConfig *Config

// GetConfig returns the global configuration instance
func GetConfig() *Config {
	if GlobalConfig == nil {
		panic("Configuration not loaded. Call LoadConfig() first.")
	}
	return GlobalConfig
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		// Don't return error if .env file doesn't exist in production
		if os.Getenv("APP_ENV") != "production" {
			fmt.Println("Warning: .env file not found, using environment variables")
		}
	}

	config := &Config{}

	// Load database configuration
	dbConfig, err := loadDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}
	config.Database = dbConfig

	// Load app configuration
	appConfig, err := loadAppConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load app config: %w", err)
	}
	config.App = appConfig

	// Set global config
	GlobalConfig = config

	return config, nil
}

// loadDatabaseConfig loads database configuration from environment variables
func loadDatabaseConfig() (DatabaseConfig, error) {
	config := DatabaseConfig{}

	// Required environment variables
	config.Host = getEnvOrDefault("DB_HOST", "localhost")
	config.Username = getEnvOrDefault("DB_USERNAME", "root")
	config.Password = os.Getenv("DB_PASSWORD") // Password should be explicitly set
	config.Name = getEnvOrDefault("DB_NAME", "pegawai_db")
	config.Charset = getEnvOrDefault("DB_CHARSET", "utf8mb4")

	// Parse port
	portStr := getEnvOrDefault("DB_PORT", "3306")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return config, fmt.Errorf("invalid DB_PORT value: %s", portStr)
	}
	config.Port = port

	// Parse connection pool settings
	maxOpenConnsStr := getEnvOrDefault("DB_MAX_OPEN_CONNS", "25")
	maxOpenConns, err := strconv.Atoi(maxOpenConnsStr)
	if err != nil {
		return config, fmt.Errorf("invalid DB_MAX_OPEN_CONNS value: %s", maxOpenConnsStr)
	}
	config.MaxOpenConns = maxOpenConns

	maxIdleConnsStr := getEnvOrDefault("DB_MAX_IDLE_CONNS", "10")
	maxIdleConns, err := strconv.Atoi(maxIdleConnsStr)
	if err != nil {
		return config, fmt.Errorf("invalid DB_MAX_IDLE_CONNS value: %s", maxIdleConnsStr)
	}
	config.MaxIdleConns = maxIdleConns

	connMaxLifetimeStr := getEnvOrDefault("DB_CONN_MAX_LIFETIME", "300")
	connMaxLifetime, err := strconv.Atoi(connMaxLifetimeStr)
	if err != nil {
		return config, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME value: %s", connMaxLifetimeStr)
	}
	config.ConnMaxLifetime = connMaxLifetime

	// Validate required fields
	if config.Name == "" {
		return config, fmt.Errorf("DB_NAME is required")
	}

	return config, nil
}

// loadAppConfig loads application configuration from environment variables
func loadAppConfig() (AppConfig, error) {
	config := AppConfig{}

	config.Environment = getEnvOrDefault("APP_ENV", "development")

	// Parse app port
	portStr := getEnvOrDefault("APP_PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return config, fmt.Errorf("invalid APP_PORT value: %s", portStr)
	}
	config.Port = port

	return config, nil
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetDSN returns the Data Source Name for database connection
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.Charset,
	)
}

// Validate validates the configuration
func (c *Config) Validate() error {
	// Validate database config
	if c.Database.Host == "" {
		return fmt.Errorf("database host is required")
	}
	if c.Database.Username == "" {
		return fmt.Errorf("database username is required")
	}
	if c.Database.Name == "" {
		return fmt.Errorf("database name is required")
	}
	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("database port must be between 1 and 65535")
	}

	// Validate app config
	if c.App.Port <= 0 || c.App.Port > 65535 {
		return fmt.Errorf("app port must be between 1 and 65535")
	}

	return nil
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.App.Environment == "production"
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.App.Environment == "development"
}
