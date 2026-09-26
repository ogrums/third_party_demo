# Mock MQTT RUX

Broker officiel vu : `tcp://43.153.69.45:1883`.
Robot SUB `cmd/L81/<clientId>/+/+` qos1. clientId ≠ SN.

## 1. Broker local

```
cd third_party_demo/mock
docker compose up -d
mosquitto_sub -h 127.0.0.1 -t '#' -v
```

Anonyme :1883. Emqx envoie user/pass du triplet ; Mosquitto 2 les ignore si `allow_anonymous true`.

## 2. Pointer le robot (MQTT ignore le DNS)

`su` + DNAT (IP_PC = IPv4 du PC) :

```
adb shell su -c "iptables -t nat -A OUTPUT -p tcp -d 43.153.69.45 --dport 1883 -j DNAT --to-destination IP_PC:1883"
```

Ou firewall routeur : DNAT 43.153.69.45:1883 → PC.
Log attendu : `Connected to: tcp://IP_PC:1883`.

HTTP (getIotTriplet) reste optionnel si DNAT MQTT seul : le triplet officiel continue de donner 43.153.69.45, d’où le DNAT.

## 3. Publier une marche

```
./pub.sh l81_9aa8495f6c9ddf95920428e3c2352d4c
# équivaut à AT+MOVEW,98,1,2
./pub.sh l81_9aa8495f6c9ddf95920428e3c2352d4c controlMotion '{"motion":"null","motion_name":"立正","number":0,"step":1,"speed":3}'
```

Windows : `pub.ps1` (mêmes args). Topic `cmd/L81/<cid>/controlMotion/<unix>` qos1, `et` dans le futur.
