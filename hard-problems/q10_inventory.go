package main

import "fmt"

// Concepts: maps, slices, range, switch, loops, functions, conditionals
// (this one ties everything together — take your time)
//
// Task:
// Implement ProcessTransactions which takes a slice of command strings and an
// initial inventory map[string]int, applies each command IN ORDER, and
// returns the final inventory map along with a slice of any error messages
// generated along the way.
//
// Supported commands (space-separated, case-sensitive action keyword):
//   "ADD <item> <qty>"    -> increases stock of <item> by <qty>
//                             (creates the item if it doesn't exist yet)
//   "REMOVE <item> <qty>" -> decreases stock of <item> by <qty>
//                             if qty > current stock: do NOT modify stock,
//                             append error "insufficient stock: <item>"
//   "CHECK <item>"        -> no mutation; if item doesn't exist,
//                             append error "unknown item: <item>"
// Any command with an unrecognized action keyword should append error
// "unknown command: <cmd>" (use the full original command string).
//
// Hint: you'll likely want strings.Fields to split each command and
// strconv.Atoi to parse the quantity — add those imports yourself.
// Use a switch statement to dispatch on the action keyword.

func ProcessTransactions(commands []string, inventory map[string]int) (map[string]int, []string) {
	// TODO: implement
	return inventory, nil
}

func main() {
	inventory := map[string]int{"apple": 5}
	commands := []string{
		"ADD apple 10",
		"REMOVE apple 3",
		"ADD banana 20",
		"REMOVE banana 25",
		"CHECK mango",
		"WIPE apple",
	}

	finalInventory, errs := ProcessTransactions(commands, inventory)

	fmt.Println("Final inventory:")
	for item, qty := range finalInventory {
		fmt.Printf("  %-8s: %d\n", item, qty)
	}

	fmt.Println("Errors:")
	for _, e := range errs {
		fmt.Println("  -", e)
	}
}
