package main

import (
	"fmt"
	"sort"
	"strings"
)

type Player struct {
	Name        string
	CurrentRoom *Room
	Inventory   []string
	BackpackOn  bool
	RoomTasks   map[string]*Room
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
		builder.WriteString(p.CurrentRoom.Description)
		builder.WriteString(p.CurrentRoom.getItemsDescription())

		if !p.IsEmptyTasks() && p.AvailableTasks(p.CurrentRoom) {
			builder.WriteString(p.getTasks())
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
	var result strings.Builder
	for _, room := range p.CurrentRoom.NearbyRooms {
		if !(args[0] == room.Name) {
			continue
		}
		if room.Name == "улица" && !p.CurrentRoom.DoorOpen {
			return "дверь закрыта"
		}
		p.CurrentRoom = room
		result.WriteString(room.EntryDescription)
		result.WriteString(p.CurrentRoom.getNearbyRoomsDescription())
		return result.String()
	}
	result.WriteString("нет пути в " + args[0])
	return result.String()
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
		delete(p.RoomTasks, "собрать рюкзак")
		return "вы надели: рюкзак"
	}
	return "некуда класть"
}

func (p *Player) apply(args []string) string {
	const NotApplicable = "не к чему применить"

	if !p.CheckItem(args[0]) {
		return "нет предмета в инвентаре - " + args[0]
	}

	if args[0] != "ключи" || p.CurrentRoom.Name != "коридор" {
		return NotApplicable
	}

	if args[1] != "дверь" {
		return NotApplicable
	}
	p.CurrentRoom.DoorOpen = true
	return "дверь открыта"
}

func (p *Player) getTasks() string {
	var builder strings.Builder
	ActiveTasksCount := len(p.RoomTasks)

	builder.WriteString("надо ")

	keys := make([]string, 0, len(p.RoomTasks))
	for key := range p.RoomTasks {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	i := 1
	for _, task := range keys {
		if i == ActiveTasksCount && ActiveTasksCount != 1 {
			builder.WriteString(" и ")
		}
		builder.WriteString(task)
		i++
	}
	builder.WriteString(". ")
	return builder.String()
}

func (p *Player) IsEmptyTasks() bool {
	return len(p.RoomTasks) == 0
}

func (p *Player) AvailableTasks(r *Room) bool {
	for _, value := range p.RoomTasks {
		if value.Name == r.Name {
			return true
		}
	}
	return false
}
