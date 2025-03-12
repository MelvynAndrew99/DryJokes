package jokes

import "math/rand"

var knockJokes = []Joke{
	{Setup: "Knock knock. Who's there? Cow says.", Punchline: "No, a cow says moooo!"},
	{Setup: "Knock knock. Who's there? Lettuce.", Punchline: "Lettuce in, it's cold out here!"},
	// Add more here!
}

func RandomKnockJoke() string {
	j := knockJokes[rand.Intn(len(knockJokes))]
	return j.Setup + " " + j.Punchline
}