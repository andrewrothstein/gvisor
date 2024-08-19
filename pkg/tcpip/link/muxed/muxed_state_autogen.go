// automatically generated by stateify.

package muxed

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (m *InjectableEndpoint) StateTypeName() string {
	return "pkg/tcpip/link/muxed.InjectableEndpoint"
}

func (m *InjectableEndpoint) StateFields() []string {
	return []string{
		"routes",
		"dispatcher",
	}
}

func (m *InjectableEndpoint) beforeSave() {}

// +checklocksignore
func (m *InjectableEndpoint) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.routes)
	stateSinkObject.Save(1, &m.dispatcher)
}

func (m *InjectableEndpoint) afterLoad(context.Context) {}

// +checklocksignore
func (m *InjectableEndpoint) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.routes)
	stateSourceObject.Load(1, &m.dispatcher)
}

func init() {
	state.Register((*InjectableEndpoint)(nil))
}