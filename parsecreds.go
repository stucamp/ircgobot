package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type creds struct {
	Nick     string
	Pass     string
	Username string
	Name     string
}

const credsfile = "./creds.json"

func openJSONfileAsByteArr(filepath string) []byte {
	file, err := ioutil.ReadFile(credsfile) // For read access.
	if err != nil {
		fmt.Println("Failed to read file urls.json")
		fmt.Println(err)
		panic(err)
	}
	return file
}
func getJSONitems(data []byte) []creds {

	var items []creds

	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println("Failed to parse the json")
		fmt.Println(err)
		panic(err)
	}
	return items
}

// GetCredsFromJSON takes a filepath to JSON file and type of source as strings, pars    es the contents, returns true if
// result was found and and URL in string form for the source
func GetCredsFromJSON() (bool, []string) {

	stuff := getJSONitems(openJSONfileAsByteArr(credsfile))
	mycreds := make([]string, 0)

	if len(stuff) > 0 {
		for k := range stuff {
			fmt.Println("Loading creds...")
			mycreds = append(mycreds, stuff[k].Nick)
			mycreds = append(mycreds, stuff[k].Pass)
			mycreds = append(mycreds, stuff[k].Username)
			mycreds = append(mycreds, stuff[k].Name)
		}
		return true, mycreds
	}

	fmt.Printf("Can't find any sources in: %s\n", credsfile)
	return false, make([]string, 0)
}
