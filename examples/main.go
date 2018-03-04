// Copyright 2018 Valerian Saliou All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main


import (
  Reporter "github.com/valeriansaliou/go-vigil-reporter/vigil_reporter"
  "time"
)


func main() {
  // 1. Create Vigil Reporter
  builder := Reporter.New("http://[::1]:8080", "REPLACE_THIS_WITH_A_SECRET_KEY")

  reporter := builder.ProbeID("relay").NodeID("socket-client").ReplicaID("192.168.1.10").Interval(time.Duration(30 * time.Second)).Build()

  // 2. Run Vigil Reporter
  reporter.Run()

  // 3. Schedule Vigil Reporter end
  time.Sleep(80 * time.Second)
}
