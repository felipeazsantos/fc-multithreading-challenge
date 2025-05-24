package configs

import (
	"log"
	"os"
)

func InitLogs() IConfig {
	lErr := log.New(os.Stderr, "Error: \t", log.Ldate|log.Ltime|log.Lshortfile)
	lInfo := log.New(os.Stdout, "Info: \t", log.Ldate|log.Ltime)

	return NewConfigs(lErr, lInfo)
}
