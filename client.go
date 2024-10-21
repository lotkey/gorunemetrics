package gorunemetrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	wrappedhttp "github.com/lotkey/gorunemetrics/internal/http"
)

type (
	Client interface {
		GetProfile(playerName string) (*PlayerProfile, error)
		GetQuests(playerName string) ([]*PlayerQuestStatus, error)
	}

	client struct {
		httpClient wrappedhttp.Client
	}
)

const (
	baseAPIURL    = "https://apps.runescape.com/runemetrics"
	profileAPIURL = baseAPIURL + "/profile/profile"
	questsAPIURL  = baseAPIURL + "/quests"
)

func NewClient(httpClient *http.Client) Client {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) GetProfile(playerName string) (*PlayerProfile, error) {
	defer c.httpClient.CloseIdleConnections()

	response, err := c.httpClient.Get(fmt.Sprintf("%s?user=%s", profileAPIURL, url.QueryEscape(playerName)))
	if err != nil {
		return nil, fmt.Errorf("failed to GET player profile: %w", err)
	}

	defer response.Body.Close()

	profile := &PlayerProfile{}

	if err := json.NewDecoder(response.Body).Decode(profile); err != nil {
		return nil, fmt.Errorf("failed to decode player profile: %w", err)
	}

	return profile, nil
}

func (c *client) GetQuests(playerName string) ([]*PlayerQuestStatus, error) {
	defer c.httpClient.CloseIdleConnections()

	response, err := c.httpClient.Get(fmt.Sprintf("%s?user=%s", questsAPIURL, url.QueryEscape(playerName)))
	if err != nil {
		return nil, fmt.Errorf("failed to GET player quest status: %w", err)
	}

	defer response.Body.Close()

	quests := struct {
		Quests []*PlayerQuestStatus `json:"quests"`
	}{}

	if err := json.NewDecoder(response.Body).Decode(&quests); err != nil {
		return nil, fmt.Errorf("failed to decode player quest status: %w", err)
	}

	return quests.Quests, nil
}
