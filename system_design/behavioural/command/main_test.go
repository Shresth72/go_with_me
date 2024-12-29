package command_test

import (
	"testing"

	"github.com/Shresth72/sysd/sysd/6_command"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommandTestSuite struct {
	suite.Suite
	remote   *command.SimpleRemoteControl
	commands map[string]command.Command
}

func (suite *CommandTestSuite) SetupTest() {
	suite.remote = &command.SimpleRemoteControl{}
	suite.commands = make(map[string]command.Command)

	// Light setup
	light := &command.Light{On: false}
	lightOn := command.NewLightOnCommand(light)
	suite.commands["LightOn"] = lightOn

	// GarageDoor setup
	garageDoor := &command.GarageDoor{OpenState: false, LightState: false}
	garageDoorOpen := command.NewGarageDoorOpenCommand(garageDoor)
	garageDoorClose := command.NewGarageDoorCloseCommand(garageDoor)
	garageDoorLightOn := command.NewGarageDoorLightOnCommand(light)

	suite.commands["GarageDoorOpen"] = garageDoorOpen
	suite.commands["GarageDoorClose"] = garageDoorClose
	suite.commands["GarageDoorLightOn"] = garageDoorLightOn
}

func TestCommandTestSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}

func (suite *CommandTestSuite) TestLightCommand() {
	// Test LightOn command
	suite.remote.SetCommand(suite.commands["LightOn"])
	suite.remote.ButtonWasPressed()

	lightOnCommand := suite.commands["LightOn"].(*command.LightOnCommand)
	assert.True(suite.T(), lightOnCommand.Light.On)
}

func (suite *CommandTestSuite) TestGarageDoorCommand() {
	// Test GarageDoorOpen command
	suite.remote.SetCommand(suite.commands["GarageDoorOpen"])
	suite.remote.ButtonWasPressed()

	garageDoorOpenCommand := suite.commands["GarageDoorOpen"].(*command.GarageDoorOpenCommand)
	assert.True(suite.T(), garageDoorOpenCommand.GarageDoor.OpenState)

	// Test GarageDoorClose command
	suite.remote.SetCommand(suite.commands["GarageDoorClose"])
	suite.remote.ButtonWasPressed()

	garageDoorCloseCommand := suite.commands["GarageDoorClose"].(*command.GarageDoorCloseCommand)
	assert.False(suite.T(), garageDoorCloseCommand.GarageDoor.OpenState)

	// Test GarageDoorLightOn command (through LightOnCommand)
	suite.remote.SetCommand(suite.commands["GarageDoorLightOn"])
	suite.remote.ButtonWasPressed()

	garageDoorLightOnCommand := suite.commands["GarageDoorLightOn"].(*command.GarageDoorLightOnCommand)
	assert.True(suite.T(), garageDoorLightOnCommand.LightOnCommand.Light.On)
}

