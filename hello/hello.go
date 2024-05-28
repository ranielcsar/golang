package hello

const (
	pt      = "pt-br"
	spanish = "spanish"

	engPrefix     = "Hello, "
	ptPrefix      = "Olá, "
	spanishPrefix = "Hola, "
)

func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return getPrefix(lang) + name + "!"
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case pt:
		prefix = ptPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = engPrefix
	}

	return
}
