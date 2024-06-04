package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/zeet-dev/cloudflare-graphql-go/pkg/cloudflaregraphql"
)

func main() {
	cfApiToken := os.Getenv("CF_API_TOKEN")

	debug := os.Getenv("DEBUG") == "true"

	api, err := cloudflare.NewWithAPIToken(cfApiToken)
	if err != nil {
		log.Fatal(err)
	}

	client, err := cloudflaregraphql.New(
		func(o *cloudflaregraphql.ClientOption) {
			o.CloudflareAPIToken = cfApiToken
			o.Debug = debug
		})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	zones, err := api.ListZonesContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, zone := range zones.Result {
		zoneTag := zone.ID

		currentDate := time.Now().UTC().Format("2006-01-02")
		sinceDate := time.Now().UTC().AddDate(0, 0, -3).Format("2006-01-02")
		result, err := client.GetZoneAnalyticsQuery(ctx, &zoneTag, sinceDate, currentDate)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Zone Analytics Results:", zone.Name)
		for _, zoneResult := range result.Viewer.Zones[0].Zones {
			date := zoneResult.Dimensions.Timeslot
			value := zoneResult.Sum.Requests
			log.Printf("Date: %s, Requests: %d\n", date, value)
		}

		workerResult, err := client.GetWorkerAnalyticsQuery(ctx, &zoneTag, time.Now())
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Worker Analytics Results:", zone.Name)
		for _, workerResult := range workerResult.Viewer.Zones[0].TotalRequestsData {
			time := workerResult.Dimensions.DatetimeHour
			value := workerResult.Sum.Requests
			log.Printf("Date: %s, Requests: %d\n", time, value)
		}
	}
}
