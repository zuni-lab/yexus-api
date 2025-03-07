package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Common configuration shared between server and indexer
type CommonConfig struct {
	Env     string
	AppName string `validate:"min=1"`
	IsProd  bool
	IsDev   bool
	IsTest  bool

	// Database
	PostgresUrl  string `validate:"url"`
	MigrationUrl string `validate:"url"`

	// Monitoring
	OpenObserveEndpoint   string `validate:"url"`
	OpenObserveCredential string `validate:"min=1"`

	// EVM
	AlchemyUrl string `validate:"url"`
}

// Server configuration
type ServerConfig struct {
	ApiHost       string `validate:"min=1"`
	Port          string `validate:"number"`
	JwtSecret     string `validate:"min=10"`
	CorsWhiteList []string

	// OpenAI
	OpenaiApiKey      string `validate:"min=1"`
	OpenaiAssistantId string `validate:"min=1"`

	// Realtime Manager
	RealtimeInterval      time.Duration `validate:"min=1s"`
	RealtimeMinBlockRange uint64        `validate:"min=1"`
	RealtimeMaxBlockRange uint64        `validate:"min=1"`

	ContractAddress string `validate:"eth_addr"`
}

// Indexer-specific configuration
type IndexerConfig struct {
	ChunkSize     uint64        `validate:"min=1"`
	Concurrency   int           `validate:"min=1"`
	StartBlock    uint64        `validate:"min=1"`
	FetchInterval time.Duration `validate:"min=1s"`
}

// ServerEnv combines all configurations
type ServerEnv struct {
	CommonConfig
	ServerConfig
	IndexerConfig
}

var Env ServerEnv

func LoadEnvWithPath(path string) {
	err := godotenv.Load(os.ExpandEnv(path))
	if err != nil {
		log.Fatal().Msgf("Error loading %s file: %s", path, err)
	}
	loadEnv()
}

func LoadEnv() {
	if os.Getenv("ENV") == "" {
		_ = os.Setenv("ENV", "development")
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			log.Fatal().Msgf("Error loading .env file: %s", err)
		}
	} else if os.Getenv("ENV") == "test" {
		err := godotenv.Load(os.ExpandEnv(".env.test"))
		if err != nil {
			log.Fatal().Msgf("Error loading .env.test file: %s", err)
		}
	}
	loadEnv()
}

func loadEnv() {
	common := loadCommonConfig()

	server := loadServerConfig()

	indexer := loadIndexerConfig()

	Env = ServerEnv{
		CommonConfig:  common,
		ServerConfig:  server,
		IndexerConfig: indexer,
	}

	validate := validator.New()
	if err := validate.Struct(Env); err != nil {
		log.Fatal().Msgf("Error validating env: %s", err)
	}
}

func loadCommonConfig() CommonConfig {
	env := os.Getenv("ENV")
	return CommonConfig{
		Env:     env,
		AppName: os.Getenv("APP_NAME"),
		IsProd:  env == "production",
		IsDev:   env == "development" || env == "",
		IsTest:  env == "test",

		PostgresUrl:  os.Getenv("POSTGRES_URL"),
		MigrationUrl: os.Getenv("MIGRATION_URL"),

		OpenObserveEndpoint:   os.Getenv("OPENOBSERVE_ENDPOINT"),
		OpenObserveCredential: os.Getenv("OPENOBSERVE_CREDENTIAL"),

		AlchemyUrl: os.Getenv("ALCHEMY_URL"),
	}
}

func loadServerConfig() ServerConfig {
	// Load CORS whitelist
	rawCORSWhiteList := os.Getenv("CORS_WHITE_LIST")
	corsWhiteList := []string{"http://localhost:3000"}
	if rawCORSWhiteList != "" {
		corsWhiteList = strings.Split(rawCORSWhiteList, ",")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	realtimeInterval := getEnvDuration("REALTIME_INTERVAL", "1s")

	return ServerConfig{
		ApiHost:       os.Getenv("API_HOST"),
		Port:          port,
		JwtSecret:     os.Getenv("JWT_SECRET"),
		CorsWhiteList: corsWhiteList,

		OpenaiApiKey:      os.Getenv("OPENAI_API_KEY"),
		OpenaiAssistantId: os.Getenv("OPENAI_ASSISTANT_ID"),

		RealtimeInterval:      realtimeInterval,
		RealtimeMinBlockRange: getEnvUint64("REALTIME_MIN_BLOCK_RANGE", 5),
		RealtimeMaxBlockRange: getEnvUint64("REALTIME_MAX_BLOCK_RANGE", 25),

		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
	}
}

func loadIndexerConfig() IndexerConfig {
	return IndexerConfig{
		ChunkSize:     getEnvUint64("INDEXER_CHUNK_SIZE", 900),
		Concurrency:   getEnvInt("INDEXER_CONCURRENCY", 10),
		StartBlock:    getEnvUint64("INDEXER_START_BLOCK", 1),
		FetchInterval: getEnvDuration("INDEXER_FETCH_INTERVAL", "1s"),
	}
}

// Helper functions for environment variable parsing
func getEnvDuration(key, defaultValue string) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", key, err)
	}
	return duration
}

func getEnvUint64(key string, defaultValue uint64) uint64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsed, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", key, err)
	}
	return parsed
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", key, err)
	}
	return parsed
}
