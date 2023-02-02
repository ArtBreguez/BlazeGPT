package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nexidian/gocliselect"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
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

func main() {
	clearTerminal()
	menu := gocliselect.NewMenu("Selecione uma opção")
	menu.AddItem("Pegar ultima jogada ", "get_latest_play")
	choice := menu.Display()
	switch choice {
	case "get_latest_play":
		clearTerminal()
		jogadas, err := getBlazeData()
		checkErr(err, "Error getting blaze data")
		getChatGPTMessage(jogadas)
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
	bold := "\033[1m"
	reset := "\033[0m"
	//fmt.Println(bold + "Texto em negrito" + reset)
	fmt.Println("A provavel próxima jogada é", bold, text, reset)
	return text
}

func clearTerminal() {
	cmd, err := exec.Command("bash", "-c", "clear").Output()
	checkErr(err, "Erro ao limpar o terminal")
	fmt.Println(string(cmd))
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
