package engine

//Command performs actions in a single event loop iteration
type Command interface {
	Execute(handler Handler)
}

//Handler for sending commands to event loop
type Handler interface {
	Post(cmd Command)
}
