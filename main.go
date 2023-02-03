package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Records struct {
	Id         string `json:"id"`
	Created_at string `json:"created_at"`
	Color      string `json:"color"`
	Roll       int    `json:"roll"`
}

type TotalPages struct {
	Total_pages int       `json:"total_pages"`
	Records     []Records `json:"records"`
}

var lastHash [32]byte

func main() {
	for {
		jogadas, err := getBlazeData()
		if err != nil {
			fmt.Println("Error getting blaze data:", err)
			continue
		}

		hash := sha256.Sum256([]byte(fmt.Sprintf("%v", jogadas)))
		if hash == lastHash {
			fmt.Println("Payload unchanged")
		} else {
			lastHash = hash
			text := getChatGPTMessage(jogadas)
			sendMessageToTelegramChannel(text)
		}

		time.Sleep(2 * time.Second)
	}
}

func getBlazeData() ([]string, error) {
	var colors []string
	data, err := http.Get("https://blaze.com/api/roulette_games/history")
	checkErr(err, "Error getting data from blaze.com")
	defer data.Body.Close()
	var result TotalPages
	err = json.NewDecoder(data.Body).Decode(&result)
	checkErr(err, "Error decoding blaze.com data")
	for _, v := range result.Records {
		colors = append(colors, v.Color)
	}
	return colors, err
}

func getChatGPTMessage(jogadas []string) string {
	result := strings.Join(jogadas, ", ")
	url := "https://api.openai.com/v1/completions"
	payload := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "Baseado nessa sequencia do jogo Double da Blaze [" + result + "] qual a possivel nova cor? Reponda da maneira mais curta possivel.",
		"max_tokens":  7,
		"temperature": 0.6,
	}

	jsonValue, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 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")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response map[string]interface{}
	json.Unmarshal(body, &response)

	choices := response["choices"].([]interface{})
	firstChoice := choices[0].(map[string]interface{})
	text := firstChoice["text"].(string)
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, ".", "", -1)
	return text
}

func sendMessageToTelegramChannel(text string) {
	var emoji string
	token := "5891096865:AAFHsDFzfFfgFDBDIUe5drg68OKsDOu9HUw"
	chatID := "-1001809111657"
	if text == "Black" {
		emoji = `⚫`
	} else if text == "Red" {
		emoji = `🔴`
	} else if text == "White" {
		emoji = `⚪`
	}
	message := "A próxima jogada é " + text + " " + emoji

	encodedMessage := url.QueryEscape(message)
	url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chatID + "&text=" + encodedMessage

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func checkErr(err error, msg ...string) {
	if err == nil {
		return
	}
	var output string
	if len(msg) != 0 {
		output = msg[0] + " "
	}
	output += err.Error()
	log.Fatalln(output)
}
