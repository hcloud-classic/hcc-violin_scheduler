package end

import "hcc/violin-scheduler/lib/mysql"

func mysqlEnd() {
	if mysql.Db != nil {
		_ = mysql.Db.Close()
	}
}
