package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

const (
	guildStatsURL = "https://api.wynncraft.com/public_api.php?action=guildStats&command=%s"
)

type Guild struct {
	Name            string    `json:"name"`
	Prefix          string    `json:"prefix"`
	Members         []Members `json:"members"`
	Xp              float64   `json:"xp"`
	Level           int       `json:"level"`
	Created         time.Time `json:"created"`
	CreatedFriendly string    `json:"createdFriendly"`
	Territories     int       `json:"territories"`
	Banner          Banner    `json:"banner"`
	Request         Request   `json:"request"`
}
type Members struct {
	Name           string    `json:"name"`
	UUID           string    `json:"uuid"`
	Rank           string    `json:"rank"`
	Contributed    int       `json:"contributed"`
	Joined         time.Time `json:"joined"`
	JoinedFriendly string    `json:"joinedFriendly"`
}
type Layers struct {
	Colour  string `json:"colour"`
	Pattern string `json:"pattern"`
}
type Banner struct {
	Base   string   `json:"base"`
	Tier   int      `json:"tier"`
	Layers []Layers `json:"layers"`
}
type Request struct {
	Timestamp int `json:"timestamp"`
	Version   int `json:"version"`
}

func GetGuildStats(guildName string) (*Guild, error) {
	res, err := get(fmt.Sprintf(guildStatsURL, url.PathEscape(guildName)))
	if err != nil {
		return nil, err
	}
	var ret Guild
	if err := json.Unmarshal([]byte(res), &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
