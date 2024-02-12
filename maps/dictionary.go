package maps

import "errors"

type Dictionary map[string]string

var ErrDefinitionNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDefinitionNotFound
	}

	return definition, nil
}
