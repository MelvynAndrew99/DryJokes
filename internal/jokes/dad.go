package jokes

import "math/rand"

var dadJokes = []Joke{
	{Setup: "Why don't skeletons fight each other?", Punchline: "They don't have the guts"},
	{Setup: "What do you call fake spaghetti?", Punchline: "An impasta."},
}

func RandomDadJoke() string {
	j := dadJokes[rand.Intn(len(dadJokes))]
	return j.Setup + " " + j.Punchline
}