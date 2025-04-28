package client

import (
	"context"
	"time"
)

// Opt is an option to configure a client.
type Opt interface {
	apply(*cfg)
}

type cfg struct {
	ctx        context.Context
	timeout    time.Duration
	apiVersion string

	verifySSL bool

	//max block size in bytes
	maxBlockSize int64

	//max number of records in a block
	maxRecords int64
}

type clientOpt struct{ fn func(*cfg) }

func (opt clientOpt) apply(cfg *cfg) { opt.fn(cfg) }

func (cfg *cfg) validate() error {
	// Validate the configuration

	return nil
}

func defaultCfg() cfg {
	return cfg{
		// Set default values for the configuration
	}
}

func SetTimeOut(timeout time.Duration) Opt {
	return clientOpt{func(cfg *cfg) { cfg.timeout = timeout }}
}
