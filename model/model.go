package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type step int

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }
func (i item) Description() string { return i.desc }

const (
	stepInput step = iota
	stepList1
	stepList2
	stepList3
	stepDone
	defaultWidth  = 30
	defaultHeight = 20
)

type Model struct {
	step          step
	textInput     textinput.Model
	frameworkList list.Model
	cssList       list.Model
	languageList  list.Model
	Name          string
	Framework     string
	Css           string
	Lang          string
	quitting      bool
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Default:(my-project)"
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 30

	return Model{
		step:          stepInput,
		textInput:     ti,
		frameworkList: frameworkList(),
		cssList:       cssList(),
		languageList:  languageList(),
		quitting:      false,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch m.step {

		case stepInput:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				m.Name = m.Input()
				if m.Name == "" {
					m.Name = "my-project"
				}
				m.step = stepList1
				return m, nil
			default:
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}

		case stepList1:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.frameworkList.SelectedItem().(item)
				if ok {
					m.Framework = i.title
				}
				m.step = stepList2
				return m, nil
			default:
				var cmd tea.Cmd
				m.frameworkList, cmd = m.frameworkList.Update(msg)
				return m, cmd
			}

		case stepList2:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.cssList.SelectedItem().(item)
				if ok {
					m.Css = i.title
				}
				m.step = stepList3
				return m, nil
			default:
				var cmd tea.Cmd
				m.cssList, cmd = m.cssList.Update(msg)
				return m, cmd
			}
		case stepList3:
			switch msg.String() {
			case "ctrl+c", "q":
				m.quitting = true
				return m, tea.Quit
			case "enter":
				i, ok := m.languageList.SelectedItem().(item)
				if ok {
					m.Lang = i.title
				}
				m.step = stepDone
				return m, nil
			default:
				var cmd tea.Cmd
				m.languageList, cmd = m.languageList.Update(msg)
				return m, cmd
			}
		case stepDone:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	switch m.step {
	case stepInput:
		return fmt.Sprintf(
			"What is the name of your project?\n\n%s",
			m.textInput.View(),
		)
	case stepList1:
		return fmt.Sprintf(
			"What framework will you be using?\n\n%s",
			m.frameworkList.View(),
		)
	case stepList2:
		return fmt.Sprintf(
			"What CSS framework will you be using?\n\n%s",
			m.cssList.View(),
		)
	case stepList3:
		return fmt.Sprintf(
			"What language will you be using?\n\n%s",
			m.languageList.View(),
		)
	case stepDone:
		return fmt.Sprintf(
			"Project Name: %s\n"+
				"Framework: %s\n"+
				"CSS: %s\n"+
				"Language: %s\n"+
				"\n\n press enter to build your project",
			m.Name,
			m.Framework,
			m.Css,
			m.Lang,
		)
	default:
		return "Loading..."
	}
}

func (m Model) Input() string {
	return m.textInput.Value()
}
