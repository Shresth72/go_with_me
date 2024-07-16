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

/*
         __                                     __
  ____  / /_     ____ ___  __  __   ____  _____/ /_
 / __ \/ __ \   / __ `__ \/ / / /  /_  / / ___/ __ \
/ /_/ / / / /  / / / / / / /_/ /    / /_(__  ) / / /
\____/_/ /_/  /_/ /_/ /_/\__, /    /___/____/_/ /_/
                        /____/
*/

