package dictionary

const (
	ErrNotFound          = DictionaryErr("couldn't find the word")
	ErrWordExists        = DictionaryErr("cannot add existing word")
	ErrWordDoesNotExists = DictionaryErr("cannot update unknown word")
)

type Dictionary map[string]string
type DictionaryErr string

// func Search(dictionary Dictionary, word string) string {
// 	return dictionary[word]
// }

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// d[word] = definition

	// return nil

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
