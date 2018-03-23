package main

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Config struct {
	host   string
	user   string
	dbname string
	port   string
}

func Config_Open(config_file string) (*ini.File, error) {
	fd, err := ini.Load(config_file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return fd, nil
}

func Config_Read(config_file string) error {
	fd, err := Config_Open(config_file)
	if err != nil {
		return err
	}

	rdb_cfg, _ := RDB_Config_Read(fd)
	RDB_Config_Print(rdb_cfg)

	mongodb_cfg, _ := MONGODB_Config_Read(fd)
	MONGODB_Config_Print(mongodb_cfg)

	redis_cfg, _ := REDIS_Config_Read(fd)
	REDIS_Config_Print(redis_cfg)

	return nil
}

func RDB_Config_Read(fd *ini.File) (*Config, error) {
	host := fd.Section("RDB").Key("host").String()
	user := fd.Section("RDB").Key("user").String()
	dbname := fd.Section("RDB").Key("dbname").String()
	port := fd.Section("RDB").Key("port").String()

	return &Config{host, user, dbname, port}, nil
}

func MONGODB_Config_Read(fd *ini.File) (*Config, error) {
	host := fd.Section("MONGODB").Key("host").String()
	user := fd.Section("MONGODB").Key("user").String()
	dbname := fd.Section("MONGODB").Key("dbname").String()
	port := fd.Section("MONGODB").Key("port").String()

	return &Config{host, user, dbname, port}, nil
}

func REDIS_Config_Read(fd *ini.File) (*Config, error) {
	host := fd.Section("REDIS").Key("host").String()
	user := fd.Section("REDIS").Key("user").String()
	dbname := fd.Section("REDIS").Key("dbname").String()
	port := fd.Section("REDIS").Key("port").String()

	return &Config{host, user, dbname, port}, nil
}

////////////////////////////////////////////////////////////////////////////////
func Config_Print(pc_st *Config) {
	fmt.Println(pc_st.host)
	fmt.Println(pc_st.user)
	fmt.Println(pc_st.dbname)
	fmt.Println(pc_st.port)
}

func RDB_Config_Print(pc_st *Config) {
	Config_Print(pc_st)
}
func MONGODB_Config_Print(pc_st *Config) {
	Config_Print(pc_st)
}

func REDIS_Config_Print(pc_st *Config) {
	Config_Print(pc_st)
}

////////////////////////////////////////////////////////////////////////////////
func main() {
	file_name := "C:\\Projects\\lescpsn\\goprojects\\src\\GoDream\\xini\\xini.ini"
	Config_Read(file_name)
}
