package main

import (
	"Vera-2020-02-21_01/lesson_08/config"
	"fmt"
	"log"
)

func main() {

	getConfig, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		"Port:", getConfig.Port,
		"\nDbUrl:", getConfig.DbUrl,
		"\nJaegerUrl:", getConfig.JaegerUrl,
		"\nSentryUrl:", getConfig.SentryUrl,
		"\nKafkaBrokerUrl:", getConfig.KafkaBrokerUrl,
		"\nSomeAppId:", getConfig.SomeAppId,
		"\nSomeAppKey:", getConfig.SomeAppKey)

}
