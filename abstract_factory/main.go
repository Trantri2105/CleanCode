package main

import "fmt"

type Application struct {
	button   Button
	checkbox Checkbox
}

func NewApplication(factory GUIFactory) *Application {
	return &Application{
		button:   factory.CreateButton(),
		checkbox: factory.CreateCheckbox(),
	}
}

func (a *Application) RenderUI() {
	fmt.Println(a.button.Render())
	fmt.Println(a.checkbox.Render())
}

func (a *Application) Click() {
	a.button.OnClick()
}

func (a *Application) ToggleCheckbox() {
	a.checkbox.Toggle()
}

func main() {
	// Create application GUI Windows
	windowsFactory := &WindowsFactory{}
	windowsApp := NewApplication(windowsFactory)

	fmt.Println("=== Windows UI ===")
	windowsApp.RenderUI()
	windowsApp.Click()
	windowsApp.ToggleCheckbox()

	fmt.Println()

	// Create application with GUI macOS
	macFactory := &MacFactory{}
	macApp := NewApplication(macFactory)

	fmt.Println("=== Mac UI ===")
	macApp.RenderUI()
	macApp.Click()
	macApp.ToggleCheckbox()
}
