#!/bin/sh
# Publish one EmqxService-shaped command.
# Usage: ./pub.sh [sn] [cmd] [json-d]
# Example: ./pub.sh EMULATOR00000000 controlMotion '{"motion":"forward","number":1}'

set -e
SN=${1:-EMULATOR00000000}
CMD=${2:-controlMotion}
DATA=${3:-'{"motion":"forward","number":1}'}
ET=$(($(date +%s) * 1000 + 120000))
TOPIC="cmd/L81/${SN}/cmd/mock"
PAYLOAD="{\"cmd\":\"${CMD}\",\"d\":${DATA},\"et\":${ET}}"

echo "publish $TOPIC"
echo "$PAYLOAD"

if command -v mosquitto_pub >/dev/null 2>&1; then
  mosquitto_pub -h 127.0.0.1 -p 1883 -t "$TOPIC" -m "$PAYLOAD"
elif command -v docker >/dev/null 2>&1; then
  docker exec rux-mqtt mosquitto_pub -h 127.0.0.1 -t "$TOPIC" -m "$PAYLOAD"
else
  echo "install mosquitto-clients or start docker compose first" >&2
  exit 1
fi
