package configs

import "log"

type IConfig interface {
	LErr() *log.Logger
	LInfo() *log.Logger
}

type config struct {
	lErr  *log.Logger
	lInfo *log.Logger
}

func NewConfigs(lErr *log.Logger, lInfo *log.Logger) IConfig {
	return &config{
		lErr:  lErr,
		lInfo: lInfo,
	}
}

func (c *config) LErr() *log.Logger {
	return c.lErr
}

func (c *config) LInfo() *log.Logger {
	return c.lInfo
}
