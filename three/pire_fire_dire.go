package three

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func fetchMeatText() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func PireFireDire(input string) string {
	if input == "" {
		text, err := fetchMeatText()
		if err != nil {
			return "{}"
		}
		input = text
	}

	beefCount := make(map[string]int)

	input = strings.ToLower(input)

	input = strings.NewReplacer(",", " ", ".", " ").Replace(input)

	words := strings.Fields(input)

	for _, word := range words {
		beefCount[word]++
	}

	response := map[string]interface{}{
		"beef": beefCount,
	}

	jsonBytes, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return "{}"
	}

	return string(jsonBytes)
}
