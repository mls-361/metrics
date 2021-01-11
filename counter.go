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
	// Counter AFAIRE.
	Counter interface {
		Dec()
		Inc()
	}

	counter struct {
		prwMutex  *sync.RWMutex
		CreatedAt int64 `json:"created_at"`
		Value     int64 `json:"value"`
		Timestamp int64 `json:"timestamp"`
	}
)

func (mcs *metrics) NewCounter(id string) Counter {
	c := &counter{
		prwMutex:  &mcs.rwMutex,
		CreatedAt: time.Now().UnixNano(),
	}

	mcs.Register(id, c)

	return c
}

// Dec AFAIRE.
func (c *counter) Dec() {
	c.prwMutex.Lock()

	c.Value--
	c.Timestamp = time.Now().UnixNano()

	c.prwMutex.Unlock()
}

// Inc AFAIRE.
func (c *counter) Inc() {
	c.prwMutex.Lock()

	c.Value++
	c.Timestamp = time.Now().UnixNano()

	c.prwMutex.Unlock()
}

/*
######################################################################################################## @(°_°)@ #######
*/
