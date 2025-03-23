package main

import (
	"fmt"
	"strings"
)

var player *Player

func main() {

	initGame()

	commands := []string{
		"осмотреться",
		"завтракать",
		"идти комната",
		"идти коридор",
		"применить ключи дверь",
		"идти комната",
		"осмотреться",
		"взять ключи",
		"надеть рюкзак",
		"осмотреться",
	}

	for _, command := range commands {
		answer := handleCommand(command)
		fmt.Printf("[%s]\n", command)
		fmt.Println(answer)
		fmt.Println()
	}

	fmt.Println(player)

}

func initGame() {

	rooms := initRooms()

	player = &Player{
		Name:        "Alexey",
		CurrentRoom: rooms["кухня"],
		Inventory:   &[]string{},
	}
}

func handleCommand(command string) string {

	commandParts := strings.Split(command, " ")
	var args []string

	if len(commandParts) > 1 {
		args = commandParts[1:]
	} else {
		args = []string{}
	}

	if cmdFunc, exists := COMMANDS[commandParts[0]]; exists {
		result := cmdFunc(player, args)
		return result
	} else {
		return "неизвестная команда"
	}
}

func initRooms() map[string]*Room {

	kitchen := &Room{
		Name: "кухня",
		Furniture: map[string]map[string]bool{
			"стол": {
				"чай": true,
			},
		},
		DoorOpen: true,
	}

	corridor := &Room{
		Name:      "коридор",
		Furniture: map[string]map[string]bool{},
		DoorOpen:  false,
	}

	livingRoom := &Room{
		Name: "комната",
		Furniture: map[string]map[string]bool{
			"стол": {
				"ключи":     true,
				"конспекты": true,
			},
			"стул": {
				"рюкзак": true,
			},
		},
		DoorOpen: true,
	}

	street := &Room{
		Name:      "улица",
		Furniture: map[string]map[string]bool{},
		DoorOpen:  true,
	}

	kitchen.NearbyRooms = []*Room{corridor}
	corridor.NearbyRooms = []*Room{kitchen, livingRoom, street}
	livingRoom.NearbyRooms = []*Room{corridor}
	street.NearbyRooms = []*Room{corridor}

	return map[string]*Room{
		"кухня":   kitchen,
		"коридор": corridor,
		"комната": livingRoom,
		"улица":   street,
	}
}
