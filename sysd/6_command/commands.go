package command

type Command interface {
	Execute()
}

// Light
type Light struct {
	On bool
}

func (l *Light) turn(v bool) {
	l.On = v
}

type LightOnCommand struct {
	Light *Light
}

func NewLightOnCommand(l *Light) *LightOnCommand {
	return &LightOnCommand{
		Light: l,
	}
}

func (lc *LightOnCommand) Execute() {
	lc.Light.turn(true)
}

// GarageDoor
type GarageDoor struct {
	OpenState  bool
	LightState bool
}

func (gd *GarageDoor) Open() {
	gd.OpenState = true
}

func (gd *GarageDoor) Close() {
	gd.OpenState = false
}

type GarageDoorCommand struct {
	GarageDoor *GarageDoor
}

func NewGarageDoorCommand(gd *GarageDoor) *GarageDoorCommand {
	return &GarageDoorCommand{GarageDoor: gd}
}

// GarageDoor commands
type GarageDoorOpenCommand struct {
	*GarageDoorCommand
}

func NewGarageDoorOpenCommand(gd *GarageDoor) *GarageDoorOpenCommand {
	return &GarageDoorOpenCommand{&GarageDoorCommand{GarageDoor: gd}}
}

func (cmd *GarageDoorOpenCommand) Execute() {
	cmd.GarageDoor.Open()
}

type GarageDoorCloseCommand struct {
	*GarageDoorCommand
}

func NewGarageDoorCloseCommand(gd *GarageDoor) *GarageDoorCloseCommand {
	return &GarageDoorCloseCommand{&GarageDoorCommand{GarageDoor: gd}}
}

func (cmd *GarageDoorCloseCommand) Execute() {
	cmd.GarageDoor.Close()
}

type GarageDoorLightOnCommand struct {
	*LightOnCommand
}

func NewGarageDoorLightOnCommand(l *Light) *GarageDoorLightOnCommand {
	lightOnCmd := NewLightOnCommand(l)
	return &GarageDoorLightOnCommand{LightOnCommand: lightOnCmd}
}

func (cmd *GarageDoorLightOnCommand) Execute() {
	cmd.LightOnCommand.Execute()
}
