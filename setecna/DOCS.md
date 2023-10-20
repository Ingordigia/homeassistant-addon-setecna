# DISCLAIMER

This add-on is developed by reverse engineering the Setecna web-interface and is not officially supported by the Setecna team, use it with caution.

# Home Assistant Add-on: Setecna add-on

## Why this addon

This add-on is meant to integrate your REG based thermal power plant in Home Assistant.

It's a web based integration, so your system needs to have access to the internet in order to comunicate with Setecna servers first.

## Prerequisites

Before you can use this add-on you need to:
1. Install an MQTT add-on
1. Enable and configure the MQTT integration

*The Setecna add-on will automatically connect to the MQTT broker installed in your Home Assistance instance.*

## How to use


Once installed use the "configuration" tab to insert the following informations:

### Required parameters
- SystemID ( find it once logged to the Setecna web-interface )
- Username
- Password

### Optional parameters
- Readonly:
    - `OFF` = All parameters will be created in HA as `sensors` or `binary_sensors`, disabling any possible interactions beetween Home Assistant and your REG system.
    - `ON` = Configuration parameters that you can change from Setecna web-interface will be created as `numbers` inside Home Assistant, enabling you to control your system inside Home Assistant
- Advanced integration:
    - `ON` = The addon will match Home Assistant's built-in entities such as `climate` or `water_heater` with REG systems's zones and DWH parameters.
    - `OFF` = Home Assistant's built-in entities will not be created by the add-on (advanced users can still create them in their homeassistant's configuration.yaml)