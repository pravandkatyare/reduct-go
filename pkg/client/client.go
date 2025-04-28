package client

import (
	"context"
	"net/url"
	"time"
)

type Client struct {
	// Connection to the Reduct server
	cfg  cfg
	opts []Opt

	ctx       context.Context
	ctxCancel func()

	url     *url.URL
	timeout time.Duration
	headers map[string]string
}

// NewClient also launches a goroutine which periodically updates the cached
// topic metadata.
func New(url, opts []Opt) (*Client, error) {
	cfg, err := validateCfg(opts...)
	if err != nil {
		return nil, err
	}

		ctx := context.Background()

	if cfg.ctx != nil {
		ctx = cfg.ctx
	}

	ctx, cancel := context.WithCancel(ctx)

	cl := &Client{
		cfg:       cfg,
		opts:      opts,
		ctx:       ctx,
		ctxCancel: cancel,

		
	}

	return cl, nil
}

// Opts returns the options that were used to create this client. This can be
// as a base to generate a new client, where you can add override options to
// the end of the original input list. If you want to know a specific option
// value, you can use OptValue or OptValues.
func (cl *Client) Opts() []Opt {
	return cl.opts
}

func (c *Client) Connect() error {
	// Connect to the Reduct server
	return nil
}

func (c *Client) Disconnect() error {
	// Disconnect from the Reduct server
	return nil
}

// ValidateOpts returns an error if the options are invalid.
func ValidateOpts(opts ...Opt) error {
	_, err := validateCfg(opts...)
	return err
}

// This function validates the configuration and returns a few things that we
// initialize while validating. The difference between this and NewClient
// initialization is all NewClient initialization is infallible.
func validateCfg(opts ...Opt) (cfg, error) {
	cfg := defaultCfg()
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}
