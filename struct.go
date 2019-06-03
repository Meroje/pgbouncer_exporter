// Copyright (c) 2017 Kristoffer K Larsen <kristoffer@larsen.so>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"database/sql"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type columnUsage int

const (
	// LABEL Use this column as a label
	LABEL columnUsage = iota
	// COUNTER Use this column as a counter
	COUNTER columnUsage = iota
	// GAUGE Use this column as a gauge
	GAUGE columnUsage = iota
)

// MetricMapNamespace Groups metric maps under a shared set of labels
type MetricMapNamespace struct {
	columnMappings map[string]MetricMap // Column mappings in this namespace
	labels         []string
}

// MetricMap Stores the prometheus metric description which a given column will be mapped
// to by the collector
type MetricMap struct {
	discard    bool                              // Should metric be discarded during mapping?
	vtype      prometheus.ValueType              // Prometheus valuetype
	desc       *prometheus.Desc                  // Prometheus descriptor
	conversion func(interface{}) (float64, bool) // Conversion function to turn PG result into float64
}

type ColumnMapping struct {
	usage       columnUsage
	metric      string
	factor      float64
	description string
}

// Exporter collects PgBouncer stats from the given server and exports
// them using the prometheus metrics package.
type Exporter struct {
	mutex sync.RWMutex

	duration, up, error prometheus.Gauge
	totalScrapes        prometheus.Counter

	metricMap map[string]MetricMapNamespace

	db *sql.DB
}
