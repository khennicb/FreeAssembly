package main

import (
	"errors"
	"reflect"
)

type Scrutin struct {
	date      string
	titre     string
	demandeur string
}

func CreateScrutin(scrutinMap map[string]interface{}) (Scrutin, error) {

	scrutinMap = scrutinMap["scrutin"].(map[string]interface{})

	scrutin := Scrutin{}

	if reflect.ValueOf(scrutinMap["dateScrutin"]).Kind() == reflect.String {
		scrutin.date = scrutinMap["dateScrutin"].(string)
	} else {
		println("no value for date")
		return Scrutin{}, errors.New("no value for date")
	}

	if reflect.ValueOf(scrutinMap["titre"]).Kind() == reflect.String {
		scrutin.titre = scrutinMap["titre"].(string)
	} else {
		return Scrutin{}, errors.New("no value for titre")
	}

	if reflect.ValueOf(scrutinMap["demandeur"]).Kind() == reflect.Map {
		if reflect.ValueOf(scrutinMap["demandeur"].(map[string]interface{})["texte"]).Kind() == reflect.String {
			scrutin.demandeur = scrutinMap["demandeur"].(map[string]interface{})["texte"].(string)
		} else {
			return Scrutin{}, errors.New("no value for demandeur")
		}
	} else {
		return Scrutin{}, errors.New("no value for demandeur")
	}

	return scrutin, nil
}
