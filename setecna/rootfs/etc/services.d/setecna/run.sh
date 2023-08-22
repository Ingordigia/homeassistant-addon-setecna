#!/usr/bin/with-contenv bashio
# ==============================================================================
# Start the service
# s6-overlay docs: https://github.com/just-containers/s6-overlay
# ==============================================================================

# Add your code here

# Declare variables
declare REG_SYSTEM_ID
declare REG_USER
declare REG_PASSWORD
declare MQTT_HOST
declare MQTT_USER
declare MQTT_PASSWORD

## Get the 'systemID' key from the user config options.
REG_SYSTEM_ID=$(bashio::config "systemID")
REG_USER=$(bashio::config "username")
REG_PASSWORD=$(bashio::config "password")
MQTT_HOST=$(bashio::services mqtt "host")
MQTT_USER=$(bashio::services mqtt "username")
MQTT_PASSWORD=$(bashio::services mqtt "password")

## Print the systemID the user supplied, defaults to "Undefined"
bashio::log.info "${systemID:=""}"

## Run your program
exec /usr/bin/app
