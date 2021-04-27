package main

import "fmt"

type StateContext struct {
	state State
}

func NewStateContext() StateContext {
	return StateContext{}
}
func (sct *StateContext) setState(state State)  {
	sct.state=state
}
func (sct *StateContext) getState() State  {
	return sct.state
}

type State interface {
	doAction(ctx StateContext)
}

type StartState struct {
	State
	StateContext
}

func (s *StartState) doAction(ctx StateContext)  {
	fmt.Println("Player is in Start State")
	ctx.setState(s)
	//s.setState(s)

}
func (s StartState) String() string  {
	return "Start State"
}

type StopState struct {
	State
	StateContext
}

func (s *StopState) doAction(ctx StateContext)  {
	fmt.Println("Player is in Stop State")
	ctx.setState(s)
	//s.setState(s)

}
func (s StopState) String() string  {
	return "Stop State"
}