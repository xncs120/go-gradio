package gradio

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func (gr *Gradio) getEventId(userInput string, args ...interface{}) string {
	if gr.systemPrompt != "" {
		userInput = gr.systemPrompt + "\n\n" + userInput
	}
	jsonData := map[string]interface{}{
		"data": append([]interface{}{userInput}, args...),
	}
	payload, _ := json.Marshal(jsonData)

	client := &http.Client{}
	req, err := http.NewRequest("POST", gr.url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	if gr.hfToken != "" {
		req.Header.Set("Authorization", "Bearer "+gr.hfToken)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	eventID, ok := response["event_id"].(string)
	if !ok {
		log.Println("No response from api")
		return ""
	}

	return eventID
}

func (gr *Gradio) getResponse(eventID string) string {
	newUrl := fmt.Sprintf("%s/%s", gr.url, eventID)

	client := &http.Client{}
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	if gr.hfToken != "" {
		req.Header.Set("Authorization", "Bearer "+gr.hfToken)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	var finalResponse string
	re := regexp.MustCompile(`data:\s*(\[.*\])`)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			jsonData := matches[1]

			var response []interface{}
			if err := json.Unmarshal([]byte(jsonData), &response); err == nil {
				if len(response) > 0 {
					if firstElement, ok := response[0].(string); ok {
						finalResponse = firstElement
					}
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
		return ""
	}

	return finalResponse
}

func (gr *Gradio) ChatCompletion(userInput string) string {
	eventId := gr.getEventId(userInput, gr.maxTokens, gr.temperature, gr.topP)
	return gr.getResponse(eventId)
}
