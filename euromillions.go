package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Returns an int >= min, < max
func randomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// Check if generated int is unique
// Append generated int to slice

func getNumbers(count int, min int, max int) []int {
	// Create slice (array)
	var euromillions []int

	// Loop to fill slice
	for len(euromillions) < count {

		// Get new random number
		newNumber := randomInt(min, max)
		euromillions = addNumber(euromillions, newNumber)
		euromillions = uniqueNumbers(euromillions)
	}
	return euromillions
}

func uniqueNumbers(euromillions []int) []int {
	// Loop map to check if number is unique
	keys := make(map[int]bool)
	euromillionsUnique := []int{}
	for _, entry := range euromillions {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			euromillionsUnique = append(euromillionsUnique, entry)
		}
	}
	return euromillionsUnique
}

// Add new number to slice
func addNumber(euromillions []int, addNumber int) []int {
	euromillions = append(euromillions, addNumber)
	// Sort slice
	sort.Ints(euromillions)
	return euromillions
}

func convertSlicetoString(slice []int) string {
	sliceConvert := []string{}

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range slice {
		number := slice[i]
		text := strconv.Itoa(number)
		sliceConvert = append(sliceConvert, text)
	}
	sliceConverted := strings.Join(sliceConvert, " ")
	return sliceConverted
}

func sendTelegram(message string) {
	var token string = "**YOUR TOKEN ID**"
	var chatID string = "**YOUR CHAT_ID**"
	//var url string

	loginURL := "https://api.telegram.org/bot" + token + "/sendMessage"

	urlData := url.Values{}
	urlData.Set("chat_id", chatID)
	urlData.Set("text", message)

	http.PostForm(loginURL, urlData)

}

func main() {
	var luckyNumbers string
	luckyNumbers = "Euromillions: " + convertSlicetoString(getNumbers(5, 1, 50)) + "\nStars: " + convertSlicetoString(getNumbers(2, 1, 12))
	sendTelegram(luckyNumbers)
}
