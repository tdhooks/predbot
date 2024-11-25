package main

type ItemList []Item

type Item struct {
	Id             uint              `json:"id"`
	GameId         uint              `json:"game_id"`
	Name           string            `json:"name"`
	DisplayName    string            `json:"display_name"`
	Price          uint              `json:"price"`
	TotalPrice     uint              `json:"total_price"`
	SlotType       string            `json:"slot_type"`
	Rarity         string            `json:"rarity"`
	AggressionType string            `json:"aggression_type"`
	HeroClass      string            `json:"hero_class"`
	RequiredLevel  string            `json:"required_level"`
	Stats          map[string]string `json:"stats"`
	Effects        []ItemEffect      `json:"effects"`
	Requirements   []string          `json:"requirements"`
	BuildPaths     []string          `json:"build_paths"`
}

type ItemEffect struct {
	Name            string `json:"name"`
	Active          bool   `json:"active"`
	Condition       string `json:"condition"`
	Cooldown        string `json:"cooldown"`
	MenuDescription string `json:"menu_description"`
}
