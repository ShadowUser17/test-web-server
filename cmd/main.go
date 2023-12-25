package main

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/spf13/cobra"

	_ "github.com/prometheus/client_golang/prometheus/collectors"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
}
