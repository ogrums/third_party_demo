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
		"我是",
		"乐天派公司",
		"研发的",
		"私有化",
		"大模型",
	},
	[]string{
		"我不知道",
		"你说的",
		"是什么，",
		"我只是一个",
		"私有化大模型",
	},
	[]string{
		"沁园春·雪",
		"【作者】毛泽东 ",
		"北国风光，千里冰封，万里雪飘。",
		"望长城内外，惟余莽莽；大河上下，顿失滔滔。",
		"山舞银蛇，原驰蜡象，欲与天公试比高。",
		"须晴日，看红装素裹，分外妖娆。",
		"江山如此多娇，引无数英雄竞折腰。",
		"惜秦皇汉武，略输文采；唐宗宋祖，稍逊风骚。",
		"一代天骄，成吉思汗，只识弯弓射大雕。",
		"俱往矣，数风流人物，还看今朝。",
	},
	[]string{
		"要成为一个优秀的产品经理，必须在心里要有一个“大我”和一个“小我”，可能你要经历这两个过程的磨练，才能找到原因背后的原因。",
		"什么是“大我”？我的意思是，一个产品经理要把自己当成CEO。在《兄弟连》里面的连长的绰号就是CEO。",
		"产品经理要对一个产品负责，虽然挂的是经理的头衔，但行驶的是总经理的职责。",
		"因为要对产品负责，产品经理还要经常去协调很多部门，要去推动很多不归他管的人和事，比如要跟设计打交道，要跟技术打交道，要跟测试打交道，还要去了解用户的想法，还要去跟市场谈支持。",
		"产品经理要操的心，一点也不比一个总经理少。所以，我说一个优秀的产品经理首先要把自己当成一个CEO，既要负责执行、推动（Executive），也要负责用户体验（Experience）。可能别人不拿",
		"你当回事，但你自己心里的有一个“大我”。你对一个产品负责，你的title不重要，但是你一定要把这个责任担负起来。",
		"所以，优秀的产品经理不必在意今天管了多少人，不必在意自己的头衔是高是低，关键最是你能不能利用公司里的资源，不管采用什么手段，最后做出来一个好产品。",
		"上亿的用户选择使用你做出来的产品，就是对产品经理最大的认可。",
		"在其他相同的条件下，你做的产品有上亿用户用，他做的产品只有几百万人用，那么你就比他要优秀。",
		"可能你会说，我是想做一个优秀的产品经理，但我太年轻，没什么资历。我觉得，产品经理要非常自信，你要有这种气势。",
		"一个优秀的产品经理心中有“大我”，但同时还得不断地“小我”，甚至要“忘我”。这就是说，",
		"产品经理要忘掉自己。根据我的经验，产品经理最容易犯的几个错误，包括今天我都还在犯，就是产品做着做着，就不是给用户做产品，而是给自己做产品，给同事做产品，给领导做产品了。",
		"所以，产品经理心中要做到“小我”，甚至“忘我”，就是产品经理必须身临其境，把自己当成一个典型用户，让自己精神分裂，这样才能体会用户心中真正的想法和需求。",
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
	fmt.Printf("收到请求的body数据：%s\n", body)
	var reqB Req
	err = json.Unmarshal(body, &reqB)
	fmt.Println("reqB:", reqB)
	answerList := answerListAll[0]
	content := reqB.Content

	fmt.Println("content:", content)

	// Simulate AI model matching answers
	if strings.Contains(content, "是谁") || strings.Contains(content, "大模型") {
		answerList = answerListAll[0]
	} else if strings.Contains(content, "园春") {
		answerList = answerListAll[2]
	} else if strings.Contains(content, "产品经理") {
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
