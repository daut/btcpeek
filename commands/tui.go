package commands

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "\n" +
		"  Welcome to btcpeek 🪙\n\n" +
		"  Bitcoin blockchain explorer\n\n" +
		"  Press 'q' or Ctrl+C to quit\n\n"
}

func NewTUI() error {
	opts := tea.WithAltScreen()
	_, err := tea.NewProgram(model{}, opts).Run()
	if err != nil {
		return err
	}
	return nil
}
