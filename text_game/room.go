package main

import "fmt"

type Room struct {
	Name        string
	NearbyRooms []*Room
	Furniture   map[string]map[string]bool
}

func (r Room) String() string {
	var ItemStr string
	for key, value := range r.Furniture {
		ItemStr += key + ": ["
		for key, value := range value {
			if value {
				ItemStr += key + ", "
			}
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
	for _, value := range r.Furniture {
		for _, value := range value {
			if value {
				return false
			}
		}
	}
	return true
}

// Функция, проверяющая, есть ли на том или ином элементе мебели предметы
func (r *Room) FurnitureIsEmpty(furniture string) bool {
	for _, value := range r.Furniture[furniture] {
		if value {
			return false
		}
	}
	return true
}

func (r *Room) CheckItem(item string) bool {
	for _, value := range r.Furniture {
		if value[item] {
			return true
		}
	}
	return false
}
