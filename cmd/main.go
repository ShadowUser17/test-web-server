package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("address", "localhost:9092", "")
	logger        = log.New(os.Stdout, "", log.Ldate|log.Ltime)
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

	var promMetricHandler = promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{EnableOpenMetrics: true})
	httpRouter.GET("/metrics", func(ctx *gin.Context) {
		logger.Printf(
			"%s %s %s",
			ctx.Request.Host,
			ctx.Request.Method,
			ctx.Request.URL.Path,
		)
		promMetricHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})

	if err := httpRouter.Run(*listenAddress); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
