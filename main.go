package main

import (
	"bufio"
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
		fmt.Printf("%s's turn\n", currentGopher.name)
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		actionSli := strings.Split(strings.ToLower(strings.TrimSpace(line)), " ")
		action = actionSli[0]
		args := actionSli[1:]
		switch action {
		case "attack":
			attack(currentGopher, otherGopher)
			break
		case "buy":
			buy(args[0], currentGopher)
			break
		case "work":
			work(currentGopher)
			break
		case "train":
			train(args[0], currentGopher)
			break
		case "exit":
			break
		default:
			fmt.Println("Invalid command !")
			fmt.Println("Options are: Attack, Buy {item}, Work, Train {stat}, Exit")
		}
		turn++
	}
	fmt.Println("Exiting ... ")
}
func attack(attacker *Gopher, defender *Gopher) {

	fmt.Printf("\n%s is attacking %s with %v", attacker.name, defender.name, attacker.weapon)
}
func buy(item string, gopher *Gopher) {
	//TODO: Implement
	fmt.Println("Buying:", item)
}
func work(gopher *Gopher) {
	fmt.Println(gopher.coins)
	rand.Seed(time.Now().UnixNano())
	goldEarned := rand.Intn((15-5)+1) + 5 // (range + 1) + minimum value
	fmt.Printf("Earned %d gold this turn !\n", goldEarned)
	gopher.coins += goldEarned
	fmt.Println(gopher.coins)
}
func train(skill string, gopher *Gopher) {
	//TODO: Implement
	fmt.Println("Training:", skill)
}
