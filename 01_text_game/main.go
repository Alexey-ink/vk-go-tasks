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

	//fmt.Println(MainPlayer)

	// commands1 := []string{
	// 	"осмотреться",
	// 	"идти коридор",
	// 	"идти комната",
	// 	"осмотреться",
	// 	"надеть рюкзак",
	// 	"взять ключи",
	// 	"взять конспекты",
	// 	"идти коридор",
	// 	"применить ключи дверь",
	// 	"идти улица",
	// }

	commands2 := []string{
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
		"взять ключи",
		"взять телефон",
		"взять ключи",
		"осмотреться",
		"взять конспекты",
		"осмотреться",
		"идти коридор",
		"идти кухня",
		"осмотреться",
		"идти коридор",
		"идти улица",
		"применить ключи дверь",
		"применить телефон шкаф",
		"применить ключи шкаф",
		"идти улица",
	}

	for _, command := range commands2 {
		answer := handleCommand(command)
		fmt.Println(answer)
	}

	//fmt.Println(MainPlayer)

}

func initGame() {
	/*
		эта функция инициализирует игровой мир и принимае
	*/

	rooms := initRooms()

	player = &Player{
		Name:        "Alexey",
		CurrentRoom: rooms["кухня"],
		Inventory:   &[]string{},
	}

	//fmt.Println(player)

}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/

	commandParts := strings.Split(command, " ")
	var args []string

	if len(commandParts) > 1 {
		args = commandParts[1:]
	} else {
		args = []string{}
	}

	//fmt.Println(commandParts)

	if cmdFunc, exists := COMMANDS[commandParts[0]]; exists {
		result := cmdFunc(player, args)
		//fmt.Println(result)
		return result
	} else {
		//fmt.Println("неизвестная команда")
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
