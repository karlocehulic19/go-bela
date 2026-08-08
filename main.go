package main

import (
	"image/color"
	"os"
	"fmt"
	tea "charm.land/bubbletea/v2"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/colorprofile"
)

var myFancyColor color.Color;

func main() {
	myFancyColor, _ = colorful.Hex("#eb4034")

	p := tea.NewProgram(model{}, tea.WithColorProfile(colorprofile.TrueColor))
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

type model struct {
    cursor   int
}

func initialModel() model { return model{}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyPressMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit
						}
					}

    return m, nil
}

func (m model) View() tea.View {
    s := `/-----\
					|     |
					|  A  |
					|     |
					\-----/`
    return tea.NewView(ansi.Style{}.ForegroundColor(myFancyColor).Styled(s))
}
