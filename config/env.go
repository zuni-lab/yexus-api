package config

import (
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type ServerEnv struct {
	Env           string
	CorsWhiteList []string

	AppName   string `validate:"min=1"`
	ApiHost   string `validate:"min=1"`
	JwtSecret string `validate:"min=10"`

	Port string `validate:"number"`

	IsProd    bool
	IsStaging bool
	IsDev     bool
	IsTest    bool

	OpenObserveEndpoint   string `validate:"url"`
	OpenObserveCredential string `validate:"min=1"`

	PostgresUrl  string `validate:"url"`
	MigrationUrl string `validate:"url"`

	AlchemyUrl string `validate:"url"`
}

var Env ServerEnv

func LoadEnvWithPath(path string) {
	err := godotenv.Load(os.ExpandEnv(path))
	if err != nil {
		log.Fatalf("Error loading %s file: %s", path, err)
	}

	loadEnv()
}

func LoadEnv() {
	if os.Getenv("ENV") == "" {
		_ = os.Setenv("ENV", "development")
		err := godotenv.Load(os.ExpandEnv(".env"))
		if err != nil {
			log.Fatalln("Error loading .env file: ", err)
		}
	} else if os.Getenv("ENV") == "test" {
		err := godotenv.Load(os.ExpandEnv(".env.test"))
		if err != nil {
			log.Fatalln("Error loading .env.test file: ", err)
		}
	}

	loadEnv()
}

func loadEnv() {
	rawCORSWhiteList := os.Getenv("CORS_WHITE_LIST")
	var corsWhiteList []string
	if rawCORSWhiteList == "" {
		corsWhiteList = []string{
			"http://localhost:3000",
		}
	} else {
		corsWhiteList = strings.Split(rawCORSWhiteList, ",")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}

	Env = ServerEnv{
		Env:           os.Getenv("ENV"),
		CorsWhiteList: corsWhiteList,

		AppName:   os.Getenv("APP_NAME"),
		ApiHost:   os.Getenv("API_HOST"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Port:      port,

		OpenObserveEndpoint:   os.Getenv("OPENOBSERVE_ENDPOINT"),
		OpenObserveCredential: os.Getenv("OPENOBSERVE_CREDENTIAL"),

		PostgresUrl:  os.Getenv("POSTGRES_URL"),
		MigrationUrl: os.Getenv("MIGRATION_URL"),

		AlchemyUrl: os.Getenv("ALCHEMY_URL"),
	}

	validate := validator.New()
	err := validate.Struct(Env)

	if err != nil {
		panic(err)
	}

	Env.IsProd = Env.Env == "production"
	Env.IsStaging = Env.Env == "staging"
	Env.IsDev = Env.Env == "development" || len(Env.Env) == 0
	Env.IsTest = Env.Env == "test"
}
