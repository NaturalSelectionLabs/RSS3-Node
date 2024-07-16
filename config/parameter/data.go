package parameter

import (
	"math/big"

	"github.com/rss3-network/protocol-go/schema/network"
)

// NumberOfMonthsToCover the number of months that a Node should cover data for
const NumberOfMonthsToCover = 3

type NetworkTolerance map[network.Network]uint64
type NetworkStartBlock map[network.Network]*big.Int
type NetworkCoreWorkerDiskSpacePerMonth map[network.Network]uint

// CurrentNetworkTolerance should be updated each epoch from vsl
var CurrentNetworkTolerance = NetworkTolerance{}

// CurrentNetworkStartBlock should be updated each epoch from vsl
var CurrentNetworkStartBlock = NetworkStartBlock{}

// CurrentNetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
// The data is calculated based on the average disk space usage during 2024 Q1.
// Actually usage may vary depending on the network's activity.
var CurrentNetworkCoreWorkerDiskSpacePerMonth = NetworkCoreWorkerDiskSpacePerMonth{}