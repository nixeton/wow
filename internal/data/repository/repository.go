package repository

import (
	"math/rand"
)

var quotes = []string{
	"The only thing we have to fear is fear itself.",
	"Injustice anywhere is a threat to justice everywhere.",
	"Imagination is more important than knowledge.",
	"I have a dream that one day this nation will rise up and live out the true meaning of its creed.",
	"To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.",
	"The best way to find yourself is to lose yourself in the service of others.",
	"What you do makes a difference, and you have to decide what kind of difference you want to make.",
	"Success is not final, failure is not fatal: It is the courage to continue that counts.",
	"Life is what happens when you're busy making other plans.",
	"The future belongs to those who believe in the beauty of their dreams.",
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetRandomQuote() string {

	return getRandomQuote(quotes)
}

func getRandomQuote(quotes []string) string {
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}
