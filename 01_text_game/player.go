package main

import (
	"fmt"
	"strings"
)

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

func (p *Player) lookAround() string {
	var builder strings.Builder

	if !p.CurrentRoom.IsEmpty() {

		if p.CurrentRoom.Name == "кухня" {
			builder.WriteString("ты находишься на кухне, ")
		}

		keys := p.CurrentRoom.getNotEmptyFurnitureKeys()

		for _, fur := range keys {
			builder.WriteString("на " + fur + "е: ") // на столе, на стуле

			items := p.CurrentRoom.getItemsInFurniture(fur)

			for _, item := range items {
				builder.WriteString(item + ", ")
			}
		}

		if p.CurrentRoom.Name == "кухня" {
			builder.WriteString("надо ")
			if !p.BackpackOn {
				builder.WriteString("собрать рюкзак и ")
			}
			builder.WriteString("идти в универ. ")
		} else {
			resultStr := builder.String()
			if len(resultStr) >= 2 {
				resultStr = resultStr[:len(resultStr)-2]
			}
			builder.Reset()
			builder.WriteString(resultStr)
			builder.WriteString(". ")
		}
	} else {
		builder.WriteString("пустая комната. ")
	}

	builder.WriteString(p.CurrentRoom.getNearbyRoomsDescription())
	return builder.String()
}

func (p *Player) goRoom(args []string) string {
	var result string

	for _, room := range p.CurrentRoom.NearbyRooms {
		if args[0] == room.Name {
			if room.Name == "улица" && !p.CurrentRoom.DoorOpen {
				return "дверь закрыта"
			}
			p.CurrentRoom = room
			result += room.Description
			result += p.CurrentRoom.getNearbyRoomsDescription()
			return result

		}
	}

	return result + "нет пути в " + args[0]
}

func (p *Player) take(args []string) string {
	if p.BackpackOn {
		if p.CurrentRoom.CheckItem(args[0]) {
			p.Inventory = append(p.Inventory, args[0])
			p.CurrentRoom.deleteItem(args[0])
			return fmt.Sprintf(("предмет добавлен в инвентарь: %s"), args[0])
		} else {
			return "нет такого"
		}
	}

	if args[0] == "рюкзак" {
		p.BackpackOn = true
		p.CurrentRoom.deleteItem(args[0])
		return "вы надели: рюкзак"
	} else {
		return "некуда класть"
	}
}

func (p *Player) apply(args []string) string {
	if !p.CheckItem(args[0]) {
		return "нет предмета в инвентаре - " + args[0]
	}

	if args[0] != "ключи" || p.CurrentRoom.Name != "коридор" {
		return "не к чему применить"
	}

	if args[1] != "дверь" {
		return "не к чему применить"
	}

	p.CurrentRoom.DoorOpen = true
	return "дверь открыта"
}
