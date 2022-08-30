package main

import (
	_ "net/http/pprof"

	"github.com/FarmerChillax/ALiCloudDDNS/cmd"
)

func main() {
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	cmd.Execute()
}
