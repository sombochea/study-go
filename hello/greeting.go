package hello

import "errors"

func Greeting() string {
	return "Hello, world!"
}

func SayHelloTo(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	return "Hello, " + name + "!", nil
}