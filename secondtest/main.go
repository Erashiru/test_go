package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.Create("hypeauditor_data.csv")
	if err != nil {
		log.Fatalf("Error creating CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Рейтинг", "Имя", "Ник", "Категории", "Подписчики", "Аудитория"}
	writer.Write(headers)

	url := "https://hypeauditor.com/top-instagram-all-russia/"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalf("Error loading page %v", err)
	}

	doc.Find(".row").Each(func(i int, s *goquery.Selection) {
		rating := s.Find(".rank").Text()
		name := s.Find(".contributor__title").Text()
		username := s.Find(".contributor__name-content").Text()
		category := s.Find(".category").Text()
		followers := s.Find(".subscribers").Text()
		country := s.Find(".audience").Text()
		// audience := s.Find(".audience").Text()

		rating = strings.TrimSpace(rating)
		name = strings.TrimSpace(name)
		username = strings.TrimSpace(username)
		category = strings.TrimSpace(category)
		followers = strings.TrimSpace(followers)
		country = strings.TrimSpace(country)
		// audience = strings.TrimSpace(audience)

		csvData := []string{rating, name, username, category, followers, country}
		if rating != "" {
			writer.Write(csvData)
		}
	})

	log.Println("Data has been writen to hypeauditor_data.csv successfully!")
}
