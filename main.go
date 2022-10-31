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
var Gopher1 = &Gopher{ // 2 static variables that point to a Gopher struct. Defines initial game state.
	name:      "gopher1",
	hitpoints: 30,
	weapon:    "bare-hands",
	coins:     20,
}
var Gopher2 = &Gopher{
	name:      "gopher2",
	hitpoints: 30,
	weapon:    "bare-hands",
	coins:     20,
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Changes the seed of the rand library
	fmt.Println("Welcome to a game of Gopher RPG")
	r := bufio.NewReader(os.Stdin)
	handleAction(r)
	if Gopher1.hitpoints <= 0 || Gopher2.hitpoints <= 0 {
		winner := getWinner()
		fmt.Printf("Game over! %s is the winner!\n", winner)
	} else {
		fmt.Println("Exiting ... ")
	}
}

func handleAction(r *bufio.Reader) { // Function to handle user input and call corresponding functions.
	action := ""
	turn := 0
	for action != "exit" && (Gopher1.hitpoints > 0 && Gopher2.hitpoints > 0) { // When user inputs "exit" exit the game, otherwise continue accepting input
		var currentGopher, otherGopher *Gopher
		if turn%2 == 0 { // Alternate turns between the 2 gophers
			currentGopher = Gopher1
			otherGopher = Gopher2
		} else {
			currentGopher = Gopher2
			otherGopher = Gopher1
		}
		fmt.Printf("\n--- %s's turn ---\n", currentGopher.name)
		fmt.Print("> ") //Indicate in terminal which lines are user's input
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		actionSli := strings.Split(strings.ToLower(strings.TrimSpace(line)), " ") // Break the user input into string and slice of arguments (if applicable), to be passed as arguments to appropriate functions
		action = actionSli[0]
		args := actionSli[1:]
		switch action { // Call the function corresponding to the user input
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
			// fmt.Println("Exiting ... ")
			break
		default: // If the first word of the user input was not one of the above, print the following.
			//Since the error was now defined, the turn is not incremented.
			err = errors.New("Invalid command")
			fmt.Println("Options are: Attack, Buy item {item}, Buy Weapon {weapon}, Work, Train {stat}, Exit")
		}
		if err != nil { // If an error was returned from a called function, print it and reprompt.
			fmt.Println(err)
		} else {
			printStats(currentGopher) // Print the current game state (stats of the 2 gophers)
			turn++                    // Increment the turn counter
		}
	}
}
func attack(attacker *Gopher, defender *Gopher) error { // Function for gopher to attack it's opponent, with damage based on currently equipped weapon
	_, ok := Weapons[attacker.weapon]
	if !ok {
		panic("Invalid state: Attacker's weapon does not exist")
	} // Check that attacker's weapon exists in the map
	dmgRange := Weapons[attacker.weapon].damage
	dmgRoll := randomClosedInt(dmgRange[0], dmgRange[1])
	defender.hitpoints -= dmgRoll
	fmt.Printf("%s attacks %s with their %s for %d damage!\n", attacker.name, defender.name, attacker.weapon, dmgRoll)
	fmt.Printf("%s has %d hitpoints remaining\n", defender.name, defender.hitpoints)
	return nil
}
func buyItem(consumableName string, gopher *Gopher) error { // Function to recover some of a gopher's hitpoints
	consumable, ok := Consumables[consumableName]
	if !ok {
		return fmt.Errorf("%s is not a valid consumable !", consumableName)
	} // Check for valid consumable
	if gopher.coins < consumable.price {
		return fmt.Errorf("%s has %d coins, but %s costs %d coins!", gopher.name, gopher.coins, consumableName, consumable.price)
	} // Check gopher has enough to buy chosen consumable
	fmt.Printf("%s bougt a %s for %d gold and recovered %d hitpoints!\n", gopher.name, consumableName, consumable.price, int(math.Min(float64(consumable.hitpointsEffect), float64(30-gopher.hitpoints))))
	gopher.hitpoints = int(math.Min(30, float64(gopher.hitpoints+consumable.hitpointsEffect))) // Prevent overhealing by only adding up to the max hp if the consumable would heal to more than max
	gopher.coins -= consumable.price
	return nil
}
func buyWeapon(weaponName string, gopher *Gopher) error { // Function to change a gopher's weapon
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
	fmt.Printf("%s bought %s for %d gold and equipped it!\n", gopher.name, weaponName, weapon.price)
	gopher.weapon = weaponName
	gopher.coins -= weapon.price
	return nil
}
func work(gopher *Gopher) error { // Function to give a gopher a random amount of gold [5,15]
	goldEarned := randomClosedInt(5, 15)
	fmt.Printf("%s earned %d gold this turn !\n", gopher.name, goldEarned)
	gopher.coins += goldEarned
	return nil
}
func train(skill string, gopher *Gopher) error { // Function to increase attibutes of a gopher
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
func printStats(g *Gopher) { // Function to print stats of each gopher at every turn
	fmt.Printf("\nGopher1: HP: %d, STR: %d, AGI: %d, INT: %d, Gold: %d\n", Gopher1.hitpoints, Gopher1.strength, Gopher1.agility, Gopher1.intellect, Gopher1.coins)
	fmt.Printf("Gopher2: HP: %d, STR: %d, AGI: %d, INT: %d, Gold: %d\n", Gopher2.hitpoints, Gopher2.strength, Gopher2.agility, Gopher2.intellect, Gopher2.coins)
}
func randomClosedInt(start, end int) int { // Function to return a random interger between 2 fully open (inclusive) intervals
	rge := end - start         // Find the range between start and end
	roll := rand.Intn(rge + 1) // Roll a random number between 0 and the end of the range + 1 to account for rand.Intn being half-open
	adjRoll := roll + start    // Offset the number by the specified start number
	return adjRoll
}
func getWinner() string {
	if Gopher2.hitpoints <= 0 {
		return Gopher1.name
	} else if Gopher1.hitpoints <= 0 {
		return Gopher2.name
	}
	panic("Invalid state: getWinner triggered without a winner")
}
