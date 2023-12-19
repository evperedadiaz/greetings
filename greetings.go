package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("nombre vacío")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"¡Hola, %s! ¡Bienvenido!",
		"¡Me alegro de verte, %s! ¡Bienvenido!",
		"¿Otra vez por aquí?, %s ¡Bienvenido!",
	}

	return formats[rand.Intn(len(formats))]
}
