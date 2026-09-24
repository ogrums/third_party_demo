// Local stand-in for the Letianpai cloud.
// Run: go run ./mock
// Listens on http://127.0.0.1:8080
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func envelope(data any) map[string]any {
	return map[string]any{"code": 0, "msg": "success", "data": data}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func main() {
	now := time.Now().Unix()
	packageItem := map[string]any{
		"is_update":      0,
		"version":        "1.0.0-mock",
		"rom_package_url": "http://127.0.0.1:8080/files/mock.zip",
		"md5":            "d41d8cd98f00b204e9800998ecf8427e",
		"upgrade_desc":   "mock package, nothing to install",
		"update_time":    now,
	}

	routes := map[string]any{
		// LtpNetWork client (NewApi / Api).
		"/robot_api/v1/bind/getSnByMac": envelope(map[string]any{
			"client_id": "mock-client",
			"hard_code": "mock-hardcode",
			"sn":        "EMULATOR00000000",
		}),
		"/robot_api/v1/bind/getIotTriplet": envelope(map[string]any{
			"client_id":     "mock-client",
			"password_hash": "mock-password",
			"user_name":     "mock-device",
			"remote_host":   "127.0.0.1",
			"remote_port":   "1883",
		}),
		"/robot_api/v1/ota/getLatestPackage": envelope(map[string]any{
			"rom_version":      "1.0.0-mock",
			"mcu_version":      "1.0.0-mock",
			"whole_package_url": "http://127.0.0.1:8080/files/mock.zip",
			"md5":              "d41d8cd98f00b204e9800998ecf8427e",
			"upgrade_desc":     "mock package, nothing to install",
			"byte_size":        0,
			"update_time":      now,
			"package_collection": map[string]any{
				"rom_package":   packageItem,
				"mcu_a_package": packageItem,
			},
		}),
		"/robot_api/v1/device/upgrade/status": envelope(map[string]any{}),
		"/robot_api/v1/device/addRecord":      envelope(map[string]any{"record_id": 1}),
		"/robot_api/v1/cloudFile/getToken": envelope(map[string]any{
			"download_url":  "http://127.0.0.1:8080/files/photo.jpg",
			"file_key":      "mock/photo.jpg",
			"upload_domain": "http://127.0.0.1:8080",
			"upload_token":  "mock-upload-token",
		}),
		"/robot_api/v1/cloudFile/getSessionToken": envelope(map[string]any{
			"access_key":     "mock-access",
			"bucket":         "mock-bucket",
			"download_url":   "http://127.0.0.1:8080/files/photo.jpg",
			"expire_time":    3600,
			"object_key":     "mock/photo.jpg",
			"object_pre_key": "mock/",
			"secret_key":     "mock-secret",
			"session_token":  "mock-session",
		}),

		// Launcher display APIs. The launcher constants are still placeholders;
		// point those paths at these routes when you want the home screen filled.
		"/robot_api/v1/device/logo": envelope(map[string]any{
			"hello_logo":   "http://127.0.0.1:8080/files/hello.png",
			"desktop_logo": "http://127.0.0.1:8080/files/desktop.png",
		}),
		"/robot_api/v1/device/general": envelope(map[string]any{
			"wea": "cloudy", "wea_img": "yun", "tem": "18", "calender_total": 1,
		}),
		"/robot_api/v1/device/weather": envelope(map[string]any{
			"province": "Ile-de-France", "city": "Paris", "town": "Paris",
			"wea": "cloudy", "wea_img": "yun", "tem": "18",
			"tem_max": "21", "tem_min": "14", "win": "W", "win_speed": "3",
			"hourWeather": []map[string]string{
				{"hour": "12", "wea": "cloudy", "wea_img": "yun", "tem": "18"},
			},
		}),
		"/robot_api/v1/device/calendar": envelope(map[string]any{
			"event_total": 1, "has_more": false,
			"memo_list": []map[string]any{{
				"memo_title": "Mock meeting", "memo_time": now + 3600, "memo_time_label": "in 1 hour",
			}},
		}),
		"/robot_api/v1/device/countdown": envelope(map[string]any{
			"event_total": 1, "has_more": false,
			"event_list": []map[string]any{{
				"event_title": "Mock birthday", "event_time": now + 86400, "remain_days": 1,
			}},
		}),
		"/robot_api/v1/device/fans": map[string]any{
			"code": 0, "msg": "success",
			"data": []map[string]string{{
				"platform": "mock", "avatar": "", "nick_name": "emulator", "fans_count": "1",
			}},
		},
		"/robot_api/v1/device/clock": envelope(map[string]any{
			"clock_id": 1, "action": "add",
			"clock_info": map[string]any{
				"clock_id": 1, "clock_hour": 7, "clock_time": "07:00",
				"clock_title": "Mock alarm", "is_on": 1,
				"repeat_method": []int{1, 2, 3, 4, 5}, "repeat_method_label": "weekdays",
			},
		}),
		"/robot_api/v1/device/bindInfo": envelope(map[string]any{
			"bind_status": 1, "bind_time": now, "client_id": "mock-client", "user_id": 1,
		}),
		"/robot_api/v1/device/serverTime": envelope(map[string]any{"timestamp": now}),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if body, ok := routes[r.URL.Path]; ok {
			log.Printf("%s %s", r.Method, r.URL.Path)
			writeJSON(w, http.StatusOK, body)
			return
		}
		writeJSON(w, http.StatusNotFound, map[string]any{
			"code": 404, "msg": "no mock for " + r.URL.Path, "data": nil,
		})
	})

	addr := "127.0.0.1:8080"
	fmt.Println("mock robot API listening on http://" + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
