package testing

import (
	"strings"
	"testing"

	"github.com/gabrielerzinger/goCoupon/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetLogger gets a new logger instance
func GetLogger(t *testing.T) logrus.FieldLogger {
	return logrus.New()
}

// GetApp creates a new app for test
func GetApp(t *testing.T) *app.App {
	config := viper.New()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.SetEnvPrefix("GOCOUPON")
	config.AutomaticEnv()
	app, err := app.NewApp("0.0.0.0", 8000, config, GetLogger(t))
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return app
}
