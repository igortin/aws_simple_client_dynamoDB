package main

import "log"

func main() {

	client, err := New("private", "us-east-1")
	log.Println("Test")
	if err != nil {
		log.Println(err.Error())
		return
	}

	rawJSONString := []byte(`{
        "instancesSet": {
            "items": [
                {
					"instanceId": "i-0abdee38da9e51098",
					"state": "Good"
                }
            ]
        }
	}`)

	event := &Event{
		EventID:   "777",
		EventInfo: "API",
		Body:      rawJSONString,
	}

	puItemOutput, err := client.PutItem(event)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(*puItemOutput)
}
