package main

import (
	"fmt"
	"strings"
)

var COMMANDS = map[string]func(*Player, []string) string{
	"осмотреться": lookAround,
	"идти":        goRoom,
	"надеть":      take,
	"взять":       take,
	"применить":   apply,
}

func lookAround(p *Player, args []string) string {
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

func goRoom(p *Player, args []string) string {
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

func take(p *Player, args []string) string {
	if p.BackpackOn {
		if p.CurrentRoom.CheckItem(args[0]) {
			p.Inventory = append(p.Inventory, args[0])
			p.CurrentRoom.deleteItem(args[0])
			return fmt.Sprintf(("предмет добавлен в инвентарь: %s"), args[0])
		} else {
			return "нет такого"
		}
	} else if args[0] == "рюкзак" {
		p.BackpackOn = true
		p.CurrentRoom.deleteItem(args[0])
		return "вы надели: рюкзак"
	} else {
		return "некуда класть"
	}
}

func apply(p *Player, args []string) string {
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
