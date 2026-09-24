package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Req struct {
	AppId     string `json:"app_id"`     // appid configured in the app
	ApiKey    string `json:"api_key"`    // apikey configured in the app
	ApiSecret string `json:"api_secret"` // apisecret configured in the app
	Content   string `json:"content"`    // Request text for ai models
	Sn        string `json:"sn"`         // Machine unique code
}

type Data struct {
	Content string `json:"content"` // Return text content
	IsEnd   bool   `json:"is_end"`  // Is it over
}

type RespData struct {
	Code int    `json:"code"` // Status code
	Msg  string `json:"msg"`  // Return prompt information
	Data *Data  `json:"data"` // The returned structural data
}

func getSendMsg(code int, msg, content string, isEnd bool) (string, error) {
	d := &Data{
		Content: content,
		IsEnd:   isEnd,
	}
	respD := &RespData{
		Code: code,
		Msg:  msg,
		Data: d,
	}
	respB, err := json.Marshal(respD)
	return string(respB), err
}

var answerListAll = [][]string{
	[]string{
		"I am",
		"a private",
		"large language",
		"model",
		"built by Letianpai.",
	},
	[]string{
		"I do not know",
		"what you",
		"are talking about.",
		"I am only",
		"a private model.",
	},
	[]string{
		"Qin Yuan Chun · Snow",
		"Mao Zedong",
		"North country scene, a thousand miles locked in ice, ten thousand miles of whirling snow.",
		"On both sides of the Great Wall, only vastness remains; the great river, up and down, has suddenly lost its torrents.",
		"Mountains dance like silver snakes, the highlands race like wax elephants, wanting to compare height with heaven.",
		"On a clear day, see the red dress wrapped in white, extraordinarily charming.",
		"This land is so rich in beauty that it has made countless heroes bow.",
		"Pity the Qin emperor and Han Wu, a little lacking in literary grace; Tang Zong and Song Zu, somewhat short of romance.",
		"A generation's proud son, Genghis Khan, knew only to bend the bow and shoot the great eagle.",
		"All are past. For truly great men, look to this age.",
	},
	[]string{
		"To become a good product manager you need a larger self and a smaller self, and you may have to go through both before you find the reason behind the reason.",
		"What is the larger self? It means treating yourself as the CEO. In Band of Brothers the company commander's nickname is CEO.",
		"A product manager is responsible for a product. The title says manager, but the job is a general manager's job.",
		"Because you own the product, you keep coordinating departments and pushing people you do not manage: design, engineering, test, users, and marketing.",
		"The worries are no fewer than a general manager's. A good product manager acts as a CEO: responsible for execution and for the experience. Other people may not",
		"take you seriously, but you keep that larger self. Your title does not matter. You still carry the responsibility.",
		"So a good product manager does not care how many people they manage today or how high the title is. What matters is whether they can use the company's resources, by whatever means, and ship a good product.",
		"Hundreds of millions of users choosing the product is the real recognition.",
		"All else equal, a product used by hundreds of millions is better than a product used by a few million.",
		"You may say you want to be good at this but you are young and have no seniority. A product manager has to be confident.",
		"A good product manager keeps a larger self, and also keeps shrinking the self, even forgetting the self. That means",
		"forgetting yourself. The easiest mistake, one I still make, is to stop building for the user and start building for yourself, for coworkers, or for a boss.",
		"The smaller self, even the forgotten self, means standing in the user's place and splitting yourself in two, so you can feel what the user actually needs.",
	},
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set response header to specify the Content Type of SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	// to deal body
	fmt.Printf("request body: %s\n", body)
	var reqB Req
	err = json.Unmarshal(body, &reqB)
	fmt.Println("reqB:", reqB)
	answerList := answerListAll[0]
	content := reqB.Content

	fmt.Println("content:", content)

	// Simulate AI model matching answers
	if strings.Contains(content, "是谁") || strings.Contains(content, "who") ||
		strings.Contains(content, "大模型") || strings.Contains(content, "model") {
		answerList = answerListAll[0]
	} else if strings.Contains(content, "园春") || strings.Contains(content, "snow") {
		answerList = answerListAll[2]
	} else if strings.Contains(content, "产品经理") || strings.Contains(content, "product manager") {
		answerList = answerListAll[3]
	} else {
		answerList = answerListAll[1]
	}

	// Simulate real-time data push
	for idx, partAnswer := range answerList {
		isEnd := false
		if idx == len(answerList)-1 {
			isEnd = true
		}
		respStr, err := getSendMsg(0, "success", partAnswer, isEnd)
		if err != nil {
			respStr, _ = getSendMsg(1, "Error in obtaining return information", "", true)
		}

		SendEventMessage(w, "message", respStr)
		// Simulated delay added, real calls need to be removed
		time.Sleep(500 * time.Millisecond)
	}
}

// Send event messages
func SendEventMessage(w http.ResponseWriter, eventType, eventData string) {
	// Construct a message in Event Source format
	message := "event: " + eventType + "\n"
	message += "data: " + eventData + "\n\n"
	fmt.Println(message)
	// Write the message to the response body
	_, err := io.WriteString(w, message)
	if err != nil {
		// Error occurred, stop sending messages
		return
	}

	// Refresh response subject to ensure data is sent
	flusher, ok := w.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}

func main() {
	http.HandleFunc("/sse", handleSSE)

	// Start the web service and listen on thelocalhost:8012
	fmt.Println("server is listening : http://127.0.0.1:8012")
	if err := http.ListenAndServe(":8012", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
