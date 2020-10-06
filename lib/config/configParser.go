package config

import (
	"hcc/violin-scheduler/lib/logger"

	"github.com/Terry-Mao/goconf"
)

var conf = goconf.New()
var config = schedulerConfig{}
var err error

func parseMysql() {
	config.MysqlConfig = conf.Get("mysql")
	if config.MysqlConfig == nil {
		logger.Logger.Panicln("no mysql section")
	}

	Mysql = mysql{}
	Mysql.ID, err = config.MysqlConfig.String("id")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Mysql.Password, err = config.MysqlConfig.String("password")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Mysql.Address, err = config.MysqlConfig.String("address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Mysql.Port, err = config.MysqlConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Mysql.Database, err = config.MysqlConfig.String("database")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

//Will be Deprecated
func parseHTTP() {
	config.HTTPConfig = conf.Get("http")
	if config.HTTPConfig == nil {
		logger.Logger.Panicln("no http section")
	}

	HTTP = http{}
	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	HTTP.Port, err = config.HTTPConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseRabbitMQ() {
	config.RabbitMQConfig = conf.Get("rabbitmq")
	if config.RabbitMQConfig == nil {
		logger.Logger.Panicln("no rabbitmq section")
	}

	RabbitMQ = rabbitmq{}
	RabbitMQ.ID, err = config.RabbitMQConfig.String("rabbitmq_id")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	RabbitMQ.Password, err = config.RabbitMQConfig.String("rabbitmq_password")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	RabbitMQ.Address, err = config.RabbitMQConfig.String("rabbitmq_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	RabbitMQ.Port, err = config.RabbitMQConfig.Int("rabbitmq_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseFlute() {
	config.FluteConfig = conf.Get("flute")
	if config.FluteConfig == nil {
		logger.Logger.Panicln("no flute section")
	}

	Flute = flute{}
	Flute.ServerAddress, err = config.FluteConfig.String("flute_server_address")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Flute.ServerPort, err = config.FluteConfig.Int("flute_server_port")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Flute.RequestTimeoutMs, err = config.FluteConfig.Int("flute_request_timeout_ms")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Flute.WaitForLeaderNodeTimeoutSec, err = config.FluteConfig.Int("flute_wait_for_leader_node_timeout_sec")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseGrpc() {
	config.GrpcConfig = conf.Get("grpc")
	if config.GrpcConfig == nil {
		logger.Logger.Panicln("no grpc section")
	}

	Grpc.Port, err = config.GrpcConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

// Parser : Parse config file
func Parser() {
	if err = conf.Parse(configLocation); err != nil {
		logger.Logger.Panicln(err)
	}
	parseFlute()
	parseMysql()
	// parseHTTP()
	// parseRabbitMQ()
	parseGrpc()
}
