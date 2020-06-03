package pingdom

import (
	"log"
	"os"

	"github.com/grnhse/go-pingdom/pingdom"
)

type Config struct {
	APIToken string `mapstructure:"api_token"`
}

// Client() returns a new client for accessing pingdom.
//
func (c *Config) Client() (*pingdom.Client, error) {

	if v := os.Getenv("PINGDOM_TOKEN"); v != "" {
		c.APIToken = v
	}

	client := pingdom.NewClient(c.APIToken)

	log.Printf("[INFO] Pingdom Client configured for user: %s", c.User)

	return client, nil
}
