package command_dp

import "log"

type Command interface {
	Execute()
}

type YangkunReceiver struct {
}

func (receiver YangkunReceiver) SingSong() {
	log.Default().Println("杨坤唱歌")
}

type SingSongCommand struct {
	receiver *YangkunReceiver
}

func NewSingSongCommand(receiver *YangkunReceiver) *SingSongCommand {
	return &SingSongCommand{receiver: receiver}
}

func (s *SingSongCommand) Execute() {
	s.receiver.SingSong()
}

type YangmiReceiver struct {
}

func (receiver YangmiReceiver) Dance() {
	log.Default().Println("杨幂跳舞")
}

type DanceCommand struct {
	receiver *YangmiReceiver
}

func NewDanceCommand(receiver *YangmiReceiver) *DanceCommand {
	return &DanceCommand{receiver: receiver}
}

func (d *DanceCommand) Execute() {
	d.receiver.Dance()
}

type commandList []Command
type Invoker struct {
	commandList
}

func NewInvoker(command ...Command) *Invoker {
	return &Invoker{commandList: command}
}

func (receiver *Invoker) AddCommand(command ...Command) {
	receiver.commandList = append(receiver.commandList, command...)
}

func (receiver *Invoker) ExeCommand() {
	for _, command := range receiver.commandList {
		command.Execute()
	}
}
