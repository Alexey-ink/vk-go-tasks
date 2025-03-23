package main

import "fmt"

type Player struct {
	Name        string
	CurrentRoom *Room
	Inventory   []string
	BackpackOn  bool
}

func (p *Player) String() string {
	var ItemStr string
	for i, v := range p.Inventory {
		if i > 0 {
			ItemStr += ", "
		}
		ItemStr += v
	}
	return fmt.Sprintf("Name of player: %s \nCurrent room: %v \nInventory: [%v]\n", p.Name, p.CurrentRoom, ItemStr)
}

func (p *Player) CheckItem(item string) bool {
	for _, v := range p.Inventory {
		if v == item {
			return true
		}
	}
	return false
}
