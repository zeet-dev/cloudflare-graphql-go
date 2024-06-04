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

		// last 3 hours
		currentTime := time.Now().UTC()
		sinceTime := time.Now().UTC().Add(-time.Hour * 3)
		result, err := client.GetZoneAnalyticsByHourQuery(ctx, &zoneTag, sinceTime, currentTime)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Zone Analytics Results:", zone.Name)
		for _, zoneResult := range result.Viewer.Zones[0].Zones {
			time := zoneResult.Dimensions.Timeslot
			value := zoneResult.Sum.Requests
			log.Printf("Time: %s, Requests: %d\n", time, value)
		}

		// last 3 days
		currentDate := time.Now().UTC().Format("2006-01-02")
		sinceDate := time.Now().UTC().AddDate(0, 0, -3).Format("2006-01-02")
		dayResult, err := client.GetZoneAnalyticsByDayQuery(ctx, &zoneTag, sinceDate, currentDate)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Zone Analytics Results:", zone.Name)
		for _, zoneResult := range dayResult.Viewer.Zones[0].Zones {
			date := zoneResult.Dimensions.Timeslot
			value := zoneResult.Sum.Requests
			log.Printf("Date: %s, Requests: %d\n", date, value)
		}

		// last 3 hours
		sinceTime = time.Now().UTC().Add(-time.Hour * 3)
		workerResult, err := client.GetWorkerAnalyticsByHourQuery(ctx, &zoneTag, sinceTime)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Worker Analytics Results:", zone.Name)
		for _, workerResult := range workerResult.Viewer.Zones[0].TotalRequestsData {
			time := workerResult.Dimensions.DatetimeHour
			value := workerResult.Sum.Requests
			log.Printf("Time: %s, Requests: %d\n", time, value)
		}
	}
}
