package cloudflaregraphql

import (
	"context"
	"log"
	"time"
)

func (c *Client) GetZoneAnalyticsQuery(
	ctx context.Context,
	zoneTag *string,
	since string,
	until string,
) (*GetZoneAnalyticsResponse, error) {
	if c.opt.Debug {
		log.Println("GetZoneAnalyticsQuery", zoneTag, since, until)
	}
	return GetZoneAnalyticsQuery(ctx, c.gqlV4, zoneTag, since, until)
}

func (c *Client) GetWorkerAnalyticsQuery(
	ctx context.Context,
	zoneTag *string,
	datetime time.Time,
) (*GetWorkerAnalyticsResponse, error) {
	if c.opt.Debug {
		log.Println("GetWorkerAnalyticsQuery", zoneTag, datetime)
	}
	return GetWorkerAnalyticsQuery(ctx, c.gqlV4, zoneTag, datetime)
}
