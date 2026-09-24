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
