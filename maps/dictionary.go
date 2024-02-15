package maps

type Dictionary map[string]string

const (
	ErrDefinitionNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists         = DictionaryErr("provided word already has a definition")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (definition string, err error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDefinitionNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrDefinitionNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}
