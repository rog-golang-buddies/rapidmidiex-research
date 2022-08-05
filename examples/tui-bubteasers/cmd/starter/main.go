package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dbg string
}

func initialModel() model {
	m := model{
		dbg: "",
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.MouseMsg:
		return m.handleMouseMsg(msg)
	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)
	}
	return m, nil
}

func (m model) View() string {
	return m.dbg
}

func (m model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	m.dbg = "KeyMsg: " + msg.String()

	switch msg.String() {
	case "ctrl+c", "esc":
		return m, tea.Quit
	}

	// switch msg.Type {
	// case tea.KeyEsc:
	// 	return m, tea.Quit
	// }

	return m, nil
}

func (m model) handleMouseMsg(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	m.dbg = fmt.Sprintf("MouseMsg: x=[%d] y=[%d] event=[%s]", msg.X, msg.Y, tea.MouseEvent(msg).String())

	// switch msg.Type {
	// case tea.MouseLeft:
	// }

	return m, nil
}

func (m model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.dbg = fmt.Sprintf("WindowSizeMsg: w=[%d] h=[%d]", msg.Width, msg.Height)

	// [0,0]-[msg.Width-1, msg.Height-1]

	return m, nil
}

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(), // motion/hover-events
		//tea.WithMouseCellMotion(), // better support
	)
	if err := p.Start(); err != nil {
		log.Println(err)
	}
}
