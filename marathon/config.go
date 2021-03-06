package marathon

import (
	"github.com/gambol99/go-marathon"
	"time"
)

type Config struct {
	Url                      string
	RequestTimeout           int
	DefaultDeploymentTimeout time.Duration
	BasicAuthUser            string
	BasicAuthPassword        string

	Client marathon.Marathon
}

func (c *Config) loadAndValidate() error {

	// this needs to return an err as well.
	marathonConfig := marathon.NewDefaultConfig()
	marathonConfig.URL = c.Url
	marathonConfig.RequestTimeout = c.RequestTimeout
	marathonConfig.HTTPBasicAuthUser = c.BasicAuthUser
	marathonConfig.HTTPBasicPassword = c.BasicAuthPassword

	client, err := marathon.NewClient(marathonConfig)
	c.Client = client
	return err
}
