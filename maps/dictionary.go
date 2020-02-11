package main

/*
var errNotFound = errors.New("could not find the word you were looking for")
var errWordExists = errors.New("word already exists")
*/

const (
	NoErr                = DictionaryErr("")
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, DictionaryErr) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, NoErr
}

func (d Dictionary) Add(word, definition string) DictionaryErr {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case NoErr:
		return ErrWordExists
	default:
		return err
	}
	return NoErr
}

func (d Dictionary) Update(word, definition string) DictionaryErr {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordDoesNotExists
	} else {
		d[word] = definition
		return NoErr
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
