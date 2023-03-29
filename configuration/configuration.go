package configuration

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var (
	defaultParallel     = 10
	errNoUrls           = fmt.Errorf("no urls provided")
	errParallelZero     = fmt.Errorf("parallel flag can not be zero")
	errParallelNegative = fmt.Errorf("parallel flag can not be negative")
)

// Configuration - configuration for the application
type Configuration struct {
	Parallel int
	Urls     []string
}

// Validate - validates the configuration
func (c Configuration) Validate() error {
	if c.Parallel == 0 {
		return errParallelZero
	}
	if c.Parallel < 0 {
		return errParallelNegative
	}
	if len(c.Urls) == 0 {
		return errNoUrls
	}
	for i, urlString := range c.Urls {
		_, err := url.ParseRequestURI(urlString)
		if err != nil {
			newUrlString := fmt.Sprintf("http://%s", urlString)
			_, err := url.ParseRequestURI(newUrlString)
			if err != nil {
				return err
			}
			c.Urls[i] = newUrlString
		}
	}
	return nil
}

// ParseCmdLine - parses the configuration based on command line flags
func ParseCmdLine() (Configuration, error) {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	parallel := flagSet.Int("parallel", defaultParallel, "Limit the number of parallel requests")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		return Configuration{}, err
	}
	config := Configuration{
		Parallel: *parallel,
		Urls:     flagSet.Args(),
	}
	return config, config.Validate()
}
