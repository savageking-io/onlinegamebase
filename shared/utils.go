package shared

import (
	log "github.com/sirupsen/logrus"
)

func SetLogLevel(level string) error {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(log.InfoLevel)
		return err
	}
	log.SetLevel(lvl)
	return nil
}
