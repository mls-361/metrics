/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type (
	// Metrics AFAIRE.
	Metrics interface {
		Handler() http.HandlerFunc
		Register(id string, metric interface{})
		NewCounter(id string) Counter
		NewGaugeFloat(id string) GaugeFloat
		NewGaugeInt(id string) GaugeInt
	}

	metrics struct {
		rwMutex   sync.RWMutex
		CreatedAt int64                  `json:"created_at"`
		Metrics   map[string]interface{} `json:"metrics"`
	}
)

func New() Metrics {
	return &metrics{
		CreatedAt: time.Now().UnixNano(),
		Metrics:   make(map[string]interface{}),
	}
}

func (mcs *metrics) Handler() http.HandlerFunc {
	return func(rw http.ResponseWriter, _ *http.Request) {
		mcs.rwMutex.RLock()
		data, err := json.Marshal(mcs)
		mcs.rwMutex.RUnlock()

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")

		_, _ = rw.Write(data)
	}
}

func (mcs *metrics) Register(id string, metric interface{}) {
	mcs.rwMutex.Lock()
	mcs.Metrics[id] = metric
	mcs.rwMutex.Unlock()
}

/*
######################################################################################################## @(°_°)@ #######
*/
