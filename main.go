package main

import (
	"flag"
	_ "flag"
	"fmt"
	enigma "github.com/Braden-Griebel/gonigma/machine"
	"io"
	"os"
)

func main() {
	var (
		rotorSpecification     string
		rotorSettings          string
		reflectorSpecification string
		reflectorSetting       string
		plugboardSetting       string
	)
	// Parse the flags
	flag.StringVar(&rotorSpecification, "rotors", "I,II,III", "Choose the rotors to use, should be 3 comma seperated roman numerals (I-VIII)")
	flag.StringVar(&rotorSpecification, "r", "I,II,III", "Choose the rotors to use, should be 3 comma seperated roman numerals (I-VIII)")

	flag.StringVar(&rotorSettings, "rotor-settings", "a,a,a", "Choose the rotor setting, should be 3 comma seperated characters")
	flag.StringVar(&rotorSettings, "s", "a,a,a", "Choose the rotor setting, should be 3 comma seperated characters")

	flag.StringVar(&reflectorSpecification, "reflector", "A", "Choose the reflector, can be A, B or C (single character)")
	flag.StringVar(&reflectorSpecification, "f", "A", "Choose the reflector, can be A, B or C (single character)")

	flag.StringVar(&reflectorSetting, "reflector-settings", "a", "Choose the reflector setting, should be a single character")
	flag.StringVar(&reflectorSetting, "fs", "a", "Choose the reflector setting, should be a single character")

	flag.StringVar(&plugboardSetting, "plugboard-setting", "",
		"Choose the plugboard settings, should be a comma seperated list of 'wires', where each wire is two characters seperated by a '-'")
	flag.StringVar(&plugboardSetting, "p", "",
		"Choose the plugboard settings, should be a comma seperated list of 'wires', where each wire is two characters seperated by a '-'")
	// Change the usage order
	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Println("Usage of gonigma: gonigma [options] filename")
		fmt.Println("If no filename is provided, will attempt to read stdin")
		fmt.Println("Options:")
		order := []string{"rotors", "r", "rotor-settings", "s", "reflector", "f", "reflector-settings", "fs", "plugboard-setting", "p"}
		for _, name := range order {
			flagVal := flagSet.Lookup(name)
			fmt.Printf("-%s\n", flagVal.Name)
			fmt.Printf("\t%s\n", flagVal.Usage)
		}
	}

	flag.Parse()

	// Create the desired engima machine setup
	machine := enigma.NewMachine(rotorSpecification,
		rotorSettings, reflectorSpecification, reflectorSetting, plugboardSetting)

	// Read the provided file into a string
	args := flag.Args()
	var input string
	if len(args) > 0 {
		dat, err := os.ReadFile(flag.Args()[0])
		if err != nil {
			_ = fmt.Errorf("Error reading file! Error: %w", err)
		}
		input = string(dat)
	} else {
		dat, err := io.ReadAll(os.Stdin)
		input = string(dat)
		if err != nil {
			_ = fmt.Errorf("Error reading Stdin! Error: %w", err)
		}

	}
	output := machine.Translate(input)
	fmt.Print(output)
}
