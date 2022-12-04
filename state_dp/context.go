package state_dp

type Context struct {
	state MemberState
}

func NewContext(state MemberState) *Context {
	return &Context{state: state}
}

func (ctx Context) Request() {
	ctx.state.Play(ctx)
}
