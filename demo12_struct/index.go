package main

import "fmt"

type profile struct {
	name    string
	address string
	phone   string
	salary  int
}

func main() {
	var profile profile
	profile.name = "Chaiwat"
	profile.address = "Khonkaen"
	profile.phone = "088-8888888"
	profile.salary = 25000

	show(profile)
	update(&profile)
	show(profile)

	profile = profile.clear()
	profile = profile.addBonus(5000)
	show(profile)
}

func show(p profile) {
	fmt.Println(p)
}

func update(p *profile) {
	p.salary = p.salary + 100
}

func (p profile) clear() profile {
	p.salary = 0
	return p
}

func (p profile) addBonus(b int) profile {
	p.salary += b
	return p
}
