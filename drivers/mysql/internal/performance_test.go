package driver

import (
	"testing"

	"github.com/datazip-inc/olake/utils/testutils"
)

func TestMySQLPerformance(t *testing.T) {
	config := &testutils.PerformanceTest{
		TestConfig:     testutils.GetTestConfig("mysql"),
		Namespace:      "performance",
		BackfillStream: "users",
		CDCStream:      "users_cdc",
		ExecuteQuery:   ExecuteQueryPerformance,
		SupportsCDC:    true,
	}

	config.TestPerformance(t)
}
