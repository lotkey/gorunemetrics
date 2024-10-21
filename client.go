package gorunemetrics

type (
	PlayerProfile struct{}

	PlayerQuestStatus struct{}

	Client interface {
		GetProfile(playerName string) (PlayerProfile, error)
		GetQuests(playerName string) (PlayerQuestStatus, error)
	}
)

func NewClient() Client {
	return nil
}
