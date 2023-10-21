package main

import "fmt"

type Light struct {
	On bool
}

func NewLight() *Light {
	return &Light{On: false}
}

func (l *Light) TurnOn() {
	fmt.Println("Light off")
	l.On = true
}

func (l *Light) TurnOff() {
	fmt.Println("Light on")
	l.On = false
}

type Command interface {
	Execute()
}

type TurnOnCommand struct {
	light *Light
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
	return &TurnOnCommand{light: light}
}

func (tc *TurnOnCommand) Execute() {
	tc.light.TurnOn()
}

type TurnOffCommand struct {
	light *Light
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
	return &TurnOffCommand{light: light}
}

func (tc *TurnOffCommand) Execute() {
	tc.light.TurnOff()
}

type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func main() {
	light := NewLight()
	remoteControl := &RemoteControl{}

	turnOnCommand := NewTurnOnCommand(light)
	turnOffCommand := NewTurnOffCommand(light)

	remoteControl.SetCommand(turnOnCommand)
	remoteControl.PressButton()

	remoteControl.SetCommand(turnOffCommand)
	remoteControl.PressButton()
}
