package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("address", "localhost:9092", "Set listen address.")
	metricsPath   = flag.String("metrics", "/metrics", "Set path for metrics.")
)

func main() {
	flag.Parse()

	var promRegistry = prometheus.NewRegistry()
	promRegistry.MustRegister(
		collectors.NewBuildInfoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(
			collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")},
		)),
	)

	var httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: "http_requests"},
		[]string{"method", "path"},
	)
	promRegistry.MustRegister(httpRequests)

	gin.SetMode(gin.ReleaseMode)
	var httpRouter = gin.New()

	var promMetricHandler = promhttp.HandlerFor(promRegistry, promhttp.HandlerOpts{EnableOpenMetrics: true})
	httpRouter.GET(*metricsPath, func(ctx *gin.Context) {
		promMetricHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})

	httpRouter.NoRoute(func(ctx *gin.Context) {
		ctx.Status(200)
		httpRequests.WithLabelValues(ctx.Request.Method, ctx.Request.URL.Path).Inc()
		fmt.Fprintf(os.Stdout, "%s -> %s %s %d\n",
			ctx.Request.RemoteAddr,
			ctx.Request.Method,
			ctx.Request.URL.Path,
			ctx.Writer.Status(),
		)

		switch ctx.Request.Method {
		case "GET":
			fmt.Fprintf(ctx.Writer, "%s\n", ctx.Request.RemoteAddr)

		case "POST", "PUT":
			fmt.Fprint(os.Stdout, "DATA: ")
			io.Copy(os.Stdout, ctx.Request.Body)
			fmt.Fprint(os.Stdout, "\n")
		}
	})

	fmt.Fprintf(os.Stdout, "Listen on %s\n", *listenAddress)
	if err := httpRouter.Run(*listenAddress); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
