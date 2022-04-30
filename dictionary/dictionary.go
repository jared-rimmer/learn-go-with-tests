package dictionary

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find word")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
