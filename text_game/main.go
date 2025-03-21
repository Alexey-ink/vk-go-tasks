package main

import (
	"fmt"
	"strings"
)

var player *Player

func main() {
	/*
		в этой функции можно ничего не писать,
		но тогда у вас не будет работать через go run main.go
		очень круто будет сделать построчный ввод команд тут, хотя это и не требуется по заданию
	*/

	initGame()

	fmt.Println(player)

	var command1 = handleCommand("осмотреться вокруг в лесу")
	args := []string{}

	if cmdFunc, exists := Commands[command1]; exists {
		result := cmdFunc(player, args)
		fmt.Println(result)
	} else {
		fmt.Println("Команда не распознана")
	}

}

func initGame() {
	/*
		эта функция инициализирует игровой мир и принимае
	*/

	rooms := initRooms()

	player = &Player{
		Name:        "Alexey",
		CurrentRoom: rooms["комната"],
		Inventory:   &[]string{"фонарик", "очки"},
	}

	fmt.Println(player)

}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/

	commandParts := strings.Split(command, " ")

	fmt.Println(commandParts)

	return commandParts[0]
}

func initRooms() map[string]*Room {

	kitchen := &Room{
		Name: "кухня",
		Furniture: map[string]map[string]bool{
			"стол": {
				"чай": true,
			},
		},
	}

	corridor := &Room{
		Name:      "коридор",
		Furniture: map[string]map[string]bool{},
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
	}

	kitchen.NearbyRooms = []*Room{corridor}
	corridor.NearbyRooms = []*Room{kitchen, livingRoom}
	livingRoom.NearbyRooms = []*Room{corridor}

	return map[string]*Room{
		"кухня":   kitchen,
		"коридор": corridor,
		"комната": livingRoom,
	}
}
