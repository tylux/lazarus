package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/buger/jsonparser"
)

func alertRunner() {
	alerts := cfg.Alerts
	for _, alert := range alerts {
		go alertLookup(alert)
	}
}

func alertLookup(a Alert) {
	for {
		log.Print("Querying Prometheus Alert Manager for active alerts")
		resp, err := http.Get(a.AlertHost + "/api/v1/alerts/")
		if err != nil {
			log.Print(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)

		//loop over json returned from alertmanager API, drill down into data, labels, alertname
		jsonparser.ArrayEach(body, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			active, err := jsonparser.GetString(value, "labels", "alertname")
			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
			if active == a.Name {
				log.Printf("Running action %s", a.Action)
				action := a.Action
				runAction(action)
				time.Sleep(120 * time.Second)
			}
		}, "data") //top level json that contains the list of alerts
		time.Sleep(60 * time.Second)
	}
}

func runAction(action string) {
	cmd := exec.Command(action)

	stdout, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
	log.Print(string(stdout))
}
