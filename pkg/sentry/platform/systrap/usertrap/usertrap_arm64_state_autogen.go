// automatically generated by stateify.

//go:build arm64
// +build arm64

package usertrap

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (s *State) StateTypeName() string {
	return "pkg/sentry/platform/systrap/usertrap.State"
}

func (s *State) StateFields() []string {
	return []string{}
}

func (s *State) beforeSave() {}

// +checklocksignore
func (s *State) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
}

func (s *State) afterLoad(context.Context) {}

// +checklocksignore
func (s *State) StateLoad(ctx context.Context, stateSourceObject state.Source) {
}

func init() {
	state.Register((*State)(nil))
}