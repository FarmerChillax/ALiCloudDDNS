package main

import (
	"fmt"

	"github.com/FarmerChillax/ALiCloudDDNS/client"
)

func main() {
	ddnsClient := client.New("www")
	fmt.Println("初始化成功:", ddnsClient)
	ddnsClient.Run()
}
