package main

import (
	"bufio"
	"errors"
	"fmt"
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
	weapon:    Weapons["bare-handed"],
	coins:     20,
}
var Gopher2 = &Gopher{
	name:      "gopher2",
	hitpoints: 30,
	weapon:    Weapons["bare-handed"],
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
			break
		case "buy":
			err = buy(args[0], currentGopher)
			break
		case "work":
			err = work(currentGopher)
			break
		case "train":
			err = train(args[0], currentGopher)
			break
		case "exit":
			break
		default:
			fmt.Println("Invalid command !")
			fmt.Println("Options are: Attack, Buy {item}, Work, Train {stat}, Exit")
		}
		if err != nil {
			fmt.Println(err)
		} else {
			turn++
		}
	}
	fmt.Println("Exiting ... ")
}
func attack(attacker *Gopher, defender *Gopher) error {
	// return errors.New("Something went wrong")
	dmgRange := attacker.weapon.damage
	dmgRoll := rand.Intn((dmgRange[1]-dmgRange[0])+1) + dmgRange[0]
	defender.hitpoints -= dmgRoll
	fmt.Printf("%s attacks %s for %d damage!\n", attacker.name, defender.name, dmgRoll)
	fmt.Printf("%s has %d hitpoints remaining\n", defender.name, defender.hitpoints)
	return nil
}
func buy(item string, gopher *Gopher) error {
	//TODO: Implement
	fmt.Println("Buying:", item)
	return nil
}
func work(gopher *Gopher) error {
	goldEarned := rand.Intn((15-5)+1) + 5 // (range + 1) + minimum value
	fmt.Printf("Earned %d gold this turn !\n", goldEarned)
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
	return nil
}
