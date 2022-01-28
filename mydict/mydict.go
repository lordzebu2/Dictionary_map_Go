package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string
// Dictionary는 map에대한 별명이다. 다른곳에서 map쓸수있게, map타입의 별명.
// type은 method를 가질 수 있다.

var (
	errNotFound = errors.New("not found")
	errCantUpdate = errors.New("can't update non-existing word")
 	errWordExists = errors.New("that word akready exists")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error){
	value, exists := d[word] //d[word]는 맵에 단어(key)를 호출함 맵은 그단어값(value)이나 ok를 준다.
								//맵에 key의 존재여부를 알려주게 하는 방법. 
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) string {
	delete(d, word)
	check := "Delete Complete"
	return check
}