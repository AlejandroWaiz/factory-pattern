package main

type Pirate struct {
	name  string
	power string
}

func (p *Pirate) setName(name string) {
	p.name = name
}

func (p *Pirate) setPower(power string) {
	p.power = power
}

func (p *Pirate) getName() string {
	return p.name
}

func (p *Pirate) getPower() string {
	return p.power
}
