package Framework

import (
	"Platform/framework/config"
	"Platform/framework/crontab"
	"Platform/framework/database"
)

func Init() {
	Config.Setup()
	Database.Init()
	Crontab.Start()
}
