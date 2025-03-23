package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var player *Player

func main() {

	initGame()

	fmt.Printf("Welcome to the game!\n\n")
	fmt.Println(player)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду:")
		if !scanner.Scan() {
			break
		}
		command := scanner.Text()
		answer := handleCommand(command)
		fmt.Printf("[%s]\n\n", answer)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
	}
}

func initGame() {
	rooms := initRooms()

	player = &Player{
		Name:        "Alexey",
		CurrentRoom: rooms["кухня"],
		Inventory:   []string{},
	}
}

func handleCommand(command string) string {

	commandParts := strings.Split(command, " ")
	var args []string

	if len(commandParts) == 0 || len(commandParts) > 3 {
		return "неизвестная команда"
	}

	if commandParts[0] == "осмотреться" && len(commandParts) > 1 {
		return "неизвестная команда"
	}

	if commandParts[0] != "применить" && len(commandParts) > 2 {
		return "неизвестная команда"
	}

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
		Name:        "кухня",
		Description: "кухня, ничего интересного. ",
		Furniture: map[string]map[string]bool{
			"стол": {
				"чай": true,
			},
		},
		DoorOpen: true,
	}

	corridor := &Room{
		Name:        "коридор",
		Description: "ничего интересного. ",
		Furniture:   map[string]map[string]bool{},
		DoorOpen:    false,
	}

	livingRoom := &Room{
		Name:        "комната",
		Description: "ты в своей комнате. ",
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
		Name:        "улица",
		Description: "на улице весна. ",
		Furniture:   map[string]map[string]bool{},
		DoorOpen:    true,
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
