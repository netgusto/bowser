package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"

	gosxnotifier "github.com/deckarep/gosx-notifier"
	"github.com/skratchdot/open-golang/open"
)

var devBrowser = "Google Chrome"
var surfBrowser = "Safari"

func handleURL(config Config, url string) chan interface{} {
	wait := make(chan interface{})

	go func() {
		log.Println("Received URL")

		var matchedBrowser *Browser

		for _, browser := range config.Browsers {
			for _, rgx := range browser.Match {
				success, _ := regexp.MatchString(rgx, url)
				if success {
					matchedBrowser = &browser
					break
				}
			}

			if matchedBrowser != nil {
				break
			}
		}

		if matchedBrowser == nil {
			// Looking for defautl browser
			for _, browser := range config.Browsers {
				if browser.Alias == "Default" {
					matchedBrowser = &browser
					break
				}
			}
		}

		if matchedBrowser == nil {
			// Cannot find default browser; fall back to Safari
			open.RunWith(url, "Safari")
			// open.RunWith("https://github.com/netgusto/bowser", "Safari")
			// open.RunWith("/tmp/hello.txt", "TextEdit")

			notifyProblem(
				"Cannot find a default browser - falling back to Safari\nSee https://github.com/netgusto/bowser",
				"Bowser - No default browser",
			)

		} else {
			err := open.RunWith(url, matchedBrowser.App)
			if err != nil {
				log.Println(err)

				if err2 := open.RunWith(url, "Safari"); err2 == nil {
					notifyProblem(
						"URL Matched \""+matchedBrowser.Alias+"\" but failed to open \""+matchedBrowser.App+"\" - Falling back to Safari",
						"Bowser - Problem with browser "+matchedBrowser.Alias,
					)
				} else {
					notifyProblem(
						"URL Matched \""+matchedBrowser.Alias+"\" but failed to open \""+matchedBrowser.App+"\" - Could not fall back to Safari",
						"Bowser - Problem with browser "+matchedBrowser.Alias,
					)
				}
			}
		}

		wait <- nil
	}()

	return wait
}

func notifyProblem(msg string, title string) {
	note := gosxnotifier.NewNotification(msg)
	note.Title = title
	note.AppIcon = getNotificationIconPath()
	note.Link = "https://github.com/netgusto/bowser#setup-default-browser"
	note.Push()
}

func getNotificationIconPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir + "/" + "bowser.png"
}
