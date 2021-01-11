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
		pmutex    *sync.RWMutex
		CreatedAt int64 `json:"created_at"`
		Value     int64 `json:"value"`
		Timestamp int64 `json:"timestamp"`
	}
)

func (ms *metrics) NewCounter(id string) Counter {
	c := &counter{
		pmutex:    &ms.mutex,
		CreatedAt: time.Now().UnixNano(),
	}

	ms.Register(id, c)

	return c
}

// Dec AFAIRE.
func (c *counter) Dec() {
	c.pmutex.Lock()

	c.Value--
	c.Timestamp = time.Now().UnixNano()

	c.pmutex.Unlock()
}

// Inc AFAIRE.
func (c *counter) Inc() {
	c.pmutex.Lock()

	c.Value++
	c.Timestamp = time.Now().UnixNano()

	c.pmutex.Unlock()
}

/*
######################################################################################################## @(°_°)@ #######
*/
