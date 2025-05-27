package models

import "time"

type Compromisso struct {
	ID        int
	Titulo    string
	Descricao string
	DataHora  time.Time
	Tipo      string // "Reunião" ou "Visita"
}
