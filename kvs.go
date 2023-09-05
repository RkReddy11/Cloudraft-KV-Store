package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	store = map[string]string{
		"abc-1": "abc-1",
		"abc-2": "abc-2",
		"xyz-1": "xyz-1",
		"xyz-2": "xyz-2",
	}

	httpLatencyStatusKeys = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_latency_status_keys",
			Help:    "Histogram of latency of each endpoint by path, status, and no. of keys (in seconds)",
			Buckets: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		[]string{"method", "path_template", "status_code", "no_of_keys"},
	)

	httpResponse = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response",
			Help: "Total count of responses by method, path, and status codes and no. of keys",
		},
		[]string{"method", "path_template", "status_code", "no_of_keys"},
	)
)

func init() {
	prometheus.MustRegister(httpLatencyStatusKeys)
	prometheus.MustRegister(httpResponse)
}

// Structure of the incoming POST request
type Post struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	r := gin.Default()

	r.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		if value, ok := store[key]; ok {
			c.String(http.StatusOK, value)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		}
	})

	r.POST("/set", func(c *gin.Context) {
		var payload Post
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store[payload.Key] = payload.Value
		c.Status(http.StatusOK)
	})

	r.GET("/search", func(c *gin.Context) {
		prefix := c.Query("prefix")
		suffix := c.Query("suffix")
		res := []string{}

		if prefix != "" {
			for k := range store {
				if strings.HasPrefix(k, prefix) {
					res = append(res, store[k])
				}
			}
		}

		if suffix != "" {
			for k := range store {
				if strings.HasSuffix(k, suffix) {
					res = append(res, store[k])
				}
			}
		}

		c.JSON(http.StatusOK, res)
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	go func() {
		if err := r.Run(":8000"); err != nil {
			panic(err)
		}
	}()

	select {}
}
