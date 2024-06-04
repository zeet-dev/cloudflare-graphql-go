# Cloudflare GraphQL Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Cloudflare GraphQL Go is a Go client library SDK for interacting with Cloudflare's GraphQL API. This library provides a convenient and idiomatic way to query Cloudflare's analytics data using GraphQL.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- Easy-to-use Go client for Cloudflare's GraphQL API
- Supports querying analytics data for workers and zones
- Automatically generated Go types from GraphQL schema
- Includes examples for quick start

## Installation

To install the Cloudflare GraphQL Go library, use `go get`:

```sh
go get github.com/zeet-dev/cloudflare-graphql-go
```

## Usage

First, import the library in your Go code:

```go
import (
    "github.com/zeet-dev/cloudflare-graphql-go/pkg/cloudflaregraphql"
    "github.com/cloudflare/cloudflare-go"
)
```

Initialize the client:

```go
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
        result, err := client.GetZoneAnalyticsByDayQuery(ctx, &zoneTag, sinceDate, currentDate)
        if err != nil {
            log.Fatal(err)
        }

        log.Println("Zone Analytics Results:", zone.Name)
        for _, zoneResult := range result.Viewer.Zones[0].Zones {
            date := zoneResult.Dimensions.Timeslot
            value := zoneResult.Sum.Requests
            log.Printf("Date: %s, Requests: %d\n", date, value)
        }
    }
}
```

See [cmd/cfgo/cfgo.go](cmd/cfgo) for more complete examples

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Write your code and tests.
4. Ensure all tests pass.
5. Submit a pull request.

Please make sure to update tests as appropriate.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to the contributors and the open-source community for their continuous support.
- [Cloudflare](https://www.cloudflare.com) for providing the GraphQL API.
