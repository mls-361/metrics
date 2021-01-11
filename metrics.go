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
	}

	metrics struct {
		mutex     sync.RWMutex
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

func (ms *metrics) Handler() http.HandlerFunc {
	return func(rw http.ResponseWriter, _ *http.Request) {
		ms.mutex.RLock()
		data, err := json.Marshal(ms)
		ms.mutex.RUnlock()

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json; charset=utf-8")

		_, _ = rw.Write(data)
	}
}

func (ms *metrics) Register(id string, metric interface{}) {
	ms.mutex.Lock()

	ms.Metrics[id] = metric

	ms.mutex.Unlock()
}

/*
######################################################################################################## @(°_°)@ #######
*/
