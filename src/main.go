package main

import (
	"log"
	"log/syslog"
	"time"

	"github.com/arkan/dotconfig"
	gosxnotifier "github.com/deckarep/gosx-notifier"
	"github.com/skratchdot/open-golang/open"
)

var urlStream chan string
var appname = "bowser"

func main() {

	config, err := getConfig()
	if err != nil {
		log.Panicln(err)
	}

	if config.Debug {
		// Configure logger to write to the syslog
		logwriter, err := syslog.New(syslog.LOG_NOTICE, appname)
		if err == nil {
			log.SetOutput(logwriter)
		}
	}

	select {
	case <-listen(config):
		if config.IsDefault {
			note := gosxnotifier.NewNotification("A default configuration has been set for Bowser in ~/.config/bowser/config.yml")
			note.Title = "Bowser - Please configure"
			note.AppIcon = getNotificationIconPath()
			note.Link = "https://github.com/netgusto/bowser#setup"
			note.Push()
		}
		log.Println("Done")
	case <-time.After(1000 * time.Millisecond):
		open.RunWith("https://github.com/netgusto/bowser", "Safari")
		log.Println("Not received any URL after 1s; exiting")
	}
}

// Config ...
type Config struct {
	Debug     bool      `yaml:"debug"`
	Browsers  []Browser `yaml:"browsers,omitempty"`
	IsDefault bool      `yaml:"-"`
}

// Browser ...
type Browser struct {
	Alias string   `yaml:"alias"`
	App   string   `yaml:"app,omitempty"`
	Match []string `yaml:"match,omitempty"`
}

func getConfig() (Config, error) {
	config := Config{}

	if err := dotconfig.Load(appname, &config); err != nil {
		if err == dotconfig.ErrConfigNotFound {
			initDefaultConfig(&config)
			if err := dotconfig.Save(appname, config); err != nil {
				return config, err
			}
		}
	} else if err != nil {
		return config, err
	}

	return config, nil
}

func initDefaultConfig(config *Config) {
	config.IsDefault = true
	config.Debug = false
	config.Browsers = []Browser{
		Browser{
			Alias: "Default",
			App:   "Safari",
		},
		// Browser{
		// 	Alias: "Dev",
		// 	App:   "Google Chrome",
		// 	Match: []string{
		// 		"^https?://127.0.0.1",
		// 		"^https?://localhost",
		// 	},
		// },
	}
}
