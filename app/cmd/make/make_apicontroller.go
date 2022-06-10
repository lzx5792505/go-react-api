package make

import (
	"fmt"
	"liu/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeAPIController = &cobra.Command{
	Use:   "apicontroller",
	Short: "Create api controller，exmaple: make apicontroller v1/user",
	Run:   runMakeAPIController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPIController(cmd *cobra.Command, args []string) {
	array := strings.Split(args[0], "/")
	if len(array) != 2 {
		console.Exit("api controller name format: v1/user")
	}
	apiVersion, name := array[0], array[1]
	model := makeModelFromString(name)
	filePath := fmt.Sprintf("app/http/controllers/api/%s/%sController.go", apiVersion, model.TableName)
	createFileFromStub(filePath, "apicontroller", model)
}
