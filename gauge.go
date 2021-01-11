/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"sync"
	"time"
)

type (
	// GaugeFloat AFAIRE.
	GaugeFloat interface {
		Set(value float64)
		Add(value float64)
		Sub(value float64)
	}

	gaugeFloat struct {
		prwMutex  *sync.RWMutex
		CreatedAt int64   `json:"created_at"`
		Value     float64 `json:"value"`
		Timestamp int64   `json:"timestamp"`
	}

	// GaugeInt AFAIRE.
	GaugeInt interface {
		Set(value int64)
		Add(value int64)
		Sub(value int64)
	}

	gaugeInt struct {
		prwMutex  *sync.RWMutex
		CreatedAt int64 `json:"created_at"`
		Value     int64 `json:"value"`
		Timestamp int64 `json:"timestamp"`
	}
)

func (mcs *metrics) NewGaugeFloat(id string) GaugeFloat {
	g := &gaugeFloat{
		prwMutex:  &mcs.rwMutex,
		CreatedAt: time.Now().UnixNano(),
	}

	mcs.Register(id, g)

	return g
}

// Set AFAIRE.
func (g *gaugeFloat) Set(value float64) {
	g.prwMutex.Lock()

	g.Value = value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

// Add AFAIRE.
func (g *gaugeFloat) Add(value float64) {
	g.prwMutex.Lock()

	g.Value += value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

// Sub AFAIRE.
func (g *gaugeFloat) Sub(value float64) {
	g.prwMutex.Lock()

	g.Value -= value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

func (mcs *metrics) NewGaugeInt(id string) GaugeInt {
	g := &gaugeInt{
		prwMutex:  &mcs.rwMutex,
		CreatedAt: time.Now().UnixNano(),
	}

	mcs.Register(id, g)

	return g
}

// Set AFAIRE.
func (g *gaugeInt) Set(value int64) {
	g.prwMutex.Lock()

	g.Value = value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

// Add AFAIRE.
func (g *gaugeInt) Add(value int64) {
	g.prwMutex.Lock()

	g.Value += value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

// Sub AFAIRE.
func (g *gaugeInt) Sub(value int64) {
	g.prwMutex.Lock()

	g.Value -= value
	g.Timestamp = time.Now().UnixNano()

	g.prwMutex.Unlock()
}

/*
######################################################################################################## @(°_°)@ #######
*/
