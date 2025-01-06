package main

import (
	"encoding/json"
	"fmt"
)

type ItemList []Item

type Item struct {
	Id             uint               `json:"id"`
	GameId         uint               `json:"game_id"`
	Name           string             `json:"name"`
	DisplayName    string             `json:"display_name"`
	Price          uint               `json:"price"`
	TotalPrice     uint               `json:"total_price"`
	SlotType       string             `json:"slot_type"`
	Rarity         string             `json:"rarity"`
	AggressionType string             `json:"aggression_type"`
	HeroClass      string             `json:"hero_class"`
	RequiredLevel  uint               `json:"required_level"`
	Stats          map[string]float64 `json:"stats"`
	Effects        []Effect           `json:"effects"`
	Requirements   []string           `json:"requirements"`
	BuildPaths     []string           `json:"build_paths"`
}

type HeroList []Hero

type Hero struct {
	Id          uint                `json:"id"`
	GameId      uint                `json:"game_id"`
	Name        string              `json:"name"`
	DisplayName string              `json:"display_name"`
	Image       string              `json:"image"`
	Stats       []uint              `json:"stats"`
	Classes     []string            `json:"classes"`
	Roles       []string            `json:"roles"`
	Abilities   []Ability           `json:"abilities"`
	BaseStats   map[string]HeroStat `json:"base_stats"`
}

// Hero stats are expressed as a list of values per level
type HeroStat []float64

type Ability struct {
	DisplayName     string    `json:"display_name"`
	Image           string    `json:"image"`
	GameDescription string    `json:"game_description"`
	MenuDescription string    `json:"menu_description"`
	Cooldown        []float64 `json:"cooldown"`
	Cost            []float64 `json:"cost"`
	Key             string    `json:"key"`
}

type Effect struct {
	Name            string `json:"name"`
	Active          bool   `json:"active"`
	Condition       string `json:"condition"`
	Cooldown        string `json:"cooldown"`
	MenuDescription string `json:"menu_description"`
}

type OmedaService struct {
	client *OmedaClient
}

func NewOmedaService(client *OmedaClient) *OmedaService {
	return &OmedaService{
		client: client,
	}
}

func (o *OmedaService) GetAllItems() (ItemList, error) {
	path := "/items.json"
	items := ItemList{}

	err := o.fetchAndUnmarshal(path, &items)

	return items, err
}

func (o *OmedaService) GetItem(name string) (Item, error) {
	path := fmt.Sprintf("/items/%s.json", name)
	item := Item{}

	err := o.fetchAndUnmarshal(path, &item)

	return item, err
}

func (o *OmedaService) GetAllHeroes() (HeroList, error) {
	path := "/heroes.json"
	heroes := HeroList{}

	err := o.fetchAndUnmarshal(path, &heroes)

	return heroes, err
}

func (o *OmedaService) GetHero(name string) (Hero, error) {
	path := fmt.Sprintf("/heroes/%s.json", name)
	hero := Hero{}

	err := o.fetchAndUnmarshal(path, &hero)

	return hero, err
}

func (o *OmedaService) fetchAndUnmarshal(path string, model any) error {
	resp, err := o.client.Get(path)
	if err != nil {
		return fmt.Errorf("fetching: %w", err)
	}

	err = json.Unmarshal(resp, model)
	if err != nil {
		return fmt.Errorf("unmarshalling model: %w", err)
	}

	return nil
}
