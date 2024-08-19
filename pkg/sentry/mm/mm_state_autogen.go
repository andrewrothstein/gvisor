// automatically generated by stateify.

package mm

import (
	"context"

	"gvisor.dev/gvisor/pkg/state"
)

func (a *aioManager) StateTypeName() string {
	return "pkg/sentry/mm.aioManager"
}

func (a *aioManager) StateFields() []string {
	return []string{
		"contexts",
	}
}

func (a *aioManager) beforeSave() {}

// +checklocksignore
func (a *aioManager) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.contexts)
}

func (a *aioManager) afterLoad(context.Context) {}

// +checklocksignore
func (a *aioManager) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.contexts)
}

func (i *ioResult) StateTypeName() string {
	return "pkg/sentry/mm.ioResult"
}

func (i *ioResult) StateFields() []string {
	return []string{
		"data",
		"ioEntry",
	}
}

func (i *ioResult) beforeSave() {}

// +checklocksignore
func (i *ioResult) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.data)
	stateSinkObject.Save(1, &i.ioEntry)
}

func (i *ioResult) afterLoad(context.Context) {}

// +checklocksignore
func (i *ioResult) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.data)
	stateSourceObject.Load(1, &i.ioEntry)
}

func (aio *AIOContext) StateTypeName() string {
	return "pkg/sentry/mm.AIOContext"
}

func (aio *AIOContext) StateFields() []string {
	return []string{
		"results",
		"maxOutstanding",
		"outstanding",
	}
}

func (aio *AIOContext) beforeSave() {}

// +checklocksignore
func (aio *AIOContext) StateSave(stateSinkObject state.Sink) {
	aio.beforeSave()
	if !state.IsZeroValue(&aio.dead) {
		state.Failf("dead is %#v, expected zero", &aio.dead)
	}
	stateSinkObject.Save(0, &aio.results)
	stateSinkObject.Save(1, &aio.maxOutstanding)
	stateSinkObject.Save(2, &aio.outstanding)
}

// +checklocksignore
func (aio *AIOContext) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &aio.results)
	stateSourceObject.Load(1, &aio.maxOutstanding)
	stateSourceObject.Load(2, &aio.outstanding)
	stateSourceObject.AfterLoad(func() { aio.afterLoad(ctx) })
}

func (m *aioMappable) StateTypeName() string {
	return "pkg/sentry/mm.aioMappable"
}

func (m *aioMappable) StateFields() []string {
	return []string{
		"aioMappableRefs",
		"fr",
	}
}

func (m *aioMappable) beforeSave() {}

// +checklocksignore
func (m *aioMappable) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.aioMappableRefs)
	stateSinkObject.Save(1, &m.fr)
}

// +checklocksignore
func (m *aioMappable) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.aioMappableRefs)
	stateSourceObject.Load(1, &m.fr)
	stateSourceObject.AfterLoad(func() { m.afterLoad(ctx) })
}

func (r *aioMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.aioMappableRefs"
}

func (r *aioMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *aioMappableRefs) beforeSave() {}

// +checklocksignore
func (r *aioMappableRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *aioMappableRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (l *ioList) StateTypeName() string {
	return "pkg/sentry/mm.ioList"
}

func (l *ioList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *ioList) beforeSave() {}

// +checklocksignore
func (l *ioList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *ioList) afterLoad(context.Context) {}

// +checklocksignore
func (l *ioList) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *ioEntry) StateTypeName() string {
	return "pkg/sentry/mm.ioEntry"
}

func (e *ioEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *ioEntry) beforeSave() {}

// +checklocksignore
func (e *ioEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *ioEntry) afterLoad(context.Context) {}

// +checklocksignore
func (e *ioEntry) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (mm *MemoryManager) StateTypeName() string {
	return "pkg/sentry/mm.MemoryManager"
}

func (mm *MemoryManager) StateFields() []string {
	return []string{
		"p",
		"layout",
		"users",
		"vmas",
		"brk",
		"usageAS",
		"lockedAS",
		"dataAS",
		"defMLockMode",
		"pmas",
		"curRSS",
		"maxRSS",
		"dumpability",
		"argv",
		"envv",
		"auxv",
		"executable",
		"aioManager",
		"sleepForActivation",
		"vdsoSigReturnAddr",
		"membarrierPrivateEnabled",
		"membarrierRSeqEnabled",
	}
}

func (mm *MemoryManager) beforeSave() {}

// +checklocksignore
func (mm *MemoryManager) StateSave(stateSinkObject state.Sink) {
	mm.beforeSave()
	if !state.IsZeroValue(&mm.active) {
		state.Failf("active is %#v, expected zero", &mm.active)
	}
	if !state.IsZeroValue(&mm.captureInvalidations) {
		state.Failf("captureInvalidations is %#v, expected zero", &mm.captureInvalidations)
	}
	stateSinkObject.Save(0, &mm.p)
	stateSinkObject.Save(1, &mm.layout)
	stateSinkObject.Save(2, &mm.users)
	stateSinkObject.Save(3, &mm.vmas)
	stateSinkObject.Save(4, &mm.brk)
	stateSinkObject.Save(5, &mm.usageAS)
	stateSinkObject.Save(6, &mm.lockedAS)
	stateSinkObject.Save(7, &mm.dataAS)
	stateSinkObject.Save(8, &mm.defMLockMode)
	stateSinkObject.Save(9, &mm.pmas)
	stateSinkObject.Save(10, &mm.curRSS)
	stateSinkObject.Save(11, &mm.maxRSS)
	stateSinkObject.Save(12, &mm.dumpability)
	stateSinkObject.Save(13, &mm.argv)
	stateSinkObject.Save(14, &mm.envv)
	stateSinkObject.Save(15, &mm.auxv)
	stateSinkObject.Save(16, &mm.executable)
	stateSinkObject.Save(17, &mm.aioManager)
	stateSinkObject.Save(18, &mm.sleepForActivation)
	stateSinkObject.Save(19, &mm.vdsoSigReturnAddr)
	stateSinkObject.Save(20, &mm.membarrierPrivateEnabled)
	stateSinkObject.Save(21, &mm.membarrierRSeqEnabled)
}

// +checklocksignore
func (mm *MemoryManager) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &mm.p)
	stateSourceObject.Load(1, &mm.layout)
	stateSourceObject.Load(2, &mm.users)
	stateSourceObject.Load(3, &mm.vmas)
	stateSourceObject.Load(4, &mm.brk)
	stateSourceObject.Load(5, &mm.usageAS)
	stateSourceObject.Load(6, &mm.lockedAS)
	stateSourceObject.Load(7, &mm.dataAS)
	stateSourceObject.Load(8, &mm.defMLockMode)
	stateSourceObject.Load(9, &mm.pmas)
	stateSourceObject.Load(10, &mm.curRSS)
	stateSourceObject.Load(11, &mm.maxRSS)
	stateSourceObject.Load(12, &mm.dumpability)
	stateSourceObject.Load(13, &mm.argv)
	stateSourceObject.Load(14, &mm.envv)
	stateSourceObject.Load(15, &mm.auxv)
	stateSourceObject.Load(16, &mm.executable)
	stateSourceObject.Load(17, &mm.aioManager)
	stateSourceObject.Load(18, &mm.sleepForActivation)
	stateSourceObject.Load(19, &mm.vdsoSigReturnAddr)
	stateSourceObject.Load(20, &mm.membarrierPrivateEnabled)
	stateSourceObject.Load(21, &mm.membarrierRSeqEnabled)
	stateSourceObject.AfterLoad(func() { mm.afterLoad(ctx) })
}

func (v *vma) StateTypeName() string {
	return "pkg/sentry/mm.vma"
}

func (v *vma) StateFields() []string {
	return []string{
		"mappable",
		"off",
		"realPerms",
		"dontfork",
		"mlockMode",
		"numaPolicy",
		"numaNodemask",
		"id",
		"hint",
		"lastFault",
	}
}

func (v *vma) beforeSave() {}

// +checklocksignore
func (v *vma) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	var realPermsValue int
	realPermsValue = v.saveRealPerms()
	stateSinkObject.SaveValue(2, realPermsValue)
	stateSinkObject.Save(0, &v.mappable)
	stateSinkObject.Save(1, &v.off)
	stateSinkObject.Save(3, &v.dontfork)
	stateSinkObject.Save(4, &v.mlockMode)
	stateSinkObject.Save(5, &v.numaPolicy)
	stateSinkObject.Save(6, &v.numaNodemask)
	stateSinkObject.Save(7, &v.id)
	stateSinkObject.Save(8, &v.hint)
	stateSinkObject.Save(9, &v.lastFault)
}

func (v *vma) afterLoad(context.Context) {}

// +checklocksignore
func (v *vma) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.mappable)
	stateSourceObject.Load(1, &v.off)
	stateSourceObject.Load(3, &v.dontfork)
	stateSourceObject.Load(4, &v.mlockMode)
	stateSourceObject.Load(5, &v.numaPolicy)
	stateSourceObject.Load(6, &v.numaNodemask)
	stateSourceObject.Load(7, &v.id)
	stateSourceObject.Load(8, &v.hint)
	stateSourceObject.Load(9, &v.lastFault)
	stateSourceObject.LoadValue(2, new(int), func(y any) { v.loadRealPerms(ctx, y.(int)) })
}

func (p *pma) StateTypeName() string {
	return "pkg/sentry/mm.pma"
}

func (p *pma) StateFields() []string {
	return []string{
		"file",
		"off",
		"translatePerms",
		"effectivePerms",
		"maxPerms",
		"needCOW",
		"private",
		"huge",
	}
}

func (p *pma) beforeSave() {}

// +checklocksignore
func (p *pma) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var fileValue string
	fileValue = p.saveFile()
	stateSinkObject.SaveValue(0, fileValue)
	stateSinkObject.Save(1, &p.off)
	stateSinkObject.Save(2, &p.translatePerms)
	stateSinkObject.Save(3, &p.effectivePerms)
	stateSinkObject.Save(4, &p.maxPerms)
	stateSinkObject.Save(5, &p.needCOW)
	stateSinkObject.Save(6, &p.private)
	stateSinkObject.Save(7, &p.huge)
}

func (p *pma) afterLoad(context.Context) {}

// +checklocksignore
func (p *pma) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(1, &p.off)
	stateSourceObject.Load(2, &p.translatePerms)
	stateSourceObject.Load(3, &p.effectivePerms)
	stateSourceObject.Load(4, &p.maxPerms)
	stateSourceObject.Load(5, &p.needCOW)
	stateSourceObject.Load(6, &p.private)
	stateSourceObject.Load(7, &p.huge)
	stateSourceObject.LoadValue(0, new(string), func(y any) { p.loadFile(ctx, y.(string)) })
}

func (s *pmaSet) StateTypeName() string {
	return "pkg/sentry/mm.pmaSet"
}

func (s *pmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *pmaSet) beforeSave() {}

// +checklocksignore
func (s *pmaSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []pmaFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *pmaSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *pmaSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]pmaFlatSegment), func(y any) { s.loadRoot(ctx, y.([]pmaFlatSegment)) })
}

func (n *pmanode) StateTypeName() string {
	return "pkg/sentry/mm.pmanode"
}

func (n *pmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *pmanode) beforeSave() {}

// +checklocksignore
func (n *pmanode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *pmanode) afterLoad(context.Context) {}

// +checklocksignore
func (n *pmanode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (p *pmaFlatSegment) StateTypeName() string {
	return "pkg/sentry/mm.pmaFlatSegment"
}

func (p *pmaFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (p *pmaFlatSegment) beforeSave() {}

// +checklocksignore
func (p *pmaFlatSegment) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Start)
	stateSinkObject.Save(1, &p.End)
	stateSinkObject.Save(2, &p.Value)
}

func (p *pmaFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (p *pmaFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Start)
	stateSourceObject.Load(1, &p.End)
	stateSourceObject.Load(2, &p.Value)
}

func (m *SpecialMappable) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappable"
}

func (m *SpecialMappable) StateFields() []string {
	return []string{
		"SpecialMappableRefs",
		"fr",
		"name",
	}
}

func (m *SpecialMappable) beforeSave() {}

// +checklocksignore
func (m *SpecialMappable) StateSave(stateSinkObject state.Sink) {
	m.beforeSave()
	stateSinkObject.Save(0, &m.SpecialMappableRefs)
	stateSinkObject.Save(1, &m.fr)
	stateSinkObject.Save(2, &m.name)
}

// +checklocksignore
func (m *SpecialMappable) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &m.SpecialMappableRefs)
	stateSourceObject.Load(1, &m.fr)
	stateSourceObject.Load(2, &m.name)
	stateSourceObject.AfterLoad(func() { m.afterLoad(ctx) })
}

func (r *SpecialMappableRefs) StateTypeName() string {
	return "pkg/sentry/mm.SpecialMappableRefs"
}

func (r *SpecialMappableRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *SpecialMappableRefs) beforeSave() {}

// +checklocksignore
func (r *SpecialMappableRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *SpecialMappableRefs) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(func() { r.afterLoad(ctx) })
}

func (s *vmaSet) StateTypeName() string {
	return "pkg/sentry/mm.vmaSet"
}

func (s *vmaSet) StateFields() []string {
	return []string{
		"root",
	}
}

func (s *vmaSet) beforeSave() {}

// +checklocksignore
func (s *vmaSet) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	var rootValue []vmaFlatSegment
	rootValue = s.saveRoot()
	stateSinkObject.SaveValue(0, rootValue)
}

func (s *vmaSet) afterLoad(context.Context) {}

// +checklocksignore
func (s *vmaSet) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.LoadValue(0, new([]vmaFlatSegment), func(y any) { s.loadRoot(ctx, y.([]vmaFlatSegment)) })
}

func (n *vmanode) StateTypeName() string {
	return "pkg/sentry/mm.vmanode"
}

func (n *vmanode) StateFields() []string {
	return []string{
		"nrSegments",
		"parent",
		"parentIndex",
		"hasChildren",
		"maxGap",
		"keys",
		"values",
		"children",
	}
}

func (n *vmanode) beforeSave() {}

// +checklocksignore
func (n *vmanode) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.nrSegments)
	stateSinkObject.Save(1, &n.parent)
	stateSinkObject.Save(2, &n.parentIndex)
	stateSinkObject.Save(3, &n.hasChildren)
	stateSinkObject.Save(4, &n.maxGap)
	stateSinkObject.Save(5, &n.keys)
	stateSinkObject.Save(6, &n.values)
	stateSinkObject.Save(7, &n.children)
}

func (n *vmanode) afterLoad(context.Context) {}

// +checklocksignore
func (n *vmanode) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.nrSegments)
	stateSourceObject.Load(1, &n.parent)
	stateSourceObject.Load(2, &n.parentIndex)
	stateSourceObject.Load(3, &n.hasChildren)
	stateSourceObject.Load(4, &n.maxGap)
	stateSourceObject.Load(5, &n.keys)
	stateSourceObject.Load(6, &n.values)
	stateSourceObject.Load(7, &n.children)
}

func (v *vmaFlatSegment) StateTypeName() string {
	return "pkg/sentry/mm.vmaFlatSegment"
}

func (v *vmaFlatSegment) StateFields() []string {
	return []string{
		"Start",
		"End",
		"Value",
	}
}

func (v *vmaFlatSegment) beforeSave() {}

// +checklocksignore
func (v *vmaFlatSegment) StateSave(stateSinkObject state.Sink) {
	v.beforeSave()
	stateSinkObject.Save(0, &v.Start)
	stateSinkObject.Save(1, &v.End)
	stateSinkObject.Save(2, &v.Value)
}

func (v *vmaFlatSegment) afterLoad(context.Context) {}

// +checklocksignore
func (v *vmaFlatSegment) StateLoad(ctx context.Context, stateSourceObject state.Source) {
	stateSourceObject.Load(0, &v.Start)
	stateSourceObject.Load(1, &v.End)
	stateSourceObject.Load(2, &v.Value)
}

func init() {
	state.Register((*aioManager)(nil))
	state.Register((*ioResult)(nil))
	state.Register((*AIOContext)(nil))
	state.Register((*aioMappable)(nil))
	state.Register((*aioMappableRefs)(nil))
	state.Register((*ioList)(nil))
	state.Register((*ioEntry)(nil))
	state.Register((*MemoryManager)(nil))
	state.Register((*vma)(nil))
	state.Register((*pma)(nil))
	state.Register((*pmaSet)(nil))
	state.Register((*pmanode)(nil))
	state.Register((*pmaFlatSegment)(nil))
	state.Register((*SpecialMappable)(nil))
	state.Register((*SpecialMappableRefs)(nil))
	state.Register((*vmaSet)(nil))
	state.Register((*vmanode)(nil))
	state.Register((*vmaFlatSegment)(nil))
}