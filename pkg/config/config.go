package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`

	AuthServerPort             string `mapstructure:"AUTH_SERVER_PORT"`
	ManagementSystemServerPort string `mapstructure:"MANAGEMENT_SYSTEM_SERVER_PORT"`
	GrpcGatewayPort            string `mapstructure:"GRPC_GATEWAY_PORT"`
	MonitorServerPort          string `mapstructure:"MONITOR_SERVER_PORT"`
	MailServerPort             string `mapstructure:"MAIL_SERVER_PORT"`

	TokenInternal string `mapstructure:"TOKEN_INTERNAL"`

	PGDatabaseHost     string `mapstructure:"PG_DATABASE_HOST"`
	PGDatabaseUser     string `mapstructure:"PG_DATABASE_USER"`
	PGDatabasePassword string `mapstructure:"PG_DATABASE_PASSWORD"`
	PGDatabaseDBName   string `mapstructure:"PG_DATABASE_DBNAME"`
	PGDatabasePort     string `mapstructure:"PG_DATABASE_PORT"`

	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDB       string `mapstructure:"REDIS_DB"`

	PathPrivateKey string `mapstructure:"PATH_PRIVATE_KEY"`
	PathPublicKey  string `mapstructure:"PATH_PUBLIC_KEY"`

	MonitorBrokerAddress   string `mapstructure:"MONITOR_BROKER_ADDRESS"`
	MonitorTopic           string `mapstructure:"MONITOR_TOPIC"`
	MonitorResultsTopic    string `mapstructure:"MONITOR_RESULTS_TOPIC"`
	MonitorConsumerGroupID string `mapstructure:"MONITOR_CONSUMER_GROUP_ID"`

	MonitorDurationMinute   int `mapstructure:"MONITOR_DURATION_MINUTE"`
	MaxConurrentPingServers int `mapstructure:"MAX_CONCURRENT_PING_SERVERS"`

	KafkaBrokerAddress          string `mapstructure:"KAFKA_BROKER_ADDRESS"`
	KafkaTopic                  string `mapstructure:"KAFKA_TOPIC"`
	ResultsKafkaTopic           string `mapstructure:"RESULTS_KAFKA_TOPIC"`
	KafkaConsumerGroupId        string `mapstructure:"KAFKA_CONSUMER_GROUP_ID"`
	KafkaResultsConsumerGroupId string `mapstructure:"KAFKA_RESULTS_CONSUMER_GROUP_ID"`

	LogFilename   string `mapstructure:"LOG_FILE_NAME"`
	LogMaxSize    int64  `mapstructure:"LOG_MAX_SIZE"`
	LogMaxBackups int64  `mapstructure:"LOG_MAX_BACKUPS"`
	LogMaxAge     int64  `mapstructure:"LOG_MAX_AGE"`
}

func NewConfig(path string, name string) (config *Config, err error) {
	cfg := &Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	if err = viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return cfg, err
}
