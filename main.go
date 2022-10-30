package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
	The game is a turn-based one. There are two gophers and they can each decide what to do on their turn. Each gopher starts with 30 hitpoints, 20 gold and all their attributes are 0.

The possible actions are:
attack - attack the other gopher with the weapon you have equipped at the moment. If none is equipped, then you are attacking bare-handed for a damage of 1 hitpoint.
Choose the actual damage dealt at random based on the weapon’s damage interval
buy weapon <item> - spend the coins necessary to buy the given weapon based on its cost. You equip it over your current weapon. You can’t buy weapons you are still illegible to use due to insufficient stats.
buy consumable <item> - spend the coins necessary to buy the given consumable based on its cost and use it directly.
work - spend the turn working for the local warlord and gain anywhere between 5 and 15 coins (picked at random)
train <skill> - train a given attribute (strength, agility or intellect) and increase it by 2. Training costs 5 gold.
exit - exits the game

Extra challenges:
Implement a game log which prints all events which occur throughout the game (see examples below)
Implement consumables which give you an attribute boost (strength, agility, intellect) for a limited duration (e.g. 3 turns). This allows you to buy and use items without having the proper training yet but only for a limited duration
*/
var Gopher1 = &Gopher{
	name:      "gopher1",
	hitpoints: 30,
	weapon:    "bare-handed",
	coins:     20,
}
var Gopher2 = &Gopher{
	name:      "gopher2",
	hitpoints: 30,
	weapon:    "bare-handed",
	coins:     20,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Welcome to a game of Gopher RPG")
	r := bufio.NewReader(os.Stdin)
	handleAction(r)
}

func handleAction(r *bufio.Reader) {
	action := ""
	turn := 0
	for action != "exit" {
		var currentGopher, otherGopher *Gopher
		if turn%2 == 0 {
			currentGopher = Gopher1
			otherGopher = Gopher2
		} else {
			currentGopher = Gopher2
			otherGopher = Gopher1
		}
		fmt.Printf("\n%s's turn\n", currentGopher.name)
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		actionSli := strings.Split(strings.ToLower(strings.TrimSpace(line)), " ")
		action = actionSli[0]
		args := actionSli[1:]
		switch action {
		case "attack":
			err = attack(currentGopher, otherGopher)
		case "buy":
			if args[0] == "item" {
				err = buyItem(args[1], currentGopher)
			} else if args[0] == "weapon" {
				err = buyWeapon(args[1], currentGopher)
			} else {
				err = errors.New("Invalid command")
			}
		case "work":
			err = work(currentGopher)
		case "train":
			err = train(args[0], currentGopher)
		case "exit":
			break
		default:
			err = errors.New("Invalid command")
			fmt.Println("Options are: Attack, Buy item {item}, Buy Weapon {weapon}, Work, Train {stat}, Exit")
		}
		if err != nil {
			fmt.Println(err)
		} else {
			printStats(currentGopher)
			turn++
		}
	}
	fmt.Println("Exiting ... ")
}
func attack(attacker *Gopher, defender *Gopher) error {
	_, ok := Weapons[attacker.weapon]
	if !ok {
		panic("Invalid state: Attacker's weapon does not exist")
	} // Check that attacker's weapon exists in the map
	dmgRange := Weapons[attacker.weapon].damage
	dmgRoll := randomClosedInt(dmgRange[0], dmgRange[1])
	defender.hitpoints -= dmgRoll
	fmt.Printf("%s attacks %s for %d damage!\n", attacker.name, defender.name, dmgRoll)
	fmt.Printf("%s has %d hitpoints remaining\n", defender.name, defender.hitpoints)
	return nil
}
func buyItem(consumableName string, gopher *Gopher) error {
	consumable, ok := Consumables[consumableName]
	if !ok {
		return fmt.Errorf("%s is not a valid consumable !", consumableName)
	}
	if gopher.coins < consumable.price {
		return fmt.Errorf("%s has %d coins, but %s costs %d coins!", gopher.name, gopher.coins, consumableName, consumable.price)
	}
	gopher.hitpoints = int(math.Min(30, float64(gopher.hitpoints+consumable.hitpointsEffect)))
	gopher.coins -= consumable.price
	return nil
}
func buyWeapon(weaponName string, gopher *Gopher) error {
	weapon, ok := Weapons[weaponName]
	if !ok {
		return fmt.Errorf("%s is not a valid weapon !", weaponName)
	} // Check weapon is a vaild selection
	if gopher.coins < weapon.price {
		return fmt.Errorf("%s has %d  coins, but %s costs %d coins!", gopher.name, gopher.coins, weaponName, weapon.price)
	} // Check gopher has enough to buy chosen weapon
	if gopher.strength < weapon.strengthReq || gopher.agility < weapon.agilityReq || gopher.intellect < weapon.intelligenceReq {
		return fmt.Errorf("%s has:	%d STR | %d AGI | %d INT\n%s requires: %d STR | %d AGI | %d INT", gopher.name, gopher.strength, gopher.agility, gopher.intellect, weaponName, weapon.strengthReq, weapon.agilityReq, weapon.intelligenceReq)
	} // Check gopher meets attribute requirements to use chosen weapon
	gopher.weapon = weaponName
	gopher.coins -= weapon.price
	return nil
}
func work(gopher *Gopher) error {
	goldEarned := randomClosedInt(5, 15)
	fmt.Printf("%s earned %d gold this turn !\n", gopher.name, goldEarned)
	gopher.coins += goldEarned
	return nil
}
func train(skill string, gopher *Gopher) error {
	if gopher.coins < 5 { // Check for enough gold to train
		return fmt.Errorf("Insuffient gold for training. You have %d but you need 5", gopher.coins)
	}
	switch skill {
	case "strength":
		gopher.strength += 2
	case "agility":
		gopher.agility += 2
	case "intellect":
		gopher.intellect += 2
	default:
		return errors.New("Invalid attribute chosen") // Check if skill is valid
	}
	gopher.coins -= 5
	fmt.Printf("%s spent 5 gold to train %s!\n", gopher.name, skill)
	return nil
}
func printStats(g *Gopher) {
	fmt.Printf("The current Gopher is %s\n", g.name)
	fmt.Printf("Gopher1: HP: %d, STR: %d, AGI: %d, INT: %d, Gold: %d\n", Gopher1.hitpoints, Gopher1.strength, Gopher1.agility, Gopher1.intellect, Gopher1.coins)
	fmt.Printf("Gopher2: HP: %d, STR: %d, AGI: %d, INT: %d, Gold: %d\n", Gopher2.hitpoints, Gopher2.strength, Gopher2.agility, Gopher2.intellect, Gopher2.coins)
}
func randomClosedInt(start, end int) int {
	rge := end - start
	fmt.Println("Rge: ", rge)
	roll := rand.Intn(rge + 1)
	fmt.Println("Roll: ", roll)
	adjRoll := roll + start
	return adjRoll
}
