package mysql

import (
	"hcc/violin-scheduler/lib/config"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/lib/syscheck"
	"testing"
)

func Test_DB_Prepare(t *testing.T) {
	err := syscheck.CheckRoot()
	if err != nil {
		t.Fatal("Failed to get root permission!")
	}

	if !logger.Prepare() {
		t.Fatal("Failed to prepare logger!")
	}
	defer func() {
		_ = logger.FpLog.Close()
	}()

	config.Parser()

	err = Prepare()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = Db.Close()
	}()
}
