package main

/*
Create a program which has three structs:
Gopher
Weapon
Consumable

Gopher should contain the following statistics about a gopher:
hitpoints - an integer
weapon - the Weapon equipped
strength - an integer
agility - an integer
intellect - an integer
coins - an integer

Weapon should contain the following data:
damage - a slice of two integers, the interval of damage the weapon can deal
price - the price of the weapon
strengthReq - an integer, strength requirements to yield the weapon
agilityReq - an integer, strength requirements to yield the weapon
intelligenceReq - an integer, intellect requirements to yield the weapon

Consumable should contain the following data:
hitpointsEffect - an integer, the effect on hitpoints
*/
type Gopher struct {
	name      string
	hitpoints int
	weapon    string
	strength  int
	agility   int
	intellect int
	coins     int
}
type Weapon struct {
	damage          []int
	price           int
	strengthReq     int
	agilityReq      int
	intelligenceReq int
}
type Consumable struct {
	hitpointsEffect int
	price           int
}
