// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package withwire

// Injectors from wire.go:

func InitializeEvent() Event {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}
