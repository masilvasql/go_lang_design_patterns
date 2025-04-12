package invoker

import "github.com/masilvasql/command/command"

type RemoteControl struct {
	buttons map[string]command.Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		buttons: make(map[string]command.Command),
	}
}

func (r *RemoteControl) SetCommand(slot string, command command.Command) {
	r.buttons[slot] = command
}

func (r *RemoteControl) PressButton(slot string) {
	if cmd, ok := r.buttons[slot]; ok {
		cmd.Execute()
	}
}
