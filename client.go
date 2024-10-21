package gorunemetrics

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	wrappedhttp "github.com/lotkey/gorunemetrics/internal/http"
)

type (
	// Client represents a RuneMetrics client.
	Client interface {
		// GetProfile returns profile data given a username.
		GetProfile(playerName string) (*PlayerProfile, error)
		// GetQuests returns quest statuses given a username.
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

var (
	// ErrUnexpectedStatusCode is returned when an unexpected status code is
	// encountered from RuneMetrics.
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	// ErrMissingPlayerData is returned when player data is missing in the
	// RuneMetrics response.
	ErrMissingPlayerData = errors.New("missing player data")
)

// NewClient creates a new RuneMetrics API client given an HTTP client.
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

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to GET player profile: %v: %w", response.Status, err)
	}

	profile := &PlayerProfile{}

	if err := json.NewDecoder(response.Body).Decode(profile); err != nil {
		return nil, fmt.Errorf("failed to decode player profile: %w", err)
	}

	if len(profile.SkillValues) == 0 || profile.Activities == nil {
		return nil, fmt.Errorf("failed to find player profile data: %w", ErrMissingPlayerData)
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

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to GET player quest status: %v: %w", response.Status, err)
	}

	quests := struct {
		Quests []*PlayerQuestStatus `json:"quests"`
	}{}

	if err := json.NewDecoder(response.Body).Decode(&quests); err != nil {
		return nil, fmt.Errorf("failed to decode player quest status: %w", err)
	}

	if quests.Quests == nil {
		return nil, fmt.Errorf("failed to find player quest status data: %w", ErrMissingPlayerData)
	}

	return quests.Quests, nil
}
