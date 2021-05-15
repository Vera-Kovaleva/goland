package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

func checkVer() {
	//...
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "verbosity output")
	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

}
