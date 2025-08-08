package main

import (
	"context"
	"encoding/json"
	apiglue "gin-sevalla-app/api_glue"
	"strings"
)

func clean_ai_json_response(json_string string) string {
	trimmed_json_string := strings.TrimSpace(json_string)
	if strings.HasPrefix(trimmed_json_string, "```json") {
		assert(strings.HasSuffix(trimmed_json_string, "```"), "```json starting string is not closed using ```")
		trimmed_json_string = trimmed_json_string[len("```json"):]
		trimmed_json_string = trimmed_json_string[:len(trimmed_json_string)-len("```")]
	}
	return trimmed_json_string
}

var ctx = context.Background()

//	var client, ai_api_initialization_err = genai.NewClient(ctx, &genai.ClientConfig{
//		APIKey:  os.Getenv("GEMINI_API_KEY"),
//		Backend: genai.BackendGeminiAPI,
//	})
var response_json_example, _ = json.Marshal(Res{})

// func init() {
// 	if ai_api_initialization_err != nil {
// 		log.Fatal(ai_api_initialization_err)
// 	}
// }

type Live_Scout struct {
	Links        map[string]struct{} `json:"links"`
	Notes        string              `json:"notes"`
	As_Of        string              `json:"as_of"`
	Scouted      bool                `json:"scouted"`
	To_Scout_For string              `json:"to_scout_for"`
}
type Res struct {
	Goal_Achieved              string   `json:"goal_achieved"`
	Links                      []string `json:"links"`
	As_Of                      string   `json:"as_of"`
	Would_Like_to_Update_Notes string   `json:"would_like_to_update_notes"`
	New_Notes                  string   `json:"new_notes"`
	Assosiated_json            string   `json:"assosiated_json"`
	Completion_Percentage      string   `json:"completion_percentage"`
	Predicted_timeline         string   `json:"predicted_timeline"`
}

// func (live_scout *Live_Scout) Scout() Res {
// 	previous_context := if_else(live_scout.Notes != "", "these are the previous notes, would you like to add to it? new notes: "+live_scout.Notes, "there are no previous notes, would you like to some to be used later as context?")

// 	result, err := client.Models.GenerateContent(
// 		ctx,
// 		"gemini-2.5-flash",
// 		genai.Text("for the following response only with json: in yes or no terms. <"+live_scout.To_Scout_For+"> please provide all assosiated links. please provide only links that actually work. also add notes regarding how close you think you are towards achieving the goal. "+previous_context+" please use the following format for your response: "+string(response_json_example)),
// 		nil,
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cleaned := clean_ai_json_response(result.Text())
// 	res := Res{}
// 	json.Unmarshal([]byte(cleaned), &res)
// 	//////
// 	if res.Goal_Achieved == "yes" {
// 		live_scout.Scouted = true
// 		live_scout.As_Of = res.As_Of
// 	}
// 	if res.Would_Like_to_Update_Notes == "yes" {
// 		live_scout.Notes = res.New_Notes
// 	}
// 	for _, link := range res.Links {
// 		live_scout.Links[link] = struct{}{}
// 	}
// 	return res
// }

var live_scouts = apiglue.NewServerState([]Live_Scout{
	{
		Links:        make(map[string]struct{}),
		To_Scout_For: "tell me when there is an ai truck driver in the state of florida",
	},
	{
		Links:        make(map[string]struct{}),
		To_Scout_For: "find me a good reliable car with a resonable at signing that i can rent for about 200-300 a month that has a decent insurence policy for 21-25 year olds.",
	},
	{
		Links:        make(map[string]struct{}),
		To_Scout_For: "let me know if there are any water parks that open up in a place that's close a bar, this place could be anywhere that is in the vecinity of kosher food, anywhere in america",
	},
	{
		Links:        make(map[string]struct{}),
		To_Scout_For: "is there a new decent laptop that i can get for under 200$ with 16 gb of ram.",
	},
})
