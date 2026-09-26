#!/bin/sh
# Usage: ./pub.sh [clientId] [cmd] [json-d]
# clientId = l81_... (PAS le SN)
CID=${1:-l81_9aa8495f6c9ddf95920428e3c2352d4c}
CMD=${2:-controlMotion}
DATA=${3:-'{"motion":"null","motion_name":"交替向前走","number":98,"step":1,"speed":2}'}
ET=$(($(date +%s) * 1000 + 300000))
TS=$(date +%s)
TOPIC="cmd/L81/${CID}/${CMD}/${TS}"
PAYLOAD="{\"cmd\":\"${CMD}\",\"d\":${DATA},\"et\":${ET}}"
echo "$TOPIC"
echo "$PAYLOAD"
if command -v mosquitto_pub >/dev/null 2>&1; then
  mosquitto_pub -h 127.0.0.1 -p 1883 -q 1 -t "$TOPIC" -m "$PAYLOAD"
elif command -v docker >/dev/null 2>&1; then
  docker exec rux-mqtt mosquitto_pub -h 127.0.0.1 -q 1 -t "$TOPIC" -m "$PAYLOAD"
else
  echo "mosquitto_pub manquant" >&2; exit 1
fi
