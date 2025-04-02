package main

import (
	"fmt"
	"sort"
	"strings"
)

type Room struct {
	Name             string
	Description      string
	EntryDescription string
	NearbyRooms      []*Room
	Furniture        map[string][]string
	DoorOpen         bool
}

func (r *Room) String() string {
	var ItemStr string
	for key, items := range r.Furniture {
		ItemStr += key + ": ["
		for _, item := range items {
			ItemStr += item + ", "
		}
		ItemStr = ItemStr[:len(ItemStr)-2]
		ItemStr += "], "
	}

	var RoomStr string
	for _, value := range r.NearbyRooms {
		RoomStr += value.Name + ", "
	}

	if len(ItemStr) > 0 {
		ItemStr = ItemStr[:len(ItemStr)-2]
	}
	if len(RoomStr) > 0 {
		RoomStr = RoomStr[:len(RoomStr)-2]
	}

	return fmt.Sprintf("%s, Items: [%v], NearbyRooms: [%v]", r.Name, ItemStr, RoomStr)
}

// Функция, проверяющая, есть ли в комнате предметы
func (r *Room) IsEmpty() bool {
	for _, items := range r.Furniture {
		if len(items) != 0 {
			return false
		}
	}
	return true
}

// Функция, проверяющая, есть ли на том или ином элементе мебели предметы
func (r *Room) FurnitureIsEmpty(furniture string) bool {
	return len(r.Furniture[furniture]) == 0
}

func (r *Room) CheckItem(item string) bool {
	for _, fur := range r.Furniture {
		for _, itm := range fur {
			if itm == item {
				return true
			}
		}
	}
	return false
}

func (r *Room) getNearbyRoomsDescription() string {
	var result string
	result += "можно пройти - "

	if r.Name == "улица" {
		result += "домой"
		return result
	}

	for _, nearbyRoom := range r.NearbyRooms {
		result += nearbyRoom.Name + ", "
	}

	if len(result) > 0 {
		result = result[:len(result)-2]
	}
	return result
}

func (r *Room) deleteItem(item string) {
	for fur, slice := range r.Furniture {
		for i, v := range slice {
			if v == item {
				r.Furniture[fur] = append(slice[:i], slice[i+1:]...)
			}
		}
	}
}

func (r *Room) getItemsDescription() string {
	var fur []string
	var builder strings.Builder

	for key := range r.Furniture {
		if !r.FurnitureIsEmpty(key) {
			fur = append(fur, key)
		}
	}

	if len(fur) == 0 {
		return ""
	}

	sort.Strings(fur)
	for _, key := range fur {
		builder.WriteString(fmt.Sprintf("на %sе: ", key)) // на столе, на стуле
		items := r.getItemsInFurniture(key)

		for _, item := range items {
			builder.WriteString(item + ", ")
		}
	}
	return builder.String()
}

func (r *Room) getItemsInFurniture(furniture string) []string {
	var result []string

	result = append(result, r.Furniture[furniture]...)
	sort.Strings(result)
	return result
}
