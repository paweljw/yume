package game

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
)

type Race int

const (
	Human Race = iota
	Dwarf
	Elf
)

type Player struct {
	Name string
	Password string
	Race Race
	Flags map[string]bool
	Version uint64
}

func (player *Player) SetPassword(unsecurePassword string) {
	hasher := sha256.New()
	hasher.Write([]byte(unsecurePassword))
	player.Password = hex.EncodeToString(hasher.Sum(nil))
}

func (player *Player) ComparePassword(unsecurePassword string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(unsecurePassword))
	return player.Password == hex.EncodeToString(hasher.Sum(nil))
}

func (player *Player) SaveToFile() error {
	jsonRepr, err := json.MarshalIndent(player, "", "  ")

	if err != nil {
		return err
	}

	filename := filenameFromPlayerName(player.Name)

	err = os.WriteFile(filename, jsonRepr, 0644)

	return err
}

func (player *Player) IsSaveable() bool {
	return PlayerFileExists(player.Name)
}

func (player *Player) HasFlag(name string) bool {
	return player.Flags[name] || false
}

func (player *Player) SetFlag(name string, value bool) {
	if player.Flags == nil {
		player.Flags = make(map[string]bool)
	}

	player.Flags[name] = value
}

func (player *Player) NameEquals(name string) bool {
	return strings.ToLower(player.Name) == strings.ToLower(name)
}

func LoadPlayerFromFile(name string) (*Player, error) {
	player := Player{}

	filename := filenameFromPlayerName(name)

	jsonRepr, err := os.ReadFile(filename)

	if err != nil {
		return &player, err
	}

	err = json.Unmarshal(jsonRepr, &player)

	return &player, err
}

func PlayerFileExists(name string) bool {
	filename := filenameFromPlayerName(name)

	if _, err := os.Stat(filename); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func filenameFromPlayerName(name string) string {
	return "./ugc/" + strings.ToLower(name) + ".json"
}
