package main

import "fmt"

type Button interface {
	Render() string
	OnClick()
}

type WindowsButton struct{}

func (b *WindowsButton) Render() string {
	return "Rendering Windows button"
}

func (b *WindowsButton) OnClick() {
	fmt.Println("Windows button clicked!")
}

type MacButton struct{}

func (b *MacButton) Render() string {
	return "Rendering macOS button"
}

func (b *MacButton) OnClick() {
	fmt.Println("MacOS button clicked!")
}
