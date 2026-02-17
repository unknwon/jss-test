package main

import (
	"math/rand"
)

// RandomInt returns a random integer between min and max (inclusive).
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

var quotes = []string{
	"The only way to do great work is to love what you do. – Steve Jobs",
	"Be the change that you wish to see in the world. – Mahatma Gandhi",
	"In the middle of difficulty lies opportunity. – Albert Einstein",
	"Stay hungry, stay foolish. – Steve Jobs",
	"Simplicity is the ultimate sophistication. – Leonardo da Vinci",
	"Talk is cheap. Show me the code. – Linus Torvalds",
	"First, solve the problem. Then, write the code. – John Johnson",
	"Code is like humor. When you have to explain it, it's bad. – Cory House",
	"Any fool can write code that a computer can understand. Good programmers write code that humans can understand. – Martin Fowler",
	"The best error message is the one that never shows up. – Thomas Fuchs",
}

// RandomQuote returns a random quote from the predefined list.
func RandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}
