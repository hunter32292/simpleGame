package game

// Event - A type of event that can occur during the game
type Event struct {
	Name     string
	Value    int
	Damage   int
	Question string
	Options  []string
	Answer   string
	outcomes map[string]string
}

// Outcome - Print the map result of the event
func (e Event) Outcome(answer string) string {
	return e.outcomes[answer]
}

// GenerateEvent - Create a random event that can occur
func GenerateEvent() Event {
	return Event{
		Name:     "Found a gold peace",
		Value:    1,
		Damage:   0,
		Question: "Do you take it?",
		Options:  []string{"yes", "no"},
		Answer:   "yes",
		outcomes: map[string]string{
			"yes": "You put the gold in your pocket",
			"no":  "An old man picks up the gold",
		},
	}
}

// ContainsOption - Check if the option given is one of the provided options
func ContainsOption(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
