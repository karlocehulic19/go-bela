package main

import (
	"fmt"
	"log"
	"os"
	"errors"
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/exp/charmtone"
)

var inputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(charmtone.Cherry.Hex()))

func main() {
	p := tea.NewProgram(model{})
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

func get_card_string(card_type string) (string, error) {
	allowed_card_types := []string{"A", "K", "Q", "J", "X", "XI", "VIII", "VII"}
	if !slices.Contains(allowed_card_types, card_type) {
		return "", errors.New("An invalid card type is specified for card string")
	}

	top_named_line := ""
	bottom_named_line := ""
	switch card_type {
	case "XI":
		top_named_line = "|%s    |\n"
		bottom_named_line = "|    %s|\n"
	case "VIII":
		top_named_line = "|%s  |\n"
		bottom_named_line = "|  %s|\n"
	case "VII":
		top_named_line = "|%s   |\n"
		bottom_named_line = "|   %s|\n"
	default:
		top_named_line = "| %s    |\n"
		bottom_named_line = "|    %s |\n"
	}

	s := fmt.Sprintf(
		"/------\\\n" +
		top_named_line +
		"|      |\n" +
		"|      |\n" +
		bottom_named_line +
	 "\\------/\n", card_type, card_type)

	return s, nil;
}

func (m model) View() tea.View {
		card_s, err := get_card_string("VII")
		if err != nil {
			log.Fatal(err)
		}
		card_s_2, err_2 := get_card_string("X")
		if err_2 != nil {
			log.Fatal(err)
		}

		card_s_3, err_3 := get_card_string("A")
		if err_3 != nil {
			log.Fatal(err)
		}
		// Create some layers.
		a := lipgloss.NewLayer(inputStyle.Render(card_s)).X(0)
		b := lipgloss.NewLayer(inputStyle.Render(card_s_2)).X(6)
		c := lipgloss.NewLayer(inputStyle.Render(card_s_3)).X(12)
		layers := []*lipgloss.Layer {a, b, c}

		// Composite 'em and render.
		outter_layer := lipgloss.NewLayer(inputStyle.Width(30).String())
		compositor := lipgloss.NewCompositor(outter_layer.AddLayers(layers...));
		output := compositor.Render()
    return tea.NewView(output)
}
