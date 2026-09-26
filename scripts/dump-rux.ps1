# Dump ADB RUX (Windows). Pre requis : adb dans le PATH, robot root, USB debug.
#   Set-ExecutionPolicy -Scope Process Bypass
#   .\dump-rux.ps1
#   .\dump-rux.ps1 -MqttHost 192.168.1.10 -MqttSeconds 60

param(
    [string]$OutDir = "",
    [string]$MqttHost = "",
    [int]$MqttSeconds = 45
)

$ErrorActionPreference = "Continue"
if (-not $OutDir) {
    $stamp = Get-Date -Format "yyyy-MM-dd_HHmm"
    $OutDir = Join-Path $PWD "rux-adb-$stamp"
}
New-Item -ItemType Directory -Force -Path $OutDir | Out-Null
function Save($name, $scriptblock) {
    $path = Join-Path $OutDir $name
    Write-Host ">> $name"
    try { & $scriptblock *>&1 | Out-File -FilePath $path -Encoding utf8 }
    catch { $_ | Out-File -FilePath $path -Encoding utf8 -Append }
}

Write-Host "Sortie : $OutDir"
adb start-server | Out-Null
adb devices | Out-File (Join-Path $OutDir "devices.txt") -Encoding utf8
$dev = adb get-serialno 2>$null
if (-not $dev -or $dev -match "unknown") {
    Write-Host "Aucun device ADB."
    exit 1
}

adb root 2>$null | Out-File (Join-Path $OutDir "adb-root.txt") -Encoding utf8
Start-Sleep -Seconds 1
adb wait-for-device

Save "getprop.txt" { adb shell getprop }
Save "getprop-filtre.txt" { adb shell "getprop | grep -iE 'sn|mac|hardcode|region|language|letian|emqx|iot|wifi'" }
Save "settings-global.txt" { adb shell settings list global }
Save "packages.txt" { adb shell "pm list packages | grep -iE 'letian|geeui|ltp|emqx|lex|ota|ident|mijia|renhejia'" }
Save "pm-path.txt" {
    @("com.letianpai.emqxservice","com.renhejia.robot.letianpaiservice","com.letianpai.otaservice") | ForEach-Object { adb shell pm path $_ }
}
Save "services.txt" { adb shell "service list | grep -iE 'letian|sensor|mcu'" }
Save "activity-services.txt" { adb shell "dumpsys activity services | grep -iE 'emqx|letianpai|mcu|task|lex'" }
Save "tty.txt" { adb shell "ls -l /dev/ttyS* /dev/ttyHS* /dev/ttyAML* /dev/ttyUSB* 2>/dev/null" }
Save "tcp.txt" { adb shell "cat /proc/net/tcp" }
Save "connectivity.txt" { adb shell "dumpsys connectivity" }
Save "prefs-grep.txt" {
    adb shell "grep -R -n 'robot-api\|your-server\|letianpai\|1883\|8883\|remote_host' /data/data/com.letianpai.emqxservice/ /data/data/com.letianpai.network/ /data/data/com.rhj.network/ 2>/dev/null | head -n 80"
}
Save "logcat-cloud.txt" { adb logcat -d -t 4000 }
Select-String -Path (Join-Path $OutDir "logcat-cloud.txt") -Pattern "Emqx|Mqtt|IotTriplet|robot_api|404|letianpai" |
    Out-File (Join-Path $OutDir "logcat-filtre.txt") -Encoding utf8

if ($MqttHost) {
    $mqttOut = Join-Path $OutDir "mqtt-dump.txt"
    Write-Host "MQTT # sur $MqttHost ${MqttSeconds}s"
    if (Get-Command mosquitto_sub -ErrorAction SilentlyContinue) {
        $p = Start-Process -FilePath "mosquitto_sub" -ArgumentList @("-h", $MqttHost, "-t", "#", "-v") -RedirectStandardOutput $mqttOut -NoNewWindow -PassThru
        Start-Sleep -Seconds $MqttSeconds
        if (-not $p.HasExited) { Stop-Process -Id $p.Id -Force }
    } else {
        "mosquitto_sub introuvable" | Out-File $mqttOut -Encoding utf8
    }
}

Write-Host "OK. Compress-Archive -Path '$OutDir' -DestinationPath '$OutDir.zip'"
