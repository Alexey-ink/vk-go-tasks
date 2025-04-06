package main

import (
	"log"
	"sort"
	"strconv"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	var wg sync.WaitGroup
	channels := make([]chan interface{}, len(cmds)+1)

	for i := 0; i < len(channels); i++ {
		channels[i] = make(chan interface{}, 100)
	}

	for i, currentCmd := range cmds {
		wg.Add(1)
		go func(i int, с cmd) {
			defer wg.Done()
			defer close(channels[i+1])
			с(channels[i], channels[i+1])
		}(i, currentCmd)
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User

	var wg sync.WaitGroup
	uniqueUsers := make(map[uint64]struct{})
	var mu sync.Mutex

	for email := range in {
		wg.Add(1)
		go func(email string) {
			defer wg.Done()
			user := GetUser(email)

			mu.Lock()
			defer mu.Unlock()
			if _, exists := uniqueUsers[user.ID]; !exists {
				uniqueUsers[user.ID] = struct{}{}
				out <- user
			}
		}(email.(string))
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// 	in - User
	// 	out - MsgID

	var wg sync.WaitGroup
	batch := make([]User, 0, GetMessagesMaxUsersBatch)

	for user := range in {
		batch = append(batch, user.(User))

		if len(batch) == GetMessagesMaxUsersBatch {
			wg.Add(1)
			go func(users []User) {
				defer wg.Done()
				messages, err := GetMessages(users...)
				if err != nil {
					log.Printf("GetMessages failed: %v", err)
					return
				}
				for _, msg := range messages {
					out <- msg
				}
			}(batch)
			batch = make([]User, 0, GetMessagesMaxUsersBatch)
		}
	}

	// Обрабатываем оставшихся пользователей
	if len(batch) > 0 {
		wg.Add(1)
		go func(users []User) {
			defer wg.Done()
			messages, err := GetMessages(users...)
			if err != nil {
				log.Printf("GetMessages failed: %v", err)
				return
			}
			for _, msg := range messages {
				out <- msg
			}
		}(batch)
	}
	wg.Wait()
}

func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData

	var wg sync.WaitGroup
	sem := make(chan struct{}, HasSpamMaxAsyncRequests) // Семафор на 5 запросов

	for msg := range in {
		sem <- struct{}{}
		wg.Add(1)

		go func(msgID MsgID) {
			defer wg.Done()
			defer func() { <-sem }()

			hasSpam, err := HasSpam(msgID)
			if err != nil {
				log.Printf("HasSpam failed: %v", err)
				return
			}
			out <- MsgData{
				ID:      msgID,
				HasSpam: hasSpam,
			}
		}(msg.(MsgID))
	}
	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string

	results := make([]MsgData, 0, 100)

	for data := range in {
		results = append(results, data.(MsgData))
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].HasSpam == results[j].HasSpam {
			return results[i].ID < results[j].ID
		}
		return results[i].HasSpam && !results[j].HasSpam
	})

	for _, result := range results {
		out <- formatResult(result)
	}
}

func formatResult(data MsgData) string {
	return formatBool(data.HasSpam) + " " + formatUint64(uint64(data.ID))
}

func formatUint64(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func formatBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
