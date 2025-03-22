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

	var command1 = "осмотреться вокруг в "
	answer := handleCommand(command1)
	fmt.Println(answer)

	command2 := "идти коридор"
	answer = handleCommand(command2)
	fmt.Println(answer)

	command3 := "идти комната"
	answer = handleCommand(command3)
	fmt.Println(answer)

	command4 := "взять ключи"
	answer = handleCommand(command4)
	fmt.Println(answer)

	command5 := "надеть рюкзак"
	answer = handleCommand(command5)
	fmt.Println(answer)

	command6 := "взять ключи"
	answer = handleCommand(command6)
	fmt.Println(answer)

	command7 := "взять ключи"
	answer = handleCommand(command7)
	fmt.Println(answer)

	fmt.Println(player)

}

func initGame() {
	/*
		эта функция инициализирует игровой мир и принимае
	*/

	rooms := initRooms()

	player = &Player{
		Name:        "Alexey",
		CurrentRoom: rooms["кухня"],
		Inventory:   &[]string{"фонарик", "очки"},
	}

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

	var result string
	fmt.Println(commandParts)

	if cmdFunc, exists := COMMANDS[commandParts[0]]; exists {
		result := cmdFunc(player, args)
		fmt.Println(result)
	} else {
		fmt.Println("Команда не распознана")
	}

	return result
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
