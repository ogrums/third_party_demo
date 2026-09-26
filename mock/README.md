# Mock cloud RUX

Two processes. Not the LLM demo (`POST /sse` :8012).

## HTTP — `go run .` in this folder

`http://127.0.0.1:8080` — envelope `{code,msg,data}`.
`getIotTriplet` already returns `remote_host=127.0.0.1` `remote_port=1883`
`user_name=mock-device` `password_hash=mock-password` `client_id=mock-client`
`getSnByMac` returns `sn=EMULATOR00000000`.

Point `LtpNetWork` `Constants.kt` at this host before a robot boot.

## MQTT — Mosquitto

```
docker compose -f mock/docker-compose.yml up
# from repo root; or from this directory: docker compose up
```

Anonymous :1883 (and websocket :9001). EmqxService also sends user/password from the triplet; Mosquitto 2 accepts them when `allow_anonymous true`.

Dump everything:

```
mosquitto_sub -h 127.0.0.1 -t '#' -v
```

After EmqxService connects it SUB `cmd/L81/<sn>/+/+` and PUB ack on `cmd_resp/L81/<sn>/…`.

Send a motion command:

```
chmod +x mock/pub.sh
./mock/pub.sh EMULATOR00000000 controlMotion '{"motion":"forward","number":1}'
```

Payload shape (from EmqxService.apk):

```json
{"cmd":"controlMotion","d":{"motion":"forward","number":1},"et":1730000000000}
```

`et` must be in the future (ms). Expired messages are dropped.
