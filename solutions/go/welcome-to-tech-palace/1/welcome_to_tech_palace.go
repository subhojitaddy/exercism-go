package techpalace

import (
    s "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + s.ToUpper(customer)
}

// AddStars creates custom star string (Eg. "********")
func AddStars(numStarsPerLine int) string {
    starString := ""
    for i := 0; i < numStarsPerLine; i++ {
        starString += "*"
    }
    return starString
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return AddStars(numStarsPerLine) + "\n" + welcomeMsg + "\n" + AddStars(numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    return s.TrimSpace(s.ReplaceAll(oldMsg, "*", " "))
}
