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
	message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
	formats := [] string{
		"Hi, %v. Welcome!",
        "Hello, %v! It's great to meet you.",
        "Greetings, %v! Have a wonderful day.",
	}

	return formats[rand.Intn(len(formats))]
}

	