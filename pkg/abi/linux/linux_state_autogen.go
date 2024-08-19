// automatically generated by stateify.

package linux

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (i *IOEvent) StateTypeName() string {
	return "pkg/abi/linux.IOEvent"
}

func (i *IOEvent) StateFields() []string {
	return []string{
		"Data",
		"Obj",
		"Result",
		"Result2",
	}
}

func (i *IOEvent) beforeSave() {}

// +checklocksignore
func (i *IOEvent) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Data)
	stateSinkObject.Save(1, &i.Obj)
	stateSinkObject.Save(2, &i.Result)
	stateSinkObject.Save(3, &i.Result2)
}

func (i *IOEvent) afterLoad(context.Context) {}

// +checklocksignore
func (i *IOEvent) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Data)
	stateSourceObject.Load(1, &i.Obj)
	stateSourceObject.Load(2, &i.Result)
	stateSourceObject.Load(3, &i.Result2)
}

func (b *BPFInstruction) StateTypeName() string {
	return "pkg/abi/linux.BPFInstruction"
}

func (b *BPFInstruction) StateFields() []string {
	return []string{
		"OpCode",
		"JumpIfTrue",
		"JumpIfFalse",
		"K",
	}
}

func (b *BPFInstruction) beforeSave() {}

// +checklocksignore
func (b *BPFInstruction) StateSave(stateSinkObject state.Sink) {
	b.beforeSave()
	stateSinkObject.Save(0, &b.OpCode)
	stateSinkObject.Save(1, &b.JumpIfTrue)
	stateSinkObject.Save(2, &b.JumpIfFalse)
	stateSinkObject.Save(3, &b.K)
}

func (b *BPFInstruction) afterLoad(context.Context) {}

// +checklocksignore
func (b *BPFInstruction) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &b.OpCode)
	stateSourceObject.Load(1, &b.JumpIfTrue)
	stateSourceObject.Load(2, &b.JumpIfFalse)
	stateSourceObject.Load(3, &b.K)
}

func (f *FUSEHeaderIn) StateTypeName() string {
	return "pkg/abi/linux.FUSEHeaderIn"
}

func (f *FUSEHeaderIn) StateFields() []string {
	return []string{
		"Len",
		"Opcode",
		"Unique",
		"NodeID",
		"UID",
		"GID",
		"PID",
	}
}

func (f *FUSEHeaderIn) beforeSave() {}

// +checklocksignore
func (f *FUSEHeaderIn) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Len)
	stateSinkObject.Save(1, &f.Opcode)
	stateSinkObject.Save(2, &f.Unique)
	stateSinkObject.Save(3, &f.NodeID)
	stateSinkObject.Save(4, &f.UID)
	stateSinkObject.Save(5, &f.GID)
	stateSinkObject.Save(6, &f.PID)
}

func (f *FUSEHeaderIn) afterLoad(context.Context) {}

// +checklocksignore
func (f *FUSEHeaderIn) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Len)
	stateSourceObject.Load(1, &f.Opcode)
	stateSourceObject.Load(2, &f.Unique)
	stateSourceObject.Load(3, &f.NodeID)
	stateSourceObject.Load(4, &f.UID)
	stateSourceObject.Load(5, &f.GID)
	stateSourceObject.Load(6, &f.PID)
}

func (f *FUSEHeaderOut) StateTypeName() string {
	return "pkg/abi/linux.FUSEHeaderOut"
}

func (f *FUSEHeaderOut) StateFields() []string {
	return []string{
		"Len",
		"Error",
		"Unique",
	}
}

func (f *FUSEHeaderOut) beforeSave() {}

// +checklocksignore
func (f *FUSEHeaderOut) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.Len)
	stateSinkObject.Save(1, &f.Error)
	stateSinkObject.Save(2, &f.Unique)
}

func (f *FUSEHeaderOut) afterLoad(context.Context) {}

// +checklocksignore
func (f *FUSEHeaderOut) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.Len)
	stateSourceObject.Load(1, &f.Error)
	stateSourceObject.Load(2, &f.Unique)
}

func (i *IOUringCqe) StateTypeName() string {
	return "pkg/abi/linux.IOUringCqe"
}

func (i *IOUringCqe) StateFields() []string {
	return []string{
		"UserData",
		"Res",
		"Flags",
	}
}

func (i *IOUringCqe) beforeSave() {}

// +checklocksignore
func (i *IOUringCqe) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.UserData)
	stateSinkObject.Save(1, &i.Res)
	stateSinkObject.Save(2, &i.Flags)
}

func (i *IOUringCqe) afterLoad(context.Context) {}

// +checklocksignore
func (i *IOUringCqe) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.UserData)
	stateSourceObject.Load(1, &i.Res)
	stateSourceObject.Load(2, &i.Flags)
}

func (i *IOUring) StateTypeName() string {
	return "pkg/abi/linux.IOUring"
}

func (i *IOUring) StateFields() []string {
	return []string{
		"Head",
		"Tail",
	}
}

func (i *IOUring) beforeSave() {}

// +checklocksignore
func (i *IOUring) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Head)
	stateSinkObject.Save(1, &i.Tail)
}

func (i *IOUring) afterLoad(context.Context) {}

// +checklocksignore
func (i *IOUring) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Head)
	stateSourceObject.Load(1, &i.Tail)
}

func (i *IORings) StateTypeName() string {
	return "pkg/abi/linux.IORings"
}

func (i *IORings) StateFields() []string {
	return []string{
		"Sq",
		"Cq",
		"SqRingMask",
		"CqRingMask",
		"SqRingEntries",
		"CqRingEntries",
		"sqDropped",
		"sqFlags",
		"cqFlags",
		"CqOverflow",
	}
}

func (i *IORings) beforeSave() {}

// +checklocksignore
func (i *IORings) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Sq)
	stateSinkObject.Save(1, &i.Cq)
	stateSinkObject.Save(2, &i.SqRingMask)
	stateSinkObject.Save(3, &i.CqRingMask)
	stateSinkObject.Save(4, &i.SqRingEntries)
	stateSinkObject.Save(5, &i.CqRingEntries)
	stateSinkObject.Save(6, &i.sqDropped)
	stateSinkObject.Save(7, &i.sqFlags)
	stateSinkObject.Save(8, &i.cqFlags)
	stateSinkObject.Save(9, &i.CqOverflow)
}

func (i *IORings) afterLoad(context.Context) {}

// +checklocksignore
func (i *IORings) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Sq)
	stateSourceObject.Load(1, &i.Cq)
	stateSourceObject.Load(2, &i.SqRingMask)
	stateSourceObject.Load(3, &i.CqRingMask)
	stateSourceObject.Load(4, &i.SqRingEntries)
	stateSourceObject.Load(5, &i.CqRingEntries)
	stateSourceObject.Load(6, &i.sqDropped)
	stateSourceObject.Load(7, &i.sqFlags)
	stateSourceObject.Load(8, &i.cqFlags)
	stateSourceObject.Load(9, &i.CqOverflow)
}

func (i *IOUringSqe) StateTypeName() string {
	return "pkg/abi/linux.IOUringSqe"
}

func (i *IOUringSqe) StateFields() []string {
	return []string{
		"Opcode",
		"Flags",
		"IoPrio",
		"Fd",
		"OffOrAddrOrCmdOp",
		"AddrOrSpliceOff",
		"Len",
		"specialFlags",
		"UserData",
		"BufIndexOrGroup",
		"personality",
		"spliceFDOrFileIndex",
		"addr3",
	}
}

func (i *IOUringSqe) beforeSave() {}

// +checklocksignore
func (i *IOUringSqe) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Opcode)
	stateSinkObject.Save(1, &i.Flags)
	stateSinkObject.Save(2, &i.IoPrio)
	stateSinkObject.Save(3, &i.Fd)
	stateSinkObject.Save(4, &i.OffOrAddrOrCmdOp)
	stateSinkObject.Save(5, &i.AddrOrSpliceOff)
	stateSinkObject.Save(6, &i.Len)
	stateSinkObject.Save(7, &i.specialFlags)
	stateSinkObject.Save(8, &i.UserData)
	stateSinkObject.Save(9, &i.BufIndexOrGroup)
	stateSinkObject.Save(10, &i.personality)
	stateSinkObject.Save(11, &i.spliceFDOrFileIndex)
	stateSinkObject.Save(12, &i.addr3)
}

func (i *IOUringSqe) afterLoad(context.Context) {}

// +checklocksignore
func (i *IOUringSqe) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Opcode)
	stateSourceObject.Load(1, &i.Flags)
	stateSourceObject.Load(2, &i.IoPrio)
	stateSourceObject.Load(3, &i.Fd)
	stateSourceObject.Load(4, &i.OffOrAddrOrCmdOp)
	stateSourceObject.Load(5, &i.AddrOrSpliceOff)
	stateSourceObject.Load(6, &i.Len)
	stateSourceObject.Load(7, &i.specialFlags)
	stateSourceObject.Load(8, &i.UserData)
	stateSourceObject.Load(9, &i.BufIndexOrGroup)
	stateSourceObject.Load(10, &i.personality)
	stateSourceObject.Load(11, &i.spliceFDOrFileIndex)
	stateSourceObject.Load(12, &i.addr3)
}

func (s *SigAction) StateTypeName() string {
	return "pkg/abi/linux.SigAction"
}

func (s *SigAction) StateFields() []string {
	return []string{
		"Handler",
		"Flags",
		"Restorer",
		"Mask",
	}
}

func (s *SigAction) beforeSave() {}

// +checklocksignore
func (s *SigAction) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Handler)
	stateSinkObject.Save(1, &s.Flags)
	stateSinkObject.Save(2, &s.Restorer)
	stateSinkObject.Save(3, &s.Mask)
}

func (s *SigAction) afterLoad(context.Context) {}

// +checklocksignore
func (s *SigAction) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Handler)
	stateSourceObject.Load(1, &s.Flags)
	stateSourceObject.Load(2, &s.Restorer)
	stateSourceObject.Load(3, &s.Mask)
}

func (s *SignalStack) StateTypeName() string {
	return "pkg/abi/linux.SignalStack"
}

func (s *SignalStack) StateFields() []string {
	return []string{
		"Addr",
		"Flags",
		"Size",
	}
}

func (s *SignalStack) beforeSave() {}

// +checklocksignore
func (s *SignalStack) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Addr)
	stateSinkObject.Save(1, &s.Flags)
	stateSinkObject.Save(2, &s.Size)
}

func (s *SignalStack) afterLoad(context.Context) {}

// +checklocksignore
func (s *SignalStack) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Addr)
	stateSourceObject.Load(1, &s.Flags)
	stateSourceObject.Load(2, &s.Size)
}

func (s *SignalInfo) StateTypeName() string {
	return "pkg/abi/linux.SignalInfo"
}

func (s *SignalInfo) StateFields() []string {
	return []string{
		"Signo",
		"Errno",
		"Code",
		"Fields",
	}
}

func (s *SignalInfo) beforeSave() {}

// +checklocksignore
func (s *SignalInfo) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Signo)
	stateSinkObject.Save(1, &s.Errno)
	stateSinkObject.Save(2, &s.Code)
	stateSinkObject.Save(3, &s.Fields)
}

func (s *SignalInfo) afterLoad(context.Context) {}

// +checklocksignore
func (s *SignalInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Signo)
	stateSourceObject.Load(1, &s.Errno)
	stateSourceObject.Load(2, &s.Code)
	stateSourceObject.Load(3, &s.Fields)
}

func (c *ControlMessageIPPacketInfo) StateTypeName() string {
	return "pkg/abi/linux.ControlMessageIPPacketInfo"
}

func (c *ControlMessageIPPacketInfo) StateFields() []string {
	return []string{
		"NIC",
		"LocalAddr",
		"DestinationAddr",
	}
}

func (c *ControlMessageIPPacketInfo) beforeSave() {}

// +checklocksignore
func (c *ControlMessageIPPacketInfo) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.NIC)
	stateSinkObject.Save(1, &c.LocalAddr)
	stateSinkObject.Save(2, &c.DestinationAddr)
}

func (c *ControlMessageIPPacketInfo) afterLoad(context.Context) {}

// +checklocksignore
func (c *ControlMessageIPPacketInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.NIC)
	stateSourceObject.Load(1, &c.LocalAddr)
	stateSourceObject.Load(2, &c.DestinationAddr)
}

func (c *ControlMessageIPv6PacketInfo) StateTypeName() string {
	return "pkg/abi/linux.ControlMessageIPv6PacketInfo"
}

func (c *ControlMessageIPv6PacketInfo) StateFields() []string {
	return []string{
		"Addr",
		"NIC",
	}
}

func (c *ControlMessageIPv6PacketInfo) beforeSave() {}

// +checklocksignore
func (c *ControlMessageIPv6PacketInfo) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.Addr)
	stateSinkObject.Save(1, &c.NIC)
}

func (c *ControlMessageIPv6PacketInfo) afterLoad(context.Context) {}

// +checklocksignore
func (c *ControlMessageIPv6PacketInfo) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.Addr)
	stateSourceObject.Load(1, &c.NIC)
}

func (i *ICMP6Filter) StateTypeName() string {
	return "pkg/abi/linux.ICMP6Filter"
}

func (i *ICMP6Filter) StateFields() []string {
	return []string{
		"Filter",
	}
}

func (i *ICMP6Filter) beforeSave() {}

// +checklocksignore
func (i *ICMP6Filter) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.Filter)
}

func (i *ICMP6Filter) afterLoad(context.Context) {}

// +checklocksignore
func (i *ICMP6Filter) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.Filter)
}

func (t *KernelTermios) StateTypeName() string {
	return "pkg/abi/linux.KernelTermios"
}

func (t *KernelTermios) StateFields() []string {
	return []string{
		"InputFlags",
		"OutputFlags",
		"ControlFlags",
		"LocalFlags",
		"LineDiscipline",
		"ControlCharacters",
		"InputSpeed",
		"OutputSpeed",
	}
}

func (t *KernelTermios) beforeSave() {}

// +checklocksignore
func (t *KernelTermios) StateSave(stateSinkObject state.Sink) {
	t.beforeSave()
	stateSinkObject.Save(0, &t.InputFlags)
	stateSinkObject.Save(1, &t.OutputFlags)
	stateSinkObject.Save(2, &t.ControlFlags)
	stateSinkObject.Save(3, &t.LocalFlags)
	stateSinkObject.Save(4, &t.LineDiscipline)
	stateSinkObject.Save(5, &t.ControlCharacters)
	stateSinkObject.Save(6, &t.InputSpeed)
	stateSinkObject.Save(7, &t.OutputSpeed)
}

func (t *KernelTermios) afterLoad(context.Context) {}

// +checklocksignore
func (t *KernelTermios) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &t.InputFlags)
	stateSourceObject.Load(1, &t.OutputFlags)
	stateSourceObject.Load(2, &t.ControlFlags)
	stateSourceObject.Load(3, &t.LocalFlags)
	stateSourceObject.Load(4, &t.LineDiscipline)
	stateSourceObject.Load(5, &t.ControlCharacters)
	stateSourceObject.Load(6, &t.InputSpeed)
	stateSourceObject.Load(7, &t.OutputSpeed)
}

func (w *WindowSize) StateTypeName() string {
	return "pkg/abi/linux.WindowSize"
}

func (w *WindowSize) StateFields() []string {
	return []string{
		"Rows",
		"Cols",
	}
}

func (w *WindowSize) beforeSave() {}

// +checklocksignore
func (w *WindowSize) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.Rows)
	stateSinkObject.Save(1, &w.Cols)
}

func (w *WindowSize) afterLoad(context.Context) {}

// +checklocksignore
func (w *WindowSize) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.Rows)
	stateSourceObject.Load(1, &w.Cols)
}

func init() {
	state.Register((*IOEvent)(nil))
	state.Register((*BPFInstruction)(nil))
	state.Register((*FUSEHeaderIn)(nil))
	state.Register((*FUSEHeaderOut)(nil))
	state.Register((*IOUringCqe)(nil))
	state.Register((*IOUring)(nil))
	state.Register((*IORings)(nil))
	state.Register((*IOUringSqe)(nil))
	state.Register((*SigAction)(nil))
	state.Register((*SignalStack)(nil))
	state.Register((*SignalInfo)(nil))
	state.Register((*ControlMessageIPPacketInfo)(nil))
	state.Register((*ControlMessageIPv6PacketInfo)(nil))
	state.Register((*ICMP6Filter)(nil))
	state.Register((*KernelTermios)(nil))
	state.Register((*WindowSize)(nil))
}