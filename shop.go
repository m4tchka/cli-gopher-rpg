package main

/*
The shop has the following items for sale with unlimited supply of them:
Consumables:
small_health_potion - consumable - 5 gold
hitpointsEffect - 5
medium_health_potion - consumable - 9 gold
hitpointsEffect - 10
big_health_potion - consumable - 18 gold
hitpointsEffect - 20
*/

var Consumables = map[string]Consumable{
	"small_health_potion":  {hitpointsEffect: 5, price: 5},
	"medium_health_potion": {hitpointsEffect: 10, price: 9},
	"large_health_potion":  {hitpointsEffect: 20, price: 18},
}

/*
Weapons:
knife - weapon - 10 gold
damage - [2-3]
all requirements are 0
sword - weapon - 35 gold
damage - [3-5]
strengthReq - 2
ninjaku - weapon - 25 gold
damage - [1-7]
agilityReq - 2
wand - weapon - 30 gold
damage - [3-3]
intellectReq - 2
gophermourne - weapon - 65 gold
damage - [6-7]
strengthReq - 5
warglaives_of_gopherinoth - weapon - 55 gold
damage - [2-9]
agilityReq - 5
codeseeker - weapon - 60 gold
damage - [4-4]
intellectReq - 5
*/
var Weapons = map[string]Weapon{
	"bare-hands":                {damage: []int{1, 1}},
	"knife":                     {damage: []int{2, 3}, price: 10},
	"sword":                     {damage: []int{3, 5}, price: 35, strengthReq: 2},
	"ninjaku":                   {damage: []int{1, 7}, price: 25, agilityReq: 2},
	"wand":                      {damage: []int{3, 3}, price: 30, intelligenceReq: 2},
	"gophermourne":              {damage: []int{6, 7}, price: 65, strengthReq: 5},
	"warglaives_of_gopherinoth": {damage: []int{2, 9}, price: 55, agilityReq: 5},
	"codeseeker":                {damage: []int{4, 4}, price: 60, intelligenceReq: 5},
}
