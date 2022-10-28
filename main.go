package main

/* The game is a turn-based one. There are two gophers and they can each decide what to do on their turn. Each gopher starts with 30 hitpoints, 20 gold and all their attributes are 0.

The possible actions are:
attack - attack the other gopher with the weapon you have equipped at the moment. If none is equipped, then you are attacking bare-handed for a damage of 1 hitpoint.
Choose the actual damage dealt at random based on the weapon’s damage interval
buy weapon <item> - spend the coins necessary to buy the given weapon based on its cost. You equip it over your current weapon. You can’t buy weapons you are still illegible to use due to insufficient stats.
buy consumable <item> - spend the coins necessary to buy the given consumable based on its cost and use it directly.
work - spend the turn working for the local warlord and gain anywhere between 5 and 15 coins (picked at random)
train <skill> - train a given attribute (strength, agility or intellect) and increase it by 2. Training costs 5 gold.
exit - exits the game

The shop has the following items for sale with unlimited supply of them:
Consumables:
small_health_potion - consumable - 5 gold
hitpointsEffect - 5
medium_health_potion - consumable - 9 gold
hitpointsEffect - 10
big_health_potion - consumable - 18 gold
hitpointsEffect - 20

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

Extra challenges:
Implement a game log which prints all events which occur throughout the game (see examples below)
Implement consumables which give you an attribute boost (strength, agility, intellect) for a limited duration (e.g. 3 turns). This allows you to buy and use items without having the proper training yet but only for a limited duration
*/

func main() {

}
