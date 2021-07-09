package main

type IPirate interface {
	setName(name string)
	setPower(power string)
	getName() string
	getPower() string
}
