package greetings
import 
(
    "errors"
    "fmt"
    "math/rand"
    "time"

)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	rand.Seed(time.Now().UnixNano())

	greetings := [] string{
		"Hi, %v. Welcome!",
        "Hello, %v! It's great to meet you.",
        "Greetings, %v! Have a wonderful day.",
	}

	randomGreeting := greetings[rand.Intn(len(greetings))]

	message := fmt.Sprintf(randomGreeting, name)
	return message, nil
}
