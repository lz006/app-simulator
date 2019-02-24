package computation

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func DoComputeSequence(results *map[string]float64) {
	updateInterval, _ := strconv.Atoi(os.Getenv("EXEC_INTERVAL"))
	//updateInterval := 10000

	for true {
		// Mandelbot
		execDuration := mandelbrotExec()
		(*results)["mandelbrot"] = execDuration

		// Pause before starting next calculation
		time.Sleep(time.Duration(updateInterval * 1000000))
	}

}

func mandelbrotExec() float64 {
	start := time.Now()
	MandelbrotSet()
	elapsed := time.Since(start)
	milliSecs := float64(elapsed / time.Millisecond)
	log.Printf("Mandelbrot took %s", fmt.Sprintf("%f", milliSecs))
	return milliSecs
}
