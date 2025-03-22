package main

import "fmt"

var COMMANDS = map[string]func(*Player, []string) string{
	"осмотреться": lookAround,
	"идти":        goRoom,
	"надеть":      take,
	"взять":       take,
	// "применить":   apply,
}

func lookAround(p *Player, args []string) string {
	var result string

	if p.CurrentRoom.IsEmpty() {
		result = "пустая комната. "
	} else {

		if p.CurrentRoom.Name == "кухня" {
			result = "ты находишься на кухне, "
		}

		for key, value := range p.CurrentRoom.Furniture {
			if !p.CurrentRoom.FurnitureIsEmpty(key) {
				result += "на " + key + "e: " // на столе, на стуле
				for key, value := range value {
					if value {
						result += key + ", "
					}
				}
			}
		}

		// Нужно добавить проверку, положил ли конспекты в рюкзак
		if p.CurrentRoom.Name == "кухня" {
			result += "надо "
			if !p.BackpackOn {
				result += "собрать рюкзак и"
			}
			result += " идти в универ. "
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
			p.CurrentRoom = room

			if room.Name == "улица" {
				result = "на улице весна. "
			} else if room.Name == "комната" {
				result = "ты в своей комнате. "
			} else if room.Name == "кухня" {
				result = "кухня, ничего интересного. "
			} else {
				result = "ничего интересного. "
			}

			result += p.CurrentRoom.getNearbyRoomsDescription()

		} else {
			result = "нет пути в " + args[0]
		}

	}

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
		return "вы надели: рюкзак"
	} else {
		return "некуда класть"
	}
}
