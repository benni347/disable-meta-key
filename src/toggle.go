package main

import (
	"fmt"
	"os"
	Read "read/src/lib"
	"regexp"
)

func main() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	configFileName := "/kwinrc"
	path := configFolderPath + configFileName

	file, err := Read.ReadFile(path)
	if err != nil {
		fmt.Printf("The error of read is: %s\n", err)
	}
	var fileLineContent []byte
	/*	var lenPrevious int*/
	/*fmt.Printf(string(file) + "\n")*/
	for i := 0; i < len(file); i++ {
		if string(file[i]) == "\n" {
			fmt.Printf("%s\n", string(fileLineContent)) // This should print the line till it reaches the \n char.
			fileLineContent = nil
		} else {
			fileLineContent = append(fileLineContent, file[i])
		}
	}
}

// checkIfLinesExist checks if the lines for the meta key already exists in the config
// Parse the lines and the i from a for loop to
// return values: bool, bool
// return values: if to exist, if they are bellow one and another
// if only [ModifierOnlyShortcuts] exists return true, false
// if [ModifierOnlyShortcuts]\nMeta= exists return either true, true or false, true
// else it will return false, false
func checkIfLinesExist(fileContent []byte) (bool, bool) {
	regexLineOnlyModifier := regexp.MustCompile(`(?m)\[ModifierOnlyShortcuts]`)
	regexLineModifierAndMetaLinesBetween := regexp.MustCompile(`(?m)\[ModifierOnlyShortcuts]([a-zA-Z]*\n)*Meta=`)
	regexLineModifierAndMetaBelow := regexp.MustCompile(`(?m)\[ModifierOnlyShortcuts]\nMeta=`)
	if regexLineOnlyModifier.Find([]byte(fileContent)) != nil {
		return true, false
	} else if regexLineModifierAndMetaLinesBetween.Find([]byte(fileContent)) != nil {
		return true, true
	} else if regexLineModifierAndMetaBelow.Find([]byte(fileContent)) != nil {
		return false, true
	}
	return false, false
}
