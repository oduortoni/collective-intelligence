package session

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const filePath = "sessions.csv"

func Remember(token string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}

	for _, record := range records {
		if record[0] == token {
			return record[1]
		}
	}

	return ""
}

func Save(username string) (string, bool) {
	token := generateToken()

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", false
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{token, username})
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return "", false
	}

	return token, true
}

func Delete(token string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}

	var newRecords [][]string
	for _, record := range records {
		if record[0] != token {
			newRecords = append(newRecords, record)
		}
	}

	if len(newRecords) == len(records) {
		return false
	}

	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(newRecords)
	if err != nil {
		fmt.Println("Error writing updated records to file:", err)
		return false
	}

	return true
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())

	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 16)

	for i := range token {
		token[i] = characters[rand.Intn(len(characters))]
	}

	return string(token)
}
