package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	teamName  string
	timestamp time.Time
	serverID  string
	result    string
}

type Team struct {
	name          string
	brokenServers map[string]bool
	penaltyTime   int
}

func parseTime(startTimeStr string) (time.Time, error) {
	return time.Parse("15:04:05", startTimeStr)
}

func calculatePenalty(startTime time.Time, requestTime time.Time) int {
	duration := requestTime.Sub(startTime)
	minutes := int(duration.Minutes())
	return minutes
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Чтение времени начала хакатона
	startTimeStr, _ := reader.ReadString('\n')
	startTimeStr = strings.TrimSpace(startTimeStr)
	startTime, _ := parseTime(startTimeStr)

	// Чтение количества запросов
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// Создание карты команд
	teams := make(map[string]*Team)

	// Чтение запросов
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Fields(line)
		teamName := strings.Trim(parts[0], `"`)
		requestTime, _ := parseTime(parts[1])
		serverID := parts[2]
		result := parts[3]

		// Если команда не существует, создаем её
		if _, exists := teams[teamName]; !exists {
			teams[teamName] = &Team{
				name:          teamName,
				brokenServers: make(map[string]bool),
				penaltyTime:   0,
			}
		}

		team := teams[teamName]

		// Обработка результатов
		if result == "ACCESSED" {
			if _, broken := team.brokenServers[serverID]; !broken {
				penalty := calculatePenalty(startTime, requestTime)
				team.penaltyTime += penalty
				team.brokenServers[serverID] = true
			}
		} else if result == "DENIED" || result == "FORBIDEN" {
			if _, broken := team.brokenServers[serverID]; !broken {
				team.penaltyTime += 20
			}
		}
	}

	// Подготовка списка команд для сортировки
	var result []struct {
		name        string
		brokenCount int
		penaltyTime int
	}

	for _, team := range teams {
		result = append(result, struct {
			name        string
			brokenCount int
			penaltyTime int
		}{
			name:        team.name,
			brokenCount: len(team.brokenServers),
			penaltyTime: team.penaltyTime,
		})
	}

	// Сортировка команд
	sort.Slice(result, func(i, j int) bool {
		if result[i].brokenCount != result[j].brokenCount {
			return result[i].brokenCount > result[j].brokenCount
		}
		if result[i].penaltyTime != result[j].penaltyTime {
			return result[i].penaltyTime < result[j].penaltyTime
		}
		return result[i].name < result[j].name
	})

	// Вывод результатов
	for i, team := range result {
		fmt.Printf("%d \"%s\" %d %d\n", i+1, team.name, team.brokenCount, team.penaltyTime)
	}
}
