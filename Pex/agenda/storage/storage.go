package storage

import (
	"agenda/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

var compromissos []models.Compromisso

func LoadCompromissos(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			compromissos = []models.Compromisso{}
			return nil
		}
		return err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	return json.Unmarshal(bytes, &compromissos)
}

func SaveCompromissos(filename string) error {
	bytes, err := json.MarshalIndent(compromissos, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0644)
}

func GetCompromissos() []models.Compromisso {
	return compromissos
}

func AddCompromisso(c models.Compromisso) {
	c.ID = len(compromissos) + 1
	compromissos = append(compromissos, c)
}

func RemoveCompromisso(id int) {
	for i, c := range compromissos {
		if c.ID == id {
			compromissos = append(compromissos[:i], compromissos[i+1:]...)
			break
		}
	}
}

func EditCompromisso(id int, novo models.Compromisso) {
	for i, c := range compromissos {
		if c.ID == id {
			novo.ID = id
			compromissos[i] = novo
			break
		}
	}
}

func FilterCompromissosPorData(data string) []models.Compromisso {
	var filtrados []models.Compromisso
	for _, c := range compromissos {
		if c.DataHora.Format("2006-01-02") == data {
			filtrados = append(filtrados, c)
		}
	}
	return filtrados
}
