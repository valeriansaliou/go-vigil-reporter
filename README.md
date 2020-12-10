# go-vigil-reporter

[![Test and Build](https://github.com/valeriansaliou/go-vigil-reporter/workflows/Test%20and%20Build/badge.svg?branch=master)](https://github.com/valeriansaliou/go-vigil-reporter/actions?query=workflow%3A%22Test+and+Build%22) [![Buy Me A Coffee](https://img.shields.io/badge/buy%20me%20a%20coffee-donate-yellow.svg)](https://www.buymeacoffee.com/valeriansaliou)

**Vigil Reporter for Golang. Used in pair with Vigil, the Microservices Status Page.**

Vigil Reporter is used to actively submit health information to Vigil from your apps. Apps are best monitored via application probes, which are able to report detailed system information such as CPU and RAM load. This lets Vigil show if an application host system is under high load.

**🇭🇺 Crafted in Budapest, Hungary.**

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
builder := Reporter.New("https://status.example.com", "YOUR_TOKEN_SECRET")

// Probe ID containing the parent Node for Replica
// Node ID containing Replica
// Unique Replica ID for instance (ie. your IP on the LAN)
// Reporting interval (in seconds; defaults to 30 seconds if not set)
reporter := builder.ProbeID("relay").NodeID("socket-client").ReplicaID("192.168.1.10").Interval(time.Duration(30 * time.Second)).Build()

// Run reporter (starts reporting)
reporter.Run()
```

## What is Vigil?

ℹ️ **Wondering what Vigil is?** Check out **[valeriansaliou/vigil](https://github.com/valeriansaliou/vigil)**.
