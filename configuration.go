package swish

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"time"
)

type Environment struct {
	BaseUrl               string
	CertificationFilePath string
}

type Configuration struct {
	Environment *Environment
	CertFile    string
	KeyFile     string
	Timeout     time.Duration
}

func NewConfiguration(environment *Environment, certFile string, keyFile string, options ...ConfigurationOption) *Configuration {
	instance := &Configuration{Environment: environment, CertFile: certFile, KeyFile: keyFile, Timeout: 60}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance
}

// ConfigurationOption definition
type ConfigurationOption func(*Configuration)

// Function to create ConfigurationOption func to set the timeout limit
func setTimeout(timeout time.Duration) ConfigurationOption {
	return func(subject *Configuration) {
		subject.Timeout = timeout
	}
}

var (
	TestEnvironment       = Environment{BaseUrl: "https://mss.cpc.getswish.net/swish-cpcapi/api/v1", CertificationFilePath: GetResourcePath("certificates/ca.test.crt")}
	ProductionEnvironment = Environment{BaseUrl: "https://mss.cpc.getswish.net/swish-cpcapi/api/v1", CertificationFilePath: GetResourcePath("certificates/ca.prod.crt")}
)

func GetResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "./resource"), nil
}

func GetResourcePath(path string) (directory string) {
	dir, err := GetResourceDirectoryPath()

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
