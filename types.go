package gorunemetrics

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type (
	Skill       int
	QuestStatus string

	Activity struct {
		Date    StringEncodedTime `json:"date"`
		Details string            `json:"details"`
		Text    string            `json:"text"`
	}

	SkillValue struct {
		ID    Skill `json:"id"`
		Level int   `json:"level"`
		Rank  int   `json:"rank"`
		XP    int   `json:"xp"`
	}

	PlayerProfile struct {
		Activities       []*Activity       `json:"activities"`
		CombatLevel      int               `json:"combatlevel"`
		LoggedIn         StringEncodedBool `json:"loggedIn"`
		MagicXP          int               `json:"magic"`
		MeleeXP          int               `json:"melee"`
		Name             string            `json:"name"`
		QuestsComplete   int               `json:"questscomplete"`
		QuestsNotStarted int               `json:"questsnotstarted"`
		QuestsStarted    int               `json:"queststarted"`
		RangedXP         int               `json:"ranged"`
		Rank             StringEncodedInt  `json:"rank"`
		SkillValues      []*SkillValue     `json:"skillvalues"`
		TotalSkill       int               `json:"totalskill"`
		TotalXP          int               `json:"totalxp"`
	}

	PlayerQuestStatus struct {
		Difficulty   int         `json:"difficulty"`
		Members      bool        `json:"members"`
		QuestPoints  int         `json:"questPoints"`
		Status       QuestStatus `json:"status"`
		Title        string      `json:"title"`
		UserEligible bool        `json:"userEligible"`
	}

	StringEncodedBool bool
	StringEncodedInt  int
	StringEncodedTime time.Time
)

const (
	Attack Skill = iota
	Defence
	Strength
	Constitution
	Ranged
	Prayer
	Magic
	Cooking
	Woodcutting
	Fletching
	Fishing
	Firemaking
	Crafting
	Smithing
	Mining
	Herblore
	Agility
	Thieving
	Slayer
	Farming
	Runecrafting
	Hunter
	Construction
	Summoning
	Dungeoneering
	Divination
	Invention
	Archaeology
	Necromancy
)

const (
	Completed  QuestStatus = "COMPLETED"
	Started    QuestStatus = "STARTED"
	NotStarted QuestStatus = "NOT_STARTED"
)

func (p *PlayerProfile) String() string {
	return fmt.Sprintf(
		"{activities: %v, "+
			"combatlevel: %v, "+
			"loggedIn: %v, "+
			"magic: %v, "+
			"melee: %v, "+
			"name: %v, "+
			"questscomplete: %v, "+
			"questsnotstarted: %v, "+
			"queststarted: %v, "+
			"ranged: %v, "+
			"rank: %v, "+
			"skillvalues: %v, "+
			"totalskill: %v, "+
			"totalxp: %v}",
		p.Activities,
		p.CombatLevel,
		p.LoggedIn,
		p.MagicXP,
		p.MeleeXP,
		p.Name,
		p.QuestsComplete,
		p.QuestsNotStarted,
		p.QuestsStarted,
		p.RangedXP,
		p.Rank,
		p.SkillValues,
		p.TotalSkill,
		p.TotalXP,
	)
}

func (a *Activity) String() string {
	return fmt.Sprintf("{date: %v, "+
		"details: %v, "+
		"text: %v}",
		a.Date,
		a.Details,
		a.Text,
	)
}

func (s *SkillValue) String() string {
	return fmt.Sprintf("{id: %v, "+
		"level: %v, "+
		"rank: %v, "+
		"xp: %v}",
		s.ID,
		s.Level,
		s.Rank,
		s.XP)
}

func (p *PlayerQuestStatus) String() string {
	return fmt.Sprintf("{difficulty: %v, "+
		"members: %v, "+
		"questPoints: %v, "+
		"status: %v, "+
		"title: %v, "+
		"userEligible: %v}",
		p.Difficulty,
		p.Members,
		p.QuestPoints,
		p.Status,
		p.Title,
		p.UserEligible,
	)
}

func (s Skill) String() string {
	switch s {
	case Attack:
		return "Attack"
	case Defence:
		return "Defence"
	case Strength:
		return "Strength"
	case Constitution:
		return "Constitution"
	case Ranged:
		return "Ranged"
	case Prayer:
		return "Prayer"
	case Magic:
		return "Magic"
	case Cooking:
		return "Cooking"
	case Woodcutting:
		return "Woodcutting"
	case Fletching:
		return "Fletching"
	case Fishing:
		return "Fishing"
	case Firemaking:
		return "Firemaking"
	case Crafting:
		return "Crafting"
	case Smithing:
		return "Smithing"
	case Mining:
		return "Mining"
	case Herblore:
		return "Herblore"
	case Agility:
		return "Agility"
	case Thieving:
		return "Thieving"
	case Slayer:
		return "Slayer"
	case Farming:
		return "Farming"
	case Runecrafting:
		return "Runecrafting"
	case Hunter:
		return "Hunter"
	case Construction:
		return "Construction"
	case Summoning:
		return "Summoning"
	case Dungeoneering:
		return "Dungeoneering"
	case Divination:
		return "Divination"
	case Invention:
		return "Invention"
	case Archaeology:
		return "Archaeology"
	case Necromancy:
		return "Necromancy"
	default:
		return ""
	}
}

func (s *StringEncodedBool) UnmarshalJSON(data []byte) error {
	stringValue := ""

	if err := json.Unmarshal(data, &stringValue); err != nil {
		return err
	}

	*s = stringValue == "true"

	return nil
}

func (s *StringEncodedInt) UnmarshalJSON(data []byte) error {
	stringValue := ""

	if err := json.Unmarshal(data, &stringValue); err != nil {
		return err
	}

	stringValue = strings.ReplaceAll(stringValue, ",", "")

	int64Value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return err
	}

	*s = StringEncodedInt(int64Value)

	return nil
}

func (s *StringEncodedTime) UnmarshalJSON(data []byte) error {
	stringValue := ""

	if err := json.Unmarshal(data, &stringValue); err != nil {
		return err
	}

	timeValue, err := time.Parse("02-Jan-2006 15:04", stringValue)
	if err != nil {
		return err
	}

	*s = StringEncodedTime(timeValue)

	return nil
}
