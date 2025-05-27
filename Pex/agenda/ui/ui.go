package ui

import (
	"agenda/models"
	"agenda/storage"
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func RunUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Agenda de Compromissos")

	idEntry := widget.NewEntry()
	tituloEntry := widget.NewEntry()
	descricaoEntry := widget.NewMultiLineEntry()
	dataEntry := widget.NewEntry()
	horaEntry := widget.NewEntry()
	tipoSelect := widget.NewSelect([]string{"Reunião", "Visita"}, func(s string) {})

	filtroDataEntry := widget.NewEntry()

	compromissos := storage.GetCompromissos()

	lista := widget.NewList(
		func() int { return len(compromissos) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			c := compromissos[i]
			o.(*widget.Label).SetText(fmt.Sprintf("%d: %s - %s", c.ID, c.Titulo, c.Tipo))
		})

	atualizarLista := func() {
		compromissos = storage.GetCompromissos()
		lista.Length = func() int { return len(compromissos) }
		lista.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			c := compromissos[i]
			o.(*widget.Label).SetText(fmt.Sprintf("%d: %s - %s", c.ID, c.Titulo, c.Tipo))
		}
		lista.Refresh()
	}

	adicionarBtn := widget.NewButton("Adicionar", func() {
		dataHoraStr := dataEntry.Text + " " + horaEntry.Text
		dataHora, err := time.Parse("2006-01-02 15:04", dataHoraStr)
		if err != nil {
			fmt.Println("Erro ao parsear data/hora:", err)
			return
		}
		novo := models.Compromisso{
			Titulo:    tituloEntry.Text,
			Descricao: descricaoEntry.Text,
			DataHora:  dataHora,
			Tipo:      tipoSelect.Selected,
		}
		storage.AddCompromisso(novo)
		storage.SaveCompromissos("compromissos.json")
		atualizarLista()
	})

	editarBtn := widget.NewButton("Editar por ID", func() {
		id, err := strconv.Atoi(idEntry.Text)
		if err != nil {
			fmt.Println("ID inválido")
			return
		}
		dataHoraStr := dataEntry.Text + " " + horaEntry.Text
		dataHora, err := time.Parse("2006-01-02 15:04", dataHoraStr)
		if err != nil {
			fmt.Println("Erro ao parsear data/hora:", err)
			return
		}
		editado := models.Compromisso{
			Titulo:    tituloEntry.Text,
			Descricao: descricaoEntry.Text,
			DataHora:  dataHora,
			Tipo:      tipoSelect.Selected,
		}
		storage.EditCompromisso(id, editado)
		storage.SaveCompromissos("compromissos.json")
		atualizarLista()
	})

	deletarBtn := widget.NewButton("Excluir por ID", func() {
		id, err := strconv.Atoi(idEntry.Text)
		if err != nil {
			fmt.Println("ID inválido")
			return
		}
		storage.RemoveCompromisso(id)
		storage.SaveCompromissos("compromissos.json")
		atualizarLista()
	})

	filtrarBtn := widget.NewButton("Filtrar por Data", func() {
		data := filtroDataEntry.Text
		filtrados := storage.FilterCompromissosPorData(data)
		lista.Length = func() int { return len(filtrados) }
		lista.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			c := filtrados[i]
			o.(*widget.Label).SetText(fmt.Sprintf("%d: %s - %s", c.ID, c.Titulo, c.Tipo))
		}
		lista.Refresh()
	})

	resetFiltroBtn := widget.NewButton("Resetar Filtro", func() {
		atualizarLista()
	})

	form := container.NewVBox(
		widget.NewLabel("ID (para editar/excluir):"), idEntry,
		widget.NewLabel("Título:"), tituloEntry,
		widget.NewLabel("Descrição:"), descricaoEntry,
		widget.NewLabel("Data (YYYY-MM-DD):"), dataEntry,
		widget.NewLabel("Hora (HH:MM):"), horaEntry,
		widget.NewLabel("Tipo:"), tipoSelect,
		adicionarBtn,
		editarBtn,
		deletarBtn,
		widget.NewSeparator(),
		widget.NewLabel("Filtro - Data (YYYY-MM-DD):"), filtroDataEntry,
		filtrarBtn,
		resetFiltroBtn,
	)

	content := container.NewHSplit(form, lista)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 400))
	myWindow.ShowAndRun()
}
