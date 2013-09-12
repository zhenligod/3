package engine

import (
	"code.google.com/p/mx3/util"
	"log"
)

type inputParam struct {
	lut
	upd_reg   [NREGION]func() []float64
	timestamp float64 // don't double-evaluate f(t)
	children  []derived
	descr
}

type derived interface {
	invalidate()
}

func (p *inputParam) init(nComp int, name, unit string, children []derived) {
	p.lut.init(nComp, p)
	p.descr = descr{name, unit}
	p.children = children
}

func (p *inputParam) update() {
	if p.timestamp != Time {
		changed := false
		// update functions of time
		for r := 0; r < regions.maxreg; r++ {
			updFunc := p.upd_reg[r]
			if updFunc != nil {
				p.bufset(r, updFunc())
				changed = true
			}
		}
		p.timestamp = Time
		if changed {
			p.invalidate()
		}
	}
}

func (p *inputParam) Cpu() [][NREGION]float32 {
	p.update()
	return p.cpu_buf
}

func (p *inputParam) setRegion(region int, v []float64) {
	util.Argument(len(v) == len(p.cpu_buf))
	p.upd_reg[region] = nil
	p.bufset(region, v)
}

func (p *inputParam) bufset(region int, v []float64) {
	for c := range p.cpu_buf {
		p.cpu_buf[c][region] = float32(v[c])
	}
}

func (p *inputParam) invalidate() {
	p.gpu_ok = false
	for _, c := range p.children {
		c.invalidate()
	}
}

type descr struct {
	name, unit string
}

func (d *descr) Name() string { return d.name }
func (d *descr) Unit() string { return d.unit }

// set in all regions except 0
// TODO: should region zero really have unset params?
// TODO: check if we always start from 1
func (p *inputParam) setUniform(v []float64) {
	for r := 1; r < NREGION; r++ {
		p.setRegion(r, v)
	}
}

func (p *inputParam) getRegion(region int) []float64 {
	cpu := p.Cpu()
	v := make([]float64, p.NComp())
	for i := range v {
		v[i] = float64(cpu[i][region])
	}
	return v
}

//func (p *inputParam) GetVec() []float64 {
//	return p.getRegion(1) // TODO: revise
//}
//
func (p *inputParam) getUniform() []float64 {
	v1 := p.getRegion(1)
	cpu := p.Cpu()
	for r := 2; r < regions.maxreg; r++ {
		for c := range v1 {
			if cpu[c][r] != float32(v1[c]) {
				log.Panicf("%v is not uniform", p.name)
			}
		}
	}
	return v1
}
