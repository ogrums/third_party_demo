# third_party_demo

Local stand-in for a third-party large language model. It is an HTTP server, not the Letianpai robot cloud. It does not implement `/robot_api/...` (bind, OTA, logo, weather, calendar).

`POST /sse` on `127.0.0.1:8012` reads a JSON body:

```json
{"app_id":"...","api_key":"...","api_secret":"...","content":"who are you","sn":"..."}
```

The response is a Server-Sent Events stream. Each `data:` line is JSON `{code, msg, data:{content, is_end}}`. The text is canned. Matching is a substring of `content`:

| Content contains | Reply |
|---|---|
| `who` or `model` | short introduction |
| `snow` | a poem, sent line by line |
| `product manager` | a short essay |
| anything else | "I do not know" |

Chinese keywords (`是谁`, `大模型`, `园春`, `产品经理`) still match.

## Build

```
# linux
./build.sh linux

# mac
./build.sh
```

## Mock robot cloud

`mock/main.go` is a separate process from the language-model demo. It answers the JSON envelope `{code, msg, data}` used by `LtpNetWork` and by the launcher parsers.

```
go run ./mock
curl -s http://127.0.0.1:8080/robot_api/v1/bind/getSnByMac
curl -s http://127.0.0.1:8080/robot_api/v1/device/weather
```

`getSnByMac` returns:

```json
{"code":0,"msg":"success","data":{"client_id":"mock-client","hard_code":"mock-hardcode","sn":"EMULATOR00000000"}}
```

Bind, OTA, upload token, logo, weather, calendar, countdown, fans, clock, bind status, and server time are the same shape. An unknown path returns `code: 404`. The launcher still calls `https://yourservice.com` plus the placeholder path `your interface url`, so point `GeeUINetworkConsts` and `Constants.kt` at `http://127.0.0.1:8080` and the route you want before this mock is used.
