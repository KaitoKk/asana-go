package asanago

import "fmt"

type Error struct {
	Message string
	Phrase  string
}

type Errors struct {
	Errors []Error
}

func (errs Errors) PrintErrors() {
	for _, e := range errs.Errors {
		fmt.Println("error: ", e.Message)
		if e.Phrase != "" {
			fmt.Println(e.Phrase)
		}
	}
}
