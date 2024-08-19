// automatically generated by stateify.

//go:build linux
// +build linux

package xdp

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (ep *endpoint) StateTypeName() string {
	return "pkg/tcpip/link/xdp.endpoint"
}

func (ep *endpoint) StateFields() []string {
	return []string{
		"fd",
		"caps",
		"networkDispatcher",
		"control",
		"stopFD",
		"addr",
	}
}

func (ep *endpoint) beforeSave() {}

// +checklocksignore
func (ep *endpoint) StateSave(stateSinkObject state.Sink) {
	ep.beforeSave()
	stateSinkObject.Save(0, &ep.fd)
	stateSinkObject.Save(1, &ep.caps)
	stateSinkObject.Save(2, &ep.networkDispatcher)
	stateSinkObject.Save(3, &ep.control)
	stateSinkObject.Save(4, &ep.stopFD)
	stateSinkObject.Save(5, &ep.addr)
}

func (ep *endpoint) afterLoad(context.Context) {}

// +checklocksignore
func (ep *endpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ep.fd)
	stateSourceObject.Load(1, &ep.caps)
	stateSourceObject.Load(2, &ep.networkDispatcher)
	stateSourceObject.Load(3, &ep.control)
	stateSourceObject.Load(4, &ep.stopFD)
	stateSourceObject.Load(5, &ep.addr)
}

func init() {
	state.Register((*endpoint)(nil))
}