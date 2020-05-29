package config

import "github.com/Terry-Mao/goconf"

var configLocation = "/etc/hcc/violin-scheduler/violin-scheduler.conf"

type schedulerConfig struct {
	MysqlConfig    *goconf.Section
	HTTPConfig     *goconf.Section
	RabbitMQConfig *goconf.Section
	FluteConfig    *goconf.Section
}
