package main

import "fmt"

type Checkbox interface {
	Render() string
	Toggle()
}

type WindowsCheckbox struct {
	checked bool
}

func (c *WindowsCheckbox) Render() string {
	if c.checked {
		return "Rendering checked Windows checkbox"
	}
	return "Rendering unchecked Windows checkbox"
}

func (c *WindowsCheckbox) Toggle() {
	c.checked = !c.checked
	fmt.Println("Windows checkbox toggled to:", c.checked)
}

type MacCheckbox struct {
	checked bool
}

func (c *MacCheckbox) Render() string {
	if c.checked {
		return "Rendering checked macOS checkbox"
	}
	return "Rendering unchecked macOS checkbox"
}

func (c *MacCheckbox) Toggle() {
	c.checked = !c.checked
	fmt.Println("MacOS checkbox toggled to:", c.checked)
}
