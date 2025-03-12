package jokes

import "math/rand"

var oneLiners = []string{
	"I told my wife she was drawing her eyebrows too high. She looked surprised.",
	"I’m reading a book on anti-gravity. It’s impossible to put down!",
	// Add more here!
}

func RandomOneLiner() string {
	return oneLiners[rand.Intn(len(oneLiners))]
}