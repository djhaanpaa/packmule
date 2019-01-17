package config

import (
	"fmt"
	"log"
	"packmule/data"
	"gopkg.in/ini.v1"

)

func GetIniFileName() (fn string) {
	iniFileName := ""
	if data.PathExists("/etc/packmule/packmule.ini") {
		iniFileName = "/etc/packmule/appconfig.ini"
	} else {
		if data.PathExists("./packmule.ini") {
			iniFileName = "./packmule.ini"
		} else {
			log.Fatal("Application Configuration Not Found")
		}
	}
	fn = iniFileName
	return
}

func GetIniConfig() (c *ini.File) {
	fn := GetIniFileName()
	cfg, err := ini.Load(fn)

	if err != nil {
		log.Fatal(err)
	} else {
		c = cfg
	}
	return
}

func GetServiceStartCommand(svc string) (cmd string) {
	cfg := GetIniConfig()
	cmd = fmt.Sprintf(cfg.Section("").Key("service_start_command").String(), svc)

	return
}

func GetServiceStopCommand(svc string) (cmd string) {
	cfg := GetIniConfig()
	cmd = fmt.Sprintf(cfg.Section("").Key("service_stop_command").String(), svc)

	return
}