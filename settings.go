package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// DaemonSettings struct
type DaemonSettings struct {
	Name string `yaml:"name"`
	Web  struct {
		Enabled bool   `yaml:"enabled"`
		Port    int    `yaml:"port"`
		Token   string `yaml:"token"`
	} `yaml:"web"`
	System struct {
		Enabled    bool `yaml:"enabled"`
		SwapReboot int  `yaml:"swapReboot"`
		SwapWarn   int  `yaml:"swapWarn"`
	} `yaml:"system"`
	Ddns struct {
		Enabled bool     `yaml:"enabled"`
		IP      string   `yaml:"ip"`
		Urls    []string `yaml:"urls"`
	} `yaml:"ddns"`
	Storage struct {
		Enabled bool     `yaml:"enabled"`
		Atto    []string `yaml:"atto"`
		Mdadm   []string `yaml:"mdadm"`
		Smart   []string `yaml:"smart"`
	} `yaml:"storage"`
	Rss struct {
		Enabled bool     `yaml:"enabled"`
		Urls    []string `yaml:"urls"`
	} `yaml:"rss"`
	Transmission struct {
		Enabled bool   `yaml:"enabled"`
		Path    string `yaml:"path"`
		Proxy   struct {
			Hostname string `yaml:"hostname"`
			Password string `yaml:"password"`
			Port     int    `yaml:"port"`
			Protocol string `yaml:"protocol"`
			Username string `yaml:"username"`
		} `yaml:"proxy"`
	} `yaml:"transmission"`
	Telegram struct {
		Enabled bool   `yaml:"enabled"`
		Apikey  string `yaml:"apikey"`
		Chatid  int    `yaml:"chatid"`
	} `yaml:"telegram"`
	plex struct {
		Enabled   bool   `yaml:"enabled"`
		PlexToken string `yaml:"plexToken"`
		Release   int    `yaml:"release"`
		Checksum  int    `yaml:"checksum"`
	} `yaml:"plex"`
}

// Settings - Object that holds all default settings from the file
var Settings DaemonSettings
var settingsFile = "./settings.yaml"

// Called by your main package to load the file or die trying
func loadSettings() {

	bytes, ferr := ioutil.ReadFile(settingsFile)
	if ferr != nil {
		log.Panicf("Failed to read settings file %s: %v\n", settingsFile, ferr)
	}

	if err := yaml.Unmarshal([]byte(bytes), &Settings); err != nil {
		log.Panicf("Error reading settings file: %s\n", settingsFile)
	}

	run = true
}
