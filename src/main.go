package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var defaultDataDirectory = "data/"

func main() {

	dataDirectory := ""
	nbArg := len(os.Args)
	fmt.Printf("number of args : %d\n", nbArg)
	if nbArg == 1 {
		dataDirectory = defaultDataDirectory+"Scrutins_XV.json/json/"
		fmt.Printf("no arg, looking into %s\n", dataDirectory)
	} else if nbArg == 2 {
		dataDirectory = os.Args[1]
		fmt.Printf("directory specified in argument is : %s\n", dataDirectory)
	} else {
		fmt.Printf("found 2 args, 0 or 1 expected\n")
		return
	}

	fileNames, err := getAllFilesInDir(dataDirectory)
	if err != nil {
		return
	}

	// Lets try that to see what is inside !

	mapScrutin, err := convertFileIntoMap(dataDirectory + "/" + fileNames[0])
	if err != nil {
		fmt.Printf("could not convert the file %s\n", fileNames[0])
		return
	}

	scrutin, err := CreateScrutin(mapScrutin)

	fmt.Printf("Scrutin details : \n")
	fmt.Printf("  date: %s\n", scrutin.date)
	fmt.Printf("  titre: %s\n", scrutin.titre)
	fmt.Printf("  demandeur: %s\n", scrutin.demandeur)

	// Checking every single file to count the errors

	errorCount := 0
	for _, fileName := range fileNames {

		mapScrutin, err := convertFileIntoMap(dataDirectory + "/" + fileName)
		if err != nil {
			fmt.Printf("could not convert the file %s\n", fileName)
			return
		}

		s, err := CreateScrutin(mapScrutin)

		fmt.Println(s.date)

		if err != nil {
			fmt.Printf("%v\n", err)
			errorCount++
		}
	}

	fmt.Printf("Error count : %d out of %d\n", errorCount, len(fileNames))

}

func getAllFilesInDir(dir string) ([]string, error) {

	f, err := os.Open(dir)
	if err != nil {
		fmt.Printf("could not open the directory %s\n", dir)
		return nil, errors.New("could not open the directory")
	}

	files, err := f.Readdirnames(0)
	if err != nil {
		fmt.Printf("could not read the directory %s\n", dir)
		return nil, errors.New("could not read the directory")
	}

	return files, nil
}

func convertFileIntoMap(filePath string) (map[string]interface{}, error) {
	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("could read the file %s\n", filePath)
		fmt.Printf("error : %v\n", err)
		return nil, errors.New("could read the file")
	}

	var scrutin map[string]interface{}
	if err := json.Unmarshal([]byte(text), &scrutin); err != nil {
		fmt.Println("failed to unmarshal json")
		return nil, err
	}

	return scrutin, nil
}
