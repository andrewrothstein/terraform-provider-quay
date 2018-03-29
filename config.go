package main

import (
	"github.com/coreos/go-quay/quay"
)

// Config :
type Config struct {
	Token        string
	Organization string
}

// Organization :
type Organization struct {
	name   string
	client *quay.Client
}

// Client :
func (c *Config) Client() (interface{}, error) {
	var org Organization
	org.name = c.Organization

	org.client = quay.NewHTTPClient(nil)
	return &org, nil
}
