package parscejson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type players struct {
	Player []player
}

type player struct {
	Id     int
	Name   string
	Hp     int
	Mp     int
	Str    int
	Int    int
	Def    int
	Res    int
	Spd    int
	Luck   int
	Race   int
	Class  int
	Rarity int
}

type bosses struct {
	Boss []boss
}
type boss struct {
	Id     int
	Name   string
	Hp     int
	Mp     int
	Str    int
	Int    int
	Def    int
	Res    int
	Spd    int
	Luck   int
	Race   int
	Class  int
	Rarity int
}

type Enemies struct {
	Monster []Enemy `json:"Monster"`
}

type Enemy struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Hp     int    `json:"Hp"`
	Mp     int    `json:"Mp"`
	Str    int    `json:"Str"`
	Int    int    `json:"Int"`
	Def    int    `json:"Def"`
	Res    int    `json:"Res"`
	Spd    int    `json:"Spd"`
	Luck   int    `json:"Luck"`
	Race   int    `json:"Race"`
	Class  int    `json:"Class"`
	Rarity int    `json:"Rarity"`
}

func ParsePlayer(Perso string) player {
	PlayerFile, err := os.Open("JsonFile/players.json")
	if err != nil {
		fmt.Println(err)
	}

	defer PlayerFile.Close()

	byteValue, err := io.ReadAll(PlayerFile)
	if err != nil {
		fmt.Println(err)
	}
	var Players players
	var result player
	json.Unmarshal(byteValue, &Players)
	for i := 0; i < len(Players.Player); i++ {
		if Players.Player[i].Name == Perso {
			result = Players.Player[i]
		}
	}
	return result
}

func ParseBoss(Choice string) boss {
	BossFile, err := os.Open("JsonFile/bosses.json")
	if err != nil {
		fmt.Println(err)
	}
	defer BossFile.Close()
	byteValue, err := io.ReadAll(BossFile)
	if err != nil {
		fmt.Println(err)
	}
	var Bosses bosses
	var result boss
	json.Unmarshal(byteValue, &Bosses)
	for i := 0; i < len(Bosses.Boss); i++ {
		if Bosses.Boss[i].Name == Choice {
			result = Bosses.Boss[i]
		}
	}
	return result
}

func ParseEnemies(Choice string) Enemy {
	EnemiesFile, err := os.Open("JsonFile/enemies.json")
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
	}
	defer EnemiesFile.Close()
	byteValue, err := io.ReadAll(EnemiesFile)
	if err != nil {
		fmt.Errorf("error reading file: %w", err)
	}

	var Enemies Enemies
	if err := json.Unmarshal(byteValue, &Enemies); err != nil {
		fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	var enemie Enemy
	for _, e := range Enemies.Monster {
		if (e.Name) == Choice {
			enemie = e
			return enemie
		}

	}
	return enemie
}
