package main
import (
	"fmt"
	"os"
	"strings"
)

// askForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func askForConfirmation() bool {
	var sure string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&sure)
	if err != nil {
		panic(err)
	}

	sure = strings.TrimSpace(sure)
	sure = strings.ToLower(sure)

	if sure == "y" || sure == "yes" {
		return true
	}
	return false
}


func createDirectory(directoryPath string) {
	//choose your permissions well
	pathErr := os.MkdirAll(directoryPath,0740)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}

func main() {
//get playbook name
	fmt.Print("Enter a Name: ")
	var playbookName string
	fmt.Scanln(&playbookName)
	playbookName = strings.TrimSpace(playbookName)
//get Path
	var path string
	path = os.Getenv("PWD")
	path = strings.TrimSpace(path)
	var fullPath = (path + "/" + playbookName)
    fullPath = strings.TrimSpace(fullPath)

	fmt.Println("Should I create", fullPath, "?")
	if askForConfirmation() {
		createDirectory(fullPath)
	}

	fmt.Println("Should I create subfolders?")
	if askForConfirmation() {
		var newPath = ( fullPath + "/doc")
		createDirectory(newPath)
	}
}

