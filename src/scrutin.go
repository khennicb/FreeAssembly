package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Scrutin struct {
	uid       string
	date      string
	titre     string
	demandeur string
}

func getMapInside(myMap map[string]interface{}, key string) (map[string]interface{}, error) {
	if reflect.ValueOf(myMap[key]).Kind() == reflect.Map {
		return myMap[key].(map[string]interface{}), nil
	} else {
		return nil, errors.New("no value for " + key)
	}
}

func getSomethingInside(myMap map[string]interface{}, keys ...string) (interface{}, error) {
	for _, key := range keys[:len(keys)-1] {
		var err error
		myMap, err = getMapInside(myMap, key)
		if err != nil {
			return "", err
		}
	}
	return myMap[keys[len(keys)-1]], nil
}

func getStringInside(myMap map[string]interface{}, keys ...string) (string, error) {

	myDeepSomething, err := getSomethingInside(myMap, keys...)
	if err != nil {
		return "", err
	}
	if reflect.ValueOf(myDeepSomething).Kind() == reflect.String {
		return myDeepSomething.(string), nil
	} else {
		return "", errors.New("no string for the key " + keys[len(keys)-1])
	}

}

func CreateScrutin(scrutinMap map[string]interface{}) (Scrutin, error) {

	scrutinMap = scrutinMap["scrutin"].(map[string]interface{})

	scrutin := Scrutin{}

	var err error
	scrutin.uid, err = getStringInside(scrutinMap, "uid")
	if err != nil {
		fmt.Println(err)
		return Scrutin{}, errors.New("no value for uid")
	}

	scrutin.date, err = getStringInside(scrutinMap, "dateScrutin")
	if err != nil {
		fmt.Println(err)
		return Scrutin{}, errors.New("no value for dateScrutin")
	}

	scrutin.titre, err = getStringInside(scrutinMap, "titre")
	if err != nil {
		fmt.Println(err)
		return Scrutin{}, errors.New("no value for titre")
	}

	scrutin.demandeur, err = getStringInside(scrutinMap, "demandeur", "texte")
	if err != nil {
		fmt.Println(err)
		return Scrutin{}, errors.New("no value for demandeur")
	}

	return scrutin, nil
}

func (s Scrutin) prettyPrint() {
	fmt.Printf("Scrutin details : \n")
	fmt.Printf("  uid: %s\n", s.uid)
	fmt.Printf("  date: %s\n", s.date)
	fmt.Printf("  titre: %s\n", s.titre)
	fmt.Printf("  demandeur: %s\n", s.demandeur)
}
