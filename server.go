package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"

	"baazaru.com/components/database"
	"baazaru.com/components/email"
	"baazaru.com/components/jsonconfig"
	"baazaru.com/components/pushover"
	"baazaru.com/components/recaptcha"
	"baazaru.com/components/server"
	"baazaru.com/components/session"
	"baazaru.com/components/view"
	"baazaru.com/components/view/plugin"
	"baazaru.com/route"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// Connect to database
	database.Connect(config.Database)

	// Configure the SMTP server
	email.Configure(config.Email)

	// Configure the Google reCAPTCHA prior to loading view plugins
	recaptcha.Configure(config.Recaptcha)

	// Configure Pushover
	pushover.Configure(config.Pushover)

	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	view.LoadPlugins(plugin.TemplateFuncMap(config.View))

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Databases      `json:"Database"`
	Email     email.SMTPInfo          `json:"Email"`
	Recaptcha recaptcha.RecaptchaInfo `json:"Recaptcha"`
	Pushover  pushover.PushoverInfo   `json:"Pushover"`
	Server    server.Server           `json:"Server"`
	Session   session.Session         `json:"Session"`
	Template  view.Template           `json:"Template"`
	View      view.View               `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
