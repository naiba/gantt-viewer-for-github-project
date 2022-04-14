package model

type Gantt struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Start        string `json:"start,omitempty"`
	End          string `json:"end,omitempty"`
	Progress     int    `json:"progress,omitempty"`
	Dependencies string `json:"dependencies,omitempty"`
	Assignees    []struct {
		AvatarUrl string `json:"avatar_url,omitempty"`
		Login     string `json:"login,omitempty"`
	} `json:"assignees,omitempty"`
}
