# third_party_demo

A local HTTP server that pretends to be a third-party large-model API. It does not call a model, and it does not talk to the robot, `RobotSdk`, AIDL, or the MCU.

The robot voice service can point its LLM URL at this process while a real provider is not available.

## Run

```bash
# Linux amd64 binary, then run it
./build.sh linux

# Host platform
./build.sh
```

`build.sh` writes `bin/main_third_party_server` and starts it. A prebuilt binary is already in `bin/`. `main()` listens on `:8012`.

## Protocol

`POST /sse` with a JSON body:

| Field | Meaning |
|---|---|
| `app_id` | ignored by the canned replies |
| `api_key` | ignored |
| `api_secret` | ignored |
| `content` | user text; a substring picks the script |
| `sn` | device serial, ignored |

The response is server-sent events. Each event is `event: message` and a JSON object `code`, `msg`, `data.content`, `data.is_end`. Chunks are flushed every 500 ms.

The script is chosen by a substring of `content`:

| Substring | Script |
|---|---|
| `是谁` or `大模型` | who-is / large-model answer |
| `园春` | Yuanchun answer |
| `产品经理` | product-manager answer |
| anything else | default answer |

The log line "收到请求的body数据" means "received request body". There are no other Chinese comments.

## What this is not

Do not ship this as the production LLM. `GeeUIAIAudioService` expects the same SSE shape (`event: message`, token text, an end flag) from the real Spark or OpenAI proxy.
