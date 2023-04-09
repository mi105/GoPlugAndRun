package main

import (
	"awesomeProject/animal"
	"awesomeProject/types"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"plugin"
	"strconv"
	"strings"
	"time"
)

func main() {
	animalMap := make(map[types.AnimalType]animal.Animal)

	// Load plugins specified in the plugins.txt file
	file, err := os.Open("plugins.txt")
	if err != nil {
		fmt.Println("Error reading plugins.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		animalType, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error parsing animal type:", err)
			continue
		}

		libName := parts[1]
		symName := parts[2]

		p, err := plugin.Open("bin/" + libName)
		if err != nil {
			fmt.Println("Error loading plugin:", err)
			continue
		}

		s, err := p.Lookup(symName)
		if err != nil {
			fmt.Println("Error looking up symbol:", err)
			continue
		}

		// Update the type assertion to match the exported variable
		symbol, ok := s.(*animal.Animal)
		if !ok {
			fmt.Println("Invalid symbol for", symName)
			continue
		}

		animalMap[types.AnimalType(animalType)] = *symbol

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning plugins.txt:", err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	animalType := types.AnimalType(rand.Intn(3)) // (determined at run-time)
	fmt.Println(animalType)
	Speak(animalMap, animalType)
}

func Speak(animalMap map[types.AnimalType]animal.Animal, animalType types.AnimalType) {
	if animal, ok := animalMap[animalType]; ok {
		animal.Speak() // Calls the appropriate function at run-time
	} else {
		fmt.Printf("Unknown animal type: %d\n", animalType)
	}
}
