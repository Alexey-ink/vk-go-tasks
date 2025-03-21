package main

import "fmt"

type Player struct {
	Name        string
	CurrentRoom *Room
	Inventory   *[]string
	BackpackOn  bool
}

func (p Player) String() string {
	var ItemStr string
	for i, v := range *p.Inventory {
		if i > 0 {
			ItemStr += ", "
		}
		ItemStr += v
	}
	return fmt.Sprintf("Name of player: %s \nCurrent room: %v \nInventory: [%v]\n", p.Name, p.CurrentRoom, ItemStr)
}

func (p Player) Take(item string) string {
	if p.BackpackOn {
		*p.Inventory = append(*p.Inventory, item)
		return fmt.Sprintf(("предмет добавлен в инвентарь: %s"), item)
	} else if item == "рюкзак" {
		p.BackpackOn = true
		return "вы надели: рюкзак"
	} else {
		return "некуда класть"
	}
}
