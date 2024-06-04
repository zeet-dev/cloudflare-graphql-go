package cloudflaregraphql

import (
	"context"
	"log"
	"time"

	"github.com/samber/lo"
)

/*
- free_tier supports 30 days
*/
func (c *Client) GetZoneAnalyticsByDayQuery(
	ctx context.Context,
	zoneTag *string,
	since string,
	until string,
) (*GetZoneAnalyticsByDayResponse, error) {
	if c.opt.Debug {
		log.Println("GetZoneAnalyticsByDayQuery", zoneTag, since, until)
	}
	return GetZoneAnalyticsByDayQuery(ctx, c.gqlV4, zoneTag, since, until)
}

/*
- free_tier supports 24 hours
*/
func (c *Client) GetZoneAnalyticsByHourQuery(
	ctx context.Context,
	zoneTag *string,
	since time.Time,
	until time.Time,
) (*GetZoneAnalyticsByHourResponse, error) {
	if c.opt.Debug {
		log.Println("GetZoneAnalyticsByHourQuery", zoneTag, since, until)
	}
	return GetZoneAnalyticsByHourQuery(ctx, c.gqlV4, zoneTag, lo.ToPtr(since.UTC()), lo.ToPtr(until.UTC()))
}

/*
- free_tier supports 30 days
*/
func (c *Client) GetWorkerAnalyticsByHourQuery(
	ctx context.Context,
	zoneTag *string,
	since time.Time,
) (*GetWorkerAnalyticsByHourResponse, error) {
	if c.opt.Debug {
		log.Println("GetWorkerAnalyticsByHourQuery", zoneTag, since)
	}
	return GetWorkerAnalyticsByHourQuery(ctx, c.gqlV4, zoneTag, since.UTC())
}
