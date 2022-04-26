package main

type Scrutin struct {
	date      string
	titre     string
	demandeur string
}

func CreateScrutin(scrutinMap map[string]interface{}) (Scrutin, error) {

	scrutin := Scrutin{}

	scrutin.date = scrutinMap["scrutin"].(map[string]interface{})["dateScrutin"].(string)

	scrutin.titre = scrutinMap["scrutin"].(map[string]interface{})["titre"].(string)

	scrutin.demandeur = scrutinMap["scrutin"].(map[string]interface{})["demandeur"].(map[string]interface{})["texte"].(string)

	return scrutin, nil
}
