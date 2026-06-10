<!-- https://developers.home-assistant.io/docs/add-ons/presentation#keeping-a-changelog -->

## 1.1.2

- Bump home-assistant/builder from 2024.08.2 to 2025.03.0
- Bump docker/login-action from 3.3.0 to 3.4.0
- Bump docker/login-action from 3.4.0 to 3.7.0
- Bump actions/checkout from 4.2.2 to 6.0.2
- Bump home-assistant/builder from 2025.03.0 to 2025.11.0
- Bump frenck/action-addon-linter from 2.18 to 2.21

## 1.1.1

- Bump home-assistant/builder from 2024.08.1 to 2024.08.2
- Bump actions/checkout from 4.1.7 to 4.2.1
- Bump frenck/action-addon-linter from 2.15 to 2.17
- Bump actions/checkout from 4.2.1 to 4.2.2
- Bump frenck/action-addon-linter from 2.17 to 2.18
- Add network capability to make this run on HA Supervised on Debian 12

## 1.1.0

- Add MTx_MODE as a sensor and MTx_FORCING as a selector to HomeAssistant
- **BREAKING CHANGE**: Change how a zone is considered active, now the plugin check if Zx_SENSOR_CHN != 0 instead of Zx_TEMP != 32769 (aligned with the web interface logic)

## 1.0.1

- Fixes decimals in command template for climate entities

## 1.0.0

- Initial release
