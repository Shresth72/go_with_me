package command

type SimpleRemoteControl struct {
	slot Command
}

func (s *SimpleRemoteControl) SetCommand(command Command) {
	s.slot = command
}

func (s *SimpleRemoteControl) ButtonWasPressed() {
	s.slot.Execute()
}
