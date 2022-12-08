package test

import (
	"lbf.com/go-dp/command_dp"
	"testing"
)

func TestCommandDp(t *testing.T) {
	singSongCommand := command_dp.NewSingSongCommand(&command_dp.YangkunReceiver{})
	danceCommand := command_dp.NewDanceCommand(&command_dp.YangmiReceiver{})
	invoker := command_dp.NewInvoker(singSongCommand, danceCommand)
	invoker.ExeCommand()
}
