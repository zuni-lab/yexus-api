package config

import (
	"crypto/ecdsa"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/yexus-api/pkg/utils"
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
	CorsWhiteList []string

	// OpenAI
	OpenaiApiKey      string `validate:"min=1"`
	OpenaiAssistantId string `validate:"min=1"`
}

type RealtimeManagerConfig struct {
	RealtimeInterval time.Duration `validate:"min=1s"`

	PrivateKey      *ecdsa.PrivateKey
	ContractAddress common.Address
	Address         common.Address
}

type WorkerConfig struct {
	YieldMetricsSource string    `validate:"url"`
	YieldMetricsRunAt  time.Time `validate:"required"`
}

// ServerEnv combines all configurations
type ServerEnv struct {
	CommonConfig
	ServerConfig
	RealtimeManagerConfig
	WorkerConfig
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
	commonConfig := loadCommonConfig()

	serverConfig := loadServerConfig()

	managerConfig := loadRealtimeManagerConfig()

	workerConfig := loadWorkerConfig()

	Env = ServerEnv{
		CommonConfig:          commonConfig,
		ServerConfig:          serverConfig,
		RealtimeManagerConfig: managerConfig,
		WorkerConfig:          workerConfig,
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

	return ServerConfig{
		ApiHost:       os.Getenv("API_HOST"),
		Port:          port,
		CorsWhiteList: corsWhiteList,

		OpenaiApiKey:      os.Getenv("OPENAI_API_KEY"),
		OpenaiAssistantId: os.Getenv("OPENAI_ASSISTANT_ID"),
	}
}

func loadRealtimeManagerConfig() RealtimeManagerConfig {
	realtimeInterval := getEnvDuration("REALTIME_INTERVAL", "5s")

	priv := getEnvPrivateKey("PRIVATE_KEY")
	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Msg("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	contractAddress, err := utils.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	if err != nil {
		log.Fatal().Msgf("Error parsing CONTRACT_ADDRESS: %s", err)
	}

	return RealtimeManagerConfig{
		RealtimeInterval: realtimeInterval,
		ContractAddress:  contractAddress,
		PrivateKey:       priv,
		Address:          crypto.PubkeyToAddress(*publicKeyECDSA),
	}
}

func loadWorkerConfig() WorkerConfig {
	runAt, err := time.Parse("15:04", os.Getenv("YIELD_METRICS_RUN_AT"))
	if err != nil {
		log.Fatal().Msgf("Error parsing YIELD_METRICS_AT: %s", err)
	}

	return WorkerConfig{
		YieldMetricsSource: os.Getenv("YIELD_METRICS_SOURCE"),
		YieldMetricsRunAt:  runAt,
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

func getEnvPrivateKey(key string) *ecdsa.PrivateKey {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal().Msgf("Error parsing %s: %s", key, "empty")
	}
	privateKey, err := crypto.HexToECDSA(value)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", key, err)
	}
	return privateKey
}
