package hexago

// Cmd is just Command's abbreviation, in the clean architecture there's a layer
// which is responsible for handling system use cases. It is common to call this
// layer `application` layer. It acts like an orchestrator which connects
// different layers. A common pattern for implementing the application layer is,
// Command Query Responsibility Segregation or CQRS. The Cmd struct, tries to act
// like a command.
type Cmd struct {
	Name    string
	Payload interface{}
}

// CmdHandler is considered for being responsible to handling, one single Cmd
// [command]. It's common to use the domain repositories in order to ship
// aggregates, and using the domain functionalities to handling use cases.
type CmdHandler interface {
	Handle(Cmd) (string, error)
}
