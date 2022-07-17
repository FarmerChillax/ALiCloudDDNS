package main

import (
	_ "net/http/pprof"

	"github.com/FarmerChillax/ALiCloudDDNS/cmd"
)

const VERSION = "0.2.0"

func main() {
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	cmd.Execute()

	// go func() {
	// 	// http.ListenAndServe(":233", nil)
	// }()
}

func SetLogger() {
	// log.SetPrefix()
	// log.SetFlags(0)
	// log.SetOutput()
}
