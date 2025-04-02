package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var player *Player

const UnknownCommandMessage = "неизвестная команда"

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
		RoomTasks: map[string]*Room{
			"собрать рюкзак": rooms["кухня"],
			"идти в универ":  rooms["кухня"],
		},
	}
}

func handleCommand(command string) string {
	commandParts := strings.Split(command, " ")
	var args []string

	if len(commandParts) == 0 || len(commandParts) > 3 {
		return UnknownCommandMessage
	}

	if commandParts[0] == "осмотреться" && len(commandParts) > 1 {
		return UnknownCommandMessage
	}

	if commandParts[0] != "применить" && len(commandParts) > 2 {
		return UnknownCommandMessage
	}

	if len(commandParts) > 1 {
		args = commandParts[1:]
	} else {
		args = []string{}
	}

	switch commandParts[0] {
	case "осмотреться":
		return player.lookAround()
	case "идти":
		return player.goRoom(args)
	case "надеть", "взять":
		return player.take(args)
	case "применить":
		return player.apply(args)
	default:
		return UnknownCommandMessage
	}
}

func initRooms() map[string]*Room {
	kitchen := &Room{
		Name:             "кухня",
		Description:      "ты находишься на кухне, ",
		EntryDescription: "кухня, ничего интересного. ",
		Furniture: map[string][]string{
			"стол": {
				"чай",
			},
		},
		DoorOpen: true,
	}

	corridor := &Room{
		Name:             "коридор",
		EntryDescription: "ничего интересного. ",
		Furniture:        map[string][]string{},
		DoorOpen:         false,
	}

	livingRoom := &Room{
		Name:             "комната",
		EntryDescription: "ты в своей комнате. ",
		Furniture: map[string][]string{
			"стол": {
				"ключи",
				"конспекты",
			},
			"стул": {
				"рюкзак",
			},
		},
		DoorOpen: true,
	}

	street := &Room{
		Name:             "улица",
		EntryDescription: "на улице весна. ",
		Furniture:        map[string][]string{},
		DoorOpen:         true,
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
