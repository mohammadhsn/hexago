package hexago

// CmdBus decouples Cmd from its execution, dependencies, side effects.
type CmdBus interface {
	Handle(Cmd) string
}
