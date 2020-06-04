package pingdom

import (
	"log"
	"net/http"
	"os"

	"github.com/AN0DA/go-pingdom/pingdom"
)

type Config struct {
	APIToken   string `mapstructure:"api_token"`
	BaseURL    string
	HTTPClient *http.Client
}

// Client() returns a new client for accessing pingdom.
//
func (c *Config) Client() (*pingdom.Client, error) {

	if v := os.Getenv("PINGDOM_TOKEN"); v != "" {
		c.APIToken = v
	}
	c.BaseURL = "https://api.pingdom.com/api/3.1"

	client, _ := pingdom.NewClientWithConfig(pingdom.ClientConfig{APIToken: c.APIToken})

	log.Printf("[INFO] Pingdom Client configured for Token: %s", c.APIToken)

	return client, nil
}
