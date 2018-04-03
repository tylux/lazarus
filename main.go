package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"

	"github.com/gorilla/mux"

	yaml "gopkg.in/yaml.v1"
)

// Config File Settings
type Config struct {
	Alerts []Alert
}

// Alert Prometheus Alert Definitions
type Alert struct {
	Name      string
	AlertHost string
	Action    string
}

//const lazarusConfig = "./lazarus.yml"

// LoadConfig - Read in Config file.
func LoadConfig() Config {
	files, err := ioutil.ReadDir("/etc/lazarus/conf.d/")
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	for _, file := range files {
		filename := fmt.Sprintf("/etc/lazarus/conf.d/%s", file.Name())
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Print("Unable to Read Config file", err)
		}

		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Print("Unable to Unmarshal Config", err)
		}

	}
	return config

}

//lazarusConfig := "./lazarus.yml"
//

var cfg = LoadConfig()

func main() {
	logwriter, e := syslog.New(syslog.LOG_NOTICE, "lazarus")
	if e == nil {
		log.SetOutput(logwriter)
	}
	log.SetFlags(0)
	// Start Alert Runner
	alertRunner()

	// TODO: Rework for more flexibility
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Print(http.ListenAndServe(":8080", router))
}

// Index Print Status of Service
func Index(w http.ResponseWriter, r *http.Request) {
	alerts := cfg.Alerts
	for _, alert := range alerts {
		fmt.Fprintln(w, "\tName: ", alert.Name)
		fmt.Fprintln(w, "\tAlertHost: ", alert.AlertHost)
		fmt.Fprintln(w, "\tAction: ", alert.Action)

	}
}
