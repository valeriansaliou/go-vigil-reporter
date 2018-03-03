# go-vigil-reporter

[![Build Status](https://img.shields.io/travis/valeriansaliou/go-vigil-reporter/master.svg)](https://travis-ci.org/valeriansaliou/go-vigil-reporter)

**Vigil Reporter for Golang. Used in pair with Vigil, the Microservices Status Page.**

Vigil Reporter is used to actively submit health information to Vigil from your apps. Apps are best monitored via application probes, which are able to report detailed system information such as CPU and RAM load. This lets Vigil show if an application host system is under high load.

## Who uses it?

_👋 You use vigil-reporter and you want to be listed there? [Contact me](https://valeriansaliou.name/)._

## How to use?

### Create reporter

`vigil-reporter` can be instantiated as such:

```go
import (
  Reporter "github.com/valeriansaliou/go-vigil-reporter/vigil_reporter"
  "time"
)

// Build reporter
// `page_url` + `reporter_token` from Vigil `config.cfg`
reporter := Reporter.New("https://status.example.com", "YOUR_TOKEN_SECRET")
  .ProbeID("relay")                           // Probe ID containing the parent Node for Replica
  .NodeID("socket-client")                    // Node ID containing Replica
  .ReplicaID("192.168.1.10")                  // Unique Replica ID for instance (ie. your IP on the LAN)
  .Interval(time.Duration(30 * time.Second))  // Reporting interval (in seconds; defaults to 30 seconds if not set)
  .Build()

// Run reporter (starts reporting)
reporter.Run()
```

## What is Vigil?

ℹ️ **Wondering what Vigil is?** Check out **[valeriansaliou/vigil](https://github.com/valeriansaliou/vigil)**.
