# Robot cloud APIs

The production paths were replaced with `your interface url` or `/your_api` before open source. Only the rows marked **known** appear as a real path in source. The other rows are the names the apps still call. Their mock route returns `{code:0, msg:"mock stub, not implemented", data:{}}`. Sample bodies already exist for the LetianpaiOS home-screen routes. Real payloads are a later step.

The mock is `go run ./mock` on `127.0.0.1:8080`. LetianpaiOS debug builds call it at `http://10.0.2.2:8080`. LtpNetWork `Constants.kt` uses the same host.

## Known paths

| Method | Path | Called by |
|---|---|---|
| GET | `/robot_api/v1/bind/getSnByMac` | LtpNetWork, LetianpaiOS |
| GET | `/robot_api/v1/bind/getIotTriplet` | LtpNetWork |
| POST | `/robot_api/v1/device/upgrade/status` | LtpNetWork |
| GET | `/robot_api/v1/ota/getLatestPackage` | LtpNetWork. GeeUIComponents `GET_LATEST_PACKAGE` and `GET_LAST_PACKAGE` are the same idea, path redacted |
| POST | `/robot_api/v1/device/addRecord` | LtpNetWork upload |
| GET | `/robot_api/v1/cloudFile/getToken` | LtpNetWork. GeeUIComponents `CLOUD_FILE_TOKEN` comment `get_robot_api_v1_cloudFile_getToken` |
| GET | `/robot_api/v1/cloudFile/getSessionToken` | LtpNetWork. GeeUIComponents `GET_SESSION_TOKEN` |
| GET | `/robot_api/v1/device/code/getInfo` | GeeUIInstaller swagger comment, GeeUIComponents `GET_BIND_CODE` |
| GET | `/robot_api/v1/device/ip/getRegion` | GeeUIComponents comment `get_robot_api_v1_device_ip_getRegion` |
| GET | `/robot_api/v1/common/getConfig` | GeeUIFace `AutoService` (`config_key=remote_stroll`) |
| GET | `/index/hello` | LtpNetWork test |
| GET | `/index/getorder` | LtpNetWork test |
| POST | `/index/addorder` | LtpNetWork test |
| POST | `/addons/shop/checkout/submit` | LtpNetWork test |
| POST | `/sse` on port 8012 | third_party_demo language model, GeeUIAIAudioService when `baseUrl` points there |

## LetianpaiOS

Wired to the mock. Sample JSON except the three stubs.

| Constant | Mock path | Sample |
|---|---|---|
| `BIND_INFO` | `/robot_api/v1/device/bindInfo` | yes |
| `CALENDAR_LIST` | `/robot_api/v1/device/calendar` | yes |
| `COUNTDOWN_LIST` | `/robot_api/v1/device/countdown` | yes |
| `FANS_INFO_LIST` | `/robot_api/v1/device/fans` | yes |
| `GENERAL_INFO` | `/robot_api/v1/device/general` | yes |
| `WEATHER_INFO` | `/robot_api/v1/device/weather` | yes |
| `CLOCK_LIST` | `/robot_api/v1/device/clock` | yes |
| `GET_SN_BY_MAC` | `/robot_api/v1/bind/getSnByMac` | yes |
| `GET_SERVER_TIME_STAMP` | `/robot_api/v1/device/serverTime` | yes |
| `GET_DEVICE_CHANNELLOGO` | `/robot_api/v1/device/logo` | yes |
| `GET_ALL_CONFIG` | `/robot_api/v1/device/allConfig` | stub |
| `GET_COMMON_CONFIG` | `/robot_api/v1/device/commonConfig` | stub |
| `UPLOAD_STATUS` | `/robot_api/v1/device/uploadStatus` | stub |

`GeeUINetConsts.CALENDAR_LIST` is a second unused placeholder.

## GeeUIComponents and GeeUIFace

Same constant file. Path still `your interface url` in those repos, so they do not call the mock until that string changes. Stub route is `/robot_api/v1/todo/<NAME>`.

`CALENDAR_LIST`, `COUNTDOWN_LIST`, `FANS_INFO_LIST`, `GENERAL_INFO`, `CUSTOM_WATCH_CONFIG`, `CLOUD_FILE_TOKEN`, `WEATHER_INFO`, `STOCK_INFO`, `IS_DEVICE_BIND`, `GET_REGION_BY_DEVICE_IP`, `CUSTOM_LIST`, `CUSTOM_PHOTO_LIST`, `COMMEMORATION_LIST`, `LAMP_CUSTOM_INFO`, `NEWS_LIST`, `GET_SN_BY_MAC`, `GET_SESSION_TOKEN`, `GET_MEDITATION_CONFIG`, `CLOCK_LIST`, `GET_ALL_CONFIG`, `GET_USER_APPS_CONFIG`, `GET_APPS_SHOW_CONFIG`, `UPLOAD_STATUS`, `UPLOAD_BATTERY_STATUS`, `UPLOAD_LEX_LOG`, `GET_COMMON_CONFIG`, `POST_MODULE_CHANGE`, `POST_RESET_STATUS`, `GET_ALL_APP_LIST`, `GET_APP_LIST`, `GET_RECHARGE_CONFIG`, `GET_APP_BG_INFO`, `GET_LATEST_PACKAGE`, `GET_BIND_CODE`, `POST_UPLOAD_APP_STATUS`, `POST_UPLOAD_USER_APP_STATUS`, `GET_USER_REMIND_LIST`, `GET_SERVER_TIME_STAMP`, `GET_TOMATO_LIST`, `GET_LAST_PACKAGE`, `POST_MANAGE_ADD`, `GET_DEVICE_CHANNELLOGO`.

GeeUITaskService has the same list commented out. GeeUITime only has `CALENDAR_LIST`.

## GeeUIAIAudioService

`com.rhj.speech.http.Api` extends LtpNetWork `NewApi`, so it also calls the known bind and OTA paths. These five methods are still `/your_api`:

| Method | Retrofit | Mock stub |
|---|---|---|
| `uploadLog` | POST | `/robot_api/v1/todo/uploadLog` |
| `uploadCall` | POST | `/robot_api/v1/todo/uploadCall` |
| `getConfigData` | GET | `/robot_api/v1/todo/getConfigData` |
| `getWakeConfig` | GET | `/robot_api/v1/todo/getWakeConfig` |
| `getAiConfig` | GET | `/robot_api/v1/todo/getAiConfig` |

The LLM stream is not one of these. It uses `AiRuntimeConfig.baseUrl` plus `POST /sse`.

## GeeUISetting

Every screen passes the literal `your interface url`. No path can be recovered. Screens: time zone, 24-hour clock, mode switch, wake language, multimodal wake, wake voice, wake word, brightness, restore, unbind, language, sleep snoring, sleep schedule, reset, charging, update, Mijia, device info.

## Apps with no cloud call

GeeUIDesktop, GeeUIGuide, GeeUIMcuService, GeeUIWiFiConnector, GeeUIBase, LetianpaiService, DemoForRobotSDK, GeeUIInstaller (only the swagger comment above).
