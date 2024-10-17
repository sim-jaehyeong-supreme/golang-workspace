package usecase

import (
	"fmt"
	"mail-sampling/util"
)

func CallUseCase() {
	fmt.Println("CallUseCase")
	util.CallUtil("template/data_connector_execute_fail_mail.html")
}
