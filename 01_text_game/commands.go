package main

import (
	"fmt"
	"sort"
)

var COMMANDS = map[string]func(*Player, []string) string{
	"осмотреться": lookAround,
	"идти":        goRoom,
	"надеть":      take,
	"взять":       take,
	"применить":   apply,
}

func lookAround(p *Player, args []string) string {
	var result string

	if p.CurrentRoom.IsEmpty() {
		result = "пустая комната. "
	} else {

		if p.CurrentRoom.Name == "кухня" {
			result = "ты находишься на кухне, "
		}

		keys := make([]string, 0, len(p.CurrentRoom.Furniture))
		for key := range p.CurrentRoom.Furniture {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, fur := range keys {
			if !p.CurrentRoom.FurnitureIsEmpty(fur) {
				result += "на " + fur + "е: " // на столе, на стуле

				items := make([]string, 0, len(p.CurrentRoom.Furniture[fur]))
				for item, exists := range p.CurrentRoom.Furniture[fur] {
					if exists {
						items = append(items, item)
					}
				}

				sort.Strings(items)
				for _, item := range items {
					result += item + ", "
				}
			}
		}

		if p.CurrentRoom.Name == "кухня" {
			result += "надо "
			if !p.BackpackOn {
				result += "собрать рюкзак и "
			}
			result += "идти в универ. "
		} else {
			result = result[:len(result)-2]
			result += ". "
		}
	}

	result += p.CurrentRoom.getNearbyRoomsDescription()
	return result
}

func goRoom(p *Player, args []string) string {
	var result string

	for _, room := range p.CurrentRoom.NearbyRooms {
		if args[0] == room.Name {
			if room.Name == "улица" && !p.CurrentRoom.DoorOpen {
				return "дверь закрыта"
			}

			p.CurrentRoom = room

			if room.Name == "улица" && p.CurrentRoom.DoorOpen {
				result = "на улице весна. "
			} else if room.Name == "комната" {
				result = "ты в своей комнате. "
			} else if room.Name == "кухня" {
				result = "кухня, ничего интересного. "
			} else {
				result = "ничего интересного. "
			}

			result += p.CurrentRoom.getNearbyRoomsDescription()
			return result

		}
	}

	result += "нет пути в " + args[0]
	return result
}

func take(p *Player, args []string) string {
	if p.BackpackOn {
		if p.CurrentRoom.CheckItem(args[0]) {
			*p.Inventory = append(*p.Inventory, args[0])
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
	if p.CheckItem(args[0]) {
		if args[0] == "ключи" && p.CurrentRoom.Name == "коридор" {
			if args[1] == "дверь" {
				p.CurrentRoom.DoorOpen = true
				return "дверь открыта"
			} else {
				return "не к чему применить"
			}
		} else {
			fmt.Println(args[0])
			fmt.Println(args[1])
			return "не к чему применить"
		}
	} else {
		return "нет предмета в инвентаре - " + args[0]
	}
}
