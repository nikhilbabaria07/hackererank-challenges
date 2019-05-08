package main

import (
	"fmt"
	"strconv"
)

type userDB struct {
	name string
	occ  int
}

var nameOcc []userDB
var tempnO userDB

func countOcc(name string) int {
	var cntr int
	for _, v := range nameOcc {
		if v.name == name {
			cntr++
		}
	}
	return cntr
}

func usernamesSystem(u []string) []string {
	for _, value := range u {
		tempnO.name = value
		tempnO.occ = countOcc(value)
		nameOcc = append(nameOcc, tempnO)
	}

	var res []string

	for _, v := range nameOcc {
		if v.occ != 0 {
			res = append(res, v.name+strconv.Itoa(v.occ))
			fmt.Println(res)

		} else {
			res = append(res, v.name)
		}

	}
	return res
}

func main() {

	//var unique []string
	users := []string{"Penn", "Teller", "Penn", "Teller", "Penn", "Teller"}
	res := usernamesSystem(users)
	fmt.Println(res)

}
