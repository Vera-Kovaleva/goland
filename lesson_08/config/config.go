package config

import (
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Port           int64
	DbUrl          string
	JaegerUrl      string
	SentryUrl      string
	KafkaBrokerUrl string
	SomeAppId      string
	SomeAppKey     string
}

func GetConfig() (Config, error) {
	port := getIntValue("PORT", "port", 8080)
	dbUrl := getStringValue("DB_URL", "db-url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	jaegerUrl := getStringValue("JAEGER_URL", "jaeger-url", "http://jaeger:16686")
	sentryUrl := getStringValue("SENTRY_URL", "sentry-url", "http://sentry:9000")
	kafkaBrokerUrl := getStringValue("KAFKA_BROKER_URL", "kafka-broker-url", "kafka:9092")
	someAppId := getStringValue("SOME_APP_ID", "some-upp-id", "testId")
	someAppKey := getStringValue("SOME_APP_KEY", "some-app-key", "testKey")

	flag.Parse()

	config := Config{
		Port:           *port,
		DbUrl:          *dbUrl,
		JaegerUrl:      *jaegerUrl,
		SentryUrl:      *sentryUrl,
		KafkaBrokerUrl: *kafkaBrokerUrl,
		SomeAppId:      *someAppId,
		SomeAppKey:     *someAppKey,
	}

	return config, nil
}

func getStringValue(envName string, flagName string, defaultValue string) *string {
	value, exists := os.LookupEnv(envName)

	if exists {
		return &value
	}

	return flag.String(flagName, defaultValue, "dddddddd")
}

func getIntValue(envName string, flagName string, defaultValue int) *int64 {
	value, exists := os.LookupEnv(envName)

	if exists {
		i, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			panic(err)
		}

		return &i
	}

	//check flags here

	return flag.Int64(flagName, int64(defaultValue), "ggggggg")

}
