package main

import (
	comp "github.com/lz006/app-simulator/pkg/computation"
	mtrcs "github.com/lz006/app-simulator/pkg/metrics"
)

func main() {

	done := make(chan bool)

	computationResults := map[string]float64{}
	go comp.DoComputeSequence(&computationResults)
	go mtrcs.StartMetricsEndpoint(&computationResults)

	<-done
}
