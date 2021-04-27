package main

import "fmt"

func main() {
	sctx:=NewStateContext()
	startCtx:=StartState{}
	startCtx.doAction(sctx)

	fmt.Println(startCtx.String())

	stoptCtx:=StopState{}
	stoptCtx.doAction(sctx)


	fmt.Println(stoptCtx.String())

}