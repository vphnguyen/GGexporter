package entities

import  "github.com/prometheus/client_golang/prometheus"



// =============================
// Config config structure
// =============================
type Config struct {
    DSN     string
    Metrics map[string]struct {
        Type        string
        Description string
        Labels      []string
        Value       string
        metricDesc  *prometheus.Desc
    }
}