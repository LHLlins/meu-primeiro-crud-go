package controller

import "github.com/LHLlins/meu-primeiro-crud-go/tree/main/src/controller/model"

// resetGlobalState resets the users slice and nextID for testing.
func resetGlobalState() {
	users = []model.User{}
	nextID = 1
}
