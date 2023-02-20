package nestedjson

import (
	"fmt"
	"golang.org/x/xerrors"
	"reflect"
)

var ErrTerminated = xerrors.New("normal shutdown of state machine")

type MinStruct struct {
	Label string
	ID string
	Category string
	Children []MinStruct
}

func SpecialFunc() {
	err := mutateUser(func(user interface{}) (err error) {
		var nextStep interface{}
		var ustate interface{}
		var processed uint64
		var terminated bool
		nextStep, processed, err = "", 1, nil
		ustate = user
		if xerrors.Is(err, ErrTerminated) {
			terminated = true
			return nil
		}
		fmt.Println(nextStep, ustate, processed, terminated)
		return err
	})
	fmt.Println(err)
}

func mutateUser(cb func(user interface{}) error) error {
	fmt.Println(reflect.TypeOf(cb))
	return nil
}