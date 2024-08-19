// automatically generated by stateify.

package pipe

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (e *Endpoint) StateTypeName() string {
	return "pkg/tcpip/link/pipe.Endpoint"
}

func (e *Endpoint) StateFields() []string {
	return []string{
		"linked",
		"dispatcher",
		"linkAddr",
		"mtu",
	}
}

func (e *Endpoint) beforeSave() {}

// +checklocksignore
func (e *Endpoint) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.linked)
	stateSinkObject.Save(1, &e.dispatcher)
	stateSinkObject.Save(2, &e.linkAddr)
	stateSinkObject.Save(3, &e.mtu)
}

func (e *Endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (e *Endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.linked)
	stateSourceObject.Load(1, &e.dispatcher)
	stateSourceObject.Load(2, &e.linkAddr)
	stateSourceObject.Load(3, &e.mtu)
}

func init() {
	state.Register((*Endpoint)(nil))
}