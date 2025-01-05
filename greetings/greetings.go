package greetings
import 
(
    "errors"
    "fmt"
    "math/rand"
    

)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	
	message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}


// Hellos returns a map of greeting messages for a list of names. 
func Hellos(names [] string) (map[string]string,error){
	// A map to associate names with messages. 
	messages := make(map[string]string)

	// Loop through the list of names. 
	for _, name := range names{
		message, err := Hello(name)
		if err != nil {
			return nil , err
		}
		messages[name]=message
	}
	return messages,nil
}


func randomFormat() string {
	formats := [] string{
		"Hi, %v. Welcome!",
        "Hello, %v! It's great to meet you.",
        "Greetings, %v! Have a wonderful day.",
	}

	return formats[rand.Intn(len(formats))]
}

	