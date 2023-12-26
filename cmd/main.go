package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("address", "localhost:9092", "")
)

func main() {
	flag.Parse()

	var promRegistry = prometheus.NewRegistry()
	promRegistry.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(
			collectors.GoRuntimeMetricsRule{
				Matcher: regexp.MustCompile("/.*"),
			},
		)),
	)

	gin.SetMode(gin.ReleaseMode)
	var httpRouter = gin.New()

	httpRouter.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"[%s] %s -> %s %s %d %s\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	}))

	var promMetricHandler = promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{EnableOpenMetrics: true})
	httpRouter.GET("/metrics", func(ctx *gin.Context) {
		promMetricHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})

	fmt.Fprintf(os.Stdout, "Listen on %s\n", *listenAddress)
	if err := httpRouter.Run(*listenAddress); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
