package main

// hello returns a salutation
func hello(name, language string) string {
	if name == "" {
		if language != "spanish" {
			name = "world"
		} else {
			name = "mundo"
		}
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}
	return
}
