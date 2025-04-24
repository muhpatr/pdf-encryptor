package logger

import (
	"log"
	"os"
	"time"
)

var logfile *os.File

// func init() {
// 	var err error
// 	logfile, err = os.OpenFile("activity.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Failed to open log file:", err)
// 	}
// 	log.SetOutput(logfile)
// }

func LogAction(action, source, dest string, duration time.Duration, err error) {
	status := "SUCCESS"
	if err != nil {
		status = "FAILED"
	}
	log.Printf("[%s] %s: %s -> %s (%.2fms) %v\n",
		time.Now().Format(time.RFC3339),
		action, source, dest, float64(duration.Microseconds())/1000.0, status,
	)
}
