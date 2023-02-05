package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
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

type Config struct {
	Channel string
	ChatID  string
	Blaze   string
	ChatGPT string
	Token   string
}

var lastHash [32]byte

func main() {
	config, err := readEnv()
	checkErr(err, "Error reading config file")
	for {
		jogadas, err := getBlazeData(config.Blaze)
		if err != nil {
			fmt.Println("Error getting blaze data:", err)
			continue
		}

		hash := sha256.Sum256([]byte(fmt.Sprintf("%v", jogadas)))
		if hash == lastHash {
		} else {
			lastHash = hash
			text := getChatGPTMessage(jogadas, config)
			sendMessageToTelegramChannel(text, config)
		}

		time.Sleep(2 * time.Second)
	}
}

func getBlazeData(endpoint string) ([]string, error) {
	var colors []string
	data, err := http.Get(endpoint)
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

func getChatGPTMessage(jogadas []string, config Config) string {
	result := strings.Join(jogadas, ", ")
	url := config.ChatGPT
	payload := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      "Baseado nessa sequencia do jogo Double da Blaze [" + result + "] qual a possivel nova cor? Reponda da maneira mais curta possivel. Uma observação é que a cor 'White' é a que tem menos probabilidade de sair, com a vermelha e preta com a mesma probabilidade. Responda em ingles. Se a probabilidade de sair a cor for menor que 90% retorne None.",
		"max_tokens":  7,
		"temperature": 1,
	}

	jsonValue, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", config.Token)

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

func sendMessageToTelegramChannel(text string, config Config) {
	var emoji string
	token := config.Channel
	chatID := config.ChatID
	if text == "Black" {
		emoji = `⚫`
	} else if text == "Red" {
		emoji = `🔴`
	} else if text == "White" {
		emoji = `⚪`
	} else {
		return
	}
	message := "A próxima jogada é " + text + " " + emoji

	encodedMessage := url.QueryEscape(message)
	url := "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + chatID + "&text=" + encodedMessage

	// create the log file in the "logs" directory
	file, err := os.OpenFile("logs/requests.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	log.SetOutput(file)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	logger := logrus.New()
	file, err = os.OpenFile("logs/requests.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		log.Fatalln("Failed to log to file, using default stderr")
	}
	logger.Info(string(body))
}

func readEnv() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	checkErr(err, "Error reading config file")
	var config Config
	err = viper.Unmarshal(&config)
	checkErr(err, "Unable to decode into struct")
	return config, err
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
