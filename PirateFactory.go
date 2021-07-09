package main

import "errors"

func GetPirate(pirate string) (IPirate, error) {

	if pirate == "luffy" {

		return NewMonkeyDLuffy(), nil

	} else if pirate == "law" {

		return NewTrafalgarDWaterLaw(), nil

	} else if pirate == "roger" {

		return NewGolDRoger(), nil

	}

	return nil, errors.New("Not a valid pirate.")

}
