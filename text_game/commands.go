package main

var Commands = map[string]func(*Player, []string) string{
	"осмотреться": lookAround,
	// "идти":        goRoom,
	// "надеть":      take,
	// "взять":       take,
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

	result += "можно пройти - "

	for _, room := range p.CurrentRoom.NearbyRooms {
		result += room.Name + ", "
	}

	result = result[:len(result)-2]
	return result
}
