# third_party_demo

Local stand-in for a third-party large language model. It is an HTTP server, not the Letianpai robot cloud. `POST /sse` on `127.0.0.1:8012`.

## Mock robot cloud (separate)

See [mock/README.md](mock/README.md).

```
# HTTP API
cd mock && go run .
curl -s http://127.0.0.1:8080/robot_api/v1/bind/getSnByMac

# MQTT broker (EmqxService)
docker compose -f mock/docker-compose.yml up
./mock/pub.sh EMULATOR00000000 controlMotion '{"motion":"forward","number":1}'
```

Point `Constants.kt` at `http://<pc>:8080`. `getIotTriplet` already advertises MQTT `127.0.0.1:1883`.

## LLM demo

`POST /sse` JSON `{app_id, api_key, api_secret, content, sn}`. SSE lines `{code, msg, data:{content, is_end}}`.

Build: `./build.sh linux` or `./build.sh`.
