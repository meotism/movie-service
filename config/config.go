package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var config *Config

type Config struct {
	Stage       string `envconfig:"STAGE"`
	Debug       bool   `envconfig:"DEBUG"`
	AutoMigrate bool   `envconfig:"AUTO_MIGRATE"`
	Port        string `envconfig:"PORT"`
	CronJobFlag bool   `envconfig:"CRON_JOB_FLAG"`
	TimeZone    string `envconfig:"TIME_ZONE"`
	SwaggerHost string `envconfig:"SWAGGER_HOST"`

	Jwt struct {
		Key                   string `envconfig:"JWT_KEY"`
		RawTokenExpire        int    `envconfig:"JWT_TOKEN_EXPIRE"`
		RawRefreshTokenExpire int    `envconfig:"JWT_REFRESH_TOKEN_EXPIRE"`
		TokenExpire           time.Duration
		RefreshTokenExpire    time.Duration
	}

	Endpoints struct {
		DatadogAgentEndpoint string `envconfig:"DATADOG_AGENT_ENDPOINT"`
		HealthCheckEndPoint  string `envconfig:"HEALTH_CHECK_ENDPOINT"`
	}

	MySQL struct {
		Masters      []string `envconfig:"DB_MASTER_HOSTS"`
		Slaves       []string `envconfig:"DB_SLAVE_HOSTS"`
		DBName       string   `envconfig:"DB_NAME"`
		User         string   `envconfig:"DB_USER"`
		Pass         string   `envconfig:"DB_PASS"`
		MaxIdleConns int      `envconfig:"DB_MAX_IDLE_CONNECTIONS"`
		MaxOpenConns int      `envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	}

	Redis struct {
		Host    string `envconfig:"REDIS_HOST"`
		Port    string `envconfig:"REDIS_PORT"`
		DB      int    `envconfig:"REDIS_DB"`
		User    string `envconfig:"REDIS_USER"`
		Pass    string `envconfig:"REDIS_PASS"`
		Timeout int    `envconfig:"REDIS_TIMEOUT"`
	}
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	config = &Config{}

	err = envconfig.Process("", config)

	if err != nil {
		panic(fmt.Sprintf("Failed to decode config env: %v", err))
	}

	if len(config.Port) == 0 {
		config.Port = "9000"
	}

	if config.MySQL.MaxIdleConns == 0 {
		config.MySQL.MaxIdleConns = 10
	}

	if config.MySQL.MaxOpenConns == 0 {
		config.MySQL.MaxOpenConns = 10
	}

	config.Jwt.TokenExpire = time.Duration(config.Jwt.RawTokenExpire) * time.Hour
	config.Jwt.RefreshTokenExpire = time.Duration(config.Jwt.RawRefreshTokenExpire) * time.Hour
}

// GetConfig .
func GetConfig() *Config {
	return config
}

// type MysqlEnvConfig struct {
// 	Host     string
// 	DB       string
// 	User     string
// 	Port     string
// 	Password string
// 	URI      string
// }

// type ServerEnvConfig struct {
// 	Port string
// }

// type Pagination struct {
// 	MovieRowPerPage int
// }

// func LoadFileEnv() {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatal("Error loading .env file: ", err)
// 	}
// }

// func GetSecretKey() string {
// 	LoadFileEnv()
// 	return os.Getenv("SECRET_KEY")
// }

// func GetMysqlEnvConfig() *MysqlEnvConfig {
// 	LoadFileEnv()
// 	var mysql MysqlEnvConfig
// 	mysql.Host = os.Getenv("MYSQL_HOST")
// 	mysql.User = os.Getenv("MYSQL_USER")
// 	mysql.Password = os.Getenv("MYSQL_PASSWORD")
// 	mysql.DB = os.Getenv("MYSQL_DB")
// 	mysql.Port = os.Getenv("MYSQL_PORT")

// 	return &mysql
// }

// func GetServerConfig() *ServerEnvConfig {
// 	LoadFileEnv()
// 	var server ServerEnvConfig
// 	server.Port = os.Getenv("SERVER_PORT")

// 	return &server
// }

// func GetPagination() *Pagination {
// 	LoadFileEnv()
// 	var pagi Pagination
// 	rowPerPage, err := strconv.Atoi(os.Getenv("MOVIE_ROW_PER_PAGE"))
// 	pagi.MovieRowPerPage = rowPerPage

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &pagi
// }
