package mydict

import "errors"

type Dictinary map[string]string

var (
	errNotFound = errors.New("Not Found")
    errWordExists = errors.New("That word already exist")
	errCantUpdate = errors.New("Cant update ono-existiong word")
)

func (d Dictinary) Serach(word string) (string,error){
	value, exists := d[word]
	if exists{
		return value,nil
	}
	return "",errNotFound
}

// Add a word to tge dictionary
func (d Dictinary) Add(word,def string) error {
	_, err := d.Serach(word)
	switch err{
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictinary) Update(word, definition string) error{
	_, err := d.Serach(word)
	switch err{
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word 
func (d Dictinary) Delete(word string){
	delete(d,word)
}