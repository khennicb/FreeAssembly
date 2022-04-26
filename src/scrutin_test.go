package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateScrutin(t *testing.T) {
	inputjson := `
	{"scrutin": 
		{"@xmlns": "http://schemas.assemblee-nationale.fr/referentiel", 
		"@xmlns:xsi": "http://www.w3.org/2001/XMLSchema-instance", 
		"uid": "VTANR5L15V1", 
		"numero": "1", 
		"organeRef": "PO717460", 
		"legislature":"15", 
		"sessionRef": "SCR5A2017E1", 
		"seanceRef": "RUANR5L15S2017IDS20603", 
		"dateScrutin": "2017-07-04", 
		"quantiemeJourSeance": "1", 
		"typeVote": 
			{"codeTypeVote": "SPS", 
			"libelleTypeVote": "scrutin public solennel", 
			"typeMajorite": "majorit\u00e9 absolue des suffrages exprim\u00e9s"}, 
		"sort": 
			{"code": "adopt\u00e9", 
			"libelle": "l'Assembl\u00e9e nationale a adopt\u00e9"}, 
		"titre": "la declaration de politique generale du Gouvernement de M. Edouard Philippe (application de l'article 49, alinea premier, de la Constitution).", 
		"demandeur": 
			{"texte": "Conference des Presidents", 
			"referenceLegislative": null}
		}
	}`

	expectedScrutin := Scrutin{
		date:      "2017-07-04",
		titre:     "la declaration de politique generale du Gouvernement de M. Edouard Philippe (application de l'article 49, alinea premier, de la Constitution).",
		demandeur: "Conference des Presidents",
	}

	var inputmap map[string]interface{}
	if err := json.Unmarshal([]byte(inputjson), &inputmap); err != nil {
		fmt.Println("failed to unmarshal json")
		t.Errorf("Test failed because of a failed unmarshal")
	}

	scrutin, err := CreateScrutin(inputmap)
	if err != nil {
		t.Errorf("Test failed. The function CreateScrutin returned an error")
	}

	if scrutin.date != expectedScrutin.date {
		t.Errorf("Test failed on %+v. got %+v; wanted %+v", expectedScrutin, scrutin.date, expectedScrutin.date)
	}

	if scrutin.titre != expectedScrutin.titre {
		t.Errorf("Test failed on %+v. got %+v; wanted %+v", expectedScrutin, scrutin.titre, expectedScrutin.titre)
	}

	if scrutin.demandeur != expectedScrutin.demandeur {
		t.Errorf("Test failed on %+v. got %+v; wanted %+v", expectedScrutin, scrutin.demandeur, expectedScrutin.demandeur)
	}

}
