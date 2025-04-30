package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	Name string `json:"name"`
}

type GitHubActivity struct {
	Type      string `json:"type"`
	Repo      Repo   `json:"repo"`
	CreatedAt string `json:"created_at"`
	Payload   struct {
		Action  string `json:"action"`
		Ref     string `json:"ref"`
		RefType string `json:"ref_type"`

		Commits []struct {
			Message string `json:"message"`
		} `json:"commits"`
	} `json:"payload"`
}

func FetchApi(username string) ([]GitHubActivity, error) {
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("status code error: %d user not found", res.StatusCode)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	activities := []GitHubActivity{}
	err = json.NewDecoder(res.Body).Decode(&activities)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func GetActivity(username string) error {
	activities, err := FetchApi(username)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s activity(s)\n\n", username)

	for _, activity := range activities {
		switch activity.Type {
		case "PushEvent":
			fmt.Printf("Pushed %v commit(s) to %s\n", len(activity.Payload.Commits), activity.Repo.Name)

		case "IssuesEvent":
			fmt.Printf("%v an issue in %s\n", activity.Payload.Action, activity.Repo.Name)

		case "WatchEvent":
			fmt.Printf("Starred %s", activity.Repo.Name)
		case "ForkEvent":
			fmt.Printf("Forked %s", activity.Repo.Name)
		case "CreateEvent":
			fmt.Printf("Created %s in %s", activity.Payload.RefType, activity.Repo.Name)
		default:
			fmt.Printf("%s in %s", activity.Type, activity.Repo.Name)
		}
	}

	return nil
}
