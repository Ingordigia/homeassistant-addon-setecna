package homeassistant

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Ingordigia/homeassistant-addon-setecna/models"
	"github.com/Ingordigia/homeassistant-addon-setecna/pkg/helpers"
	"github.com/Ingordigia/homeassistant-addon-setecna/pkg/mqtt"
)

func CreateClimates(responseMap map[string]string, systemID string) (msgs mqtt.Messages) {
	var systemSeason helpers.Season
	if responseMap["GLOBAL_SEASON"] == strconv.Itoa(0) {
		systemSeason = helpers.Winter
	} else {
		systemSeason = helpers.Summer
	}
	for i := 1; i <= 32; i++ {
		if responseMap["Z"+fmt.Sprint(i)+"_SENSOR_CHN"] != "0" {
			var j []byte
			var err error = nil
			if responseMap["Z"+fmt.Sprint(i)+"_RH"] != "32769" {
				var climate = new(models.ClimateWithHumidity)
				climate.Init(i, systemID, systemSeason)

				j, err = json.Marshal(climate)
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				var climate = new(models.ClimateWithoutHumidity)
				climate.Init(i, systemID, systemSeason)

				j, err = json.Marshal(climate)
				if err != nil {
					log.Println(err)
					return
				}
			}
			message := mqtt.Message{
				Topic:   "homeassistant/climate/" + systemID + "_zone_" + fmt.Sprint(i) + "/config",
				Message: string(j),
				Qos:     0,
			}
			msgs = append(msgs, message)
		}
	}
	return msgs
}

func RemoveClimates(responseMap map[string]string, systemID string) (msgs mqtt.Messages) {
	for i := 1; i <= 32; i++ {
		if responseMap["Z"+fmt.Sprint(i)+"_TEMP"] != "32769" {
			message := mqtt.Message{
				Topic:   "homeassistant/climate/" + systemID + "_zone_" + fmt.Sprint(i) + "/config",
				Message: "",
				Qos:     0,
			}
			msgs = append(msgs, message)
		}
	}
	return msgs
}
