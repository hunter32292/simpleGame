package game

import "fmt"

// Game - The Game object for simple game
type Game struct {
	Clock  int
	Player *Player
	// TODO: Add Location Tile
}

// TODO: Add Tile Struct for the current place that you are on

// Player - The Player who will be part of the simple game
type Player struct {
	Money     int
	Items     []Items
	Health    int
	LocationX int
	LocationY int
}

func (p Player) String() string {
	return fmt.Sprintf("Money: %d \nItems: %s \nHealth: %d \n", p.Money, p.Items, p.Health)
}

// Items The interface class for items that can be owned by the player
// cost - The cost to buy this item
// value - The ammount of money you can get from selling this item
// dividends - The amount of return value per clock tick of this item
// name - The Name of the item
type Items interface {
	cost()
	value()
	dividends()
	name()
}
