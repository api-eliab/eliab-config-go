package eliab

import (
	"testing"

	"github.com/jgolang/log"
)

func TestLoadConfiguration(t *testing.T) {

	loadConfiguration()

	log.Info(Get.General["SESSIONS"].PortServer)
	log.Info(Get.DataBase.DataBase)
	log.Info(Get.OneSignal)
	log.Info(Get.Services)
	log.Info(Get.AuthenticationToken)
	log.Info(Get.Mode)
	log.Info(Get.AppVersion)

}
