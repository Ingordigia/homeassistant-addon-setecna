package main

import (
	"fmt"
	"os"
)

var user string = os.Getenv("REG_USER")
var password string = os.Getenv("REG_PASSWORD")
var systemID string = os.Getenv("REG_SYSTEM_ID")

var refreshTopic string = "HA/RESTARTED"
var changeTopic string = "IDROSISTEMI/CHANGE/#"
var loginURL string = "https://s5a.eu/login"
var fetchUpdatesURL string = "https://s5a.eu/station/" + systemID + "/getres?timestamp="
var askRefreshURL string = "https://s5a.eu/station/" + systemID + "/askrefresh?connrq=1"
var pushUpdatesURL string = "https://s5a.eu/station/" + systemID + "/putmprop?statid=" + systemID + "&userid=guest&pcount=1" //&p0=ACS_SET_COMFORT&nb0=400

func main() {

	fmt.Println("SystemID:", systemID)
	fmt.Println("Username:", user)

}
