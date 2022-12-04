package state_dp

import "log"

type stateCode string

const (
	unMember = "普通用户状态"
	seMember = "标准版会员状态"
	peMember = "高级版会员状态"
)

type MemberState interface {
	Play(ctx Context)
	getStateCode() stateCode
}

type SEMemberState struct {
	stateCode
}

func NewSEMemberState() *SEMemberState {
	return &SEMemberState{stateCode: seMember}

}

func (receiver *SEMemberState) getStateCode() stateCode {
	return receiver.stateCode
}

func (receiver *SEMemberState) Play(ctx Context) {
	getCurrentState(ctx)
	log.Default().Println("播放720p资源")
}

type UnMemberState struct {
	stateCode
}

func NewUnMemberState() *UnMemberState {
	return &UnMemberState{stateCode: unMember}
}

func (receiver *UnMemberState) Play(ctx Context) {
	getCurrentState(ctx)
	log.Default().Println("无法播放资源")
}

func (receiver *UnMemberState) getStateCode() stateCode {
	return receiver.stateCode
}

type PeMemberState struct {
	stateCode
}

func NewPeMemberState() *PeMemberState {
	return &PeMemberState{stateCode: peMember}
}

func (PeMemberState) Play(ctx Context) {
	getCurrentState(ctx)
	log.Default().Println("播放1080p资源")
}

func (receiver *PeMemberState) getStateCode() stateCode {
	return receiver.stateCode
}

func getCurrentState(ctx Context) {
	log.Default().Println("当前状态是：", ctx.state.getStateCode())
}
