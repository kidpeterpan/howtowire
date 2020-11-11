//+build wireinject

package withwire

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
