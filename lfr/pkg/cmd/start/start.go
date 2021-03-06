package start

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lgdd/liferay-cli/lfr/pkg/util/fileutil"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/printutil"
	"github.com/lgdd/liferay-cli/lfr/pkg/util/procutil"
	"github.com/spf13/cobra"
)

var (
	// Cmd is the command 'start' which allows to start the Liferay bundle
	Cmd = &cobra.Command{
		Use:   "start",
		Short: "Start a Liferay Tomcat bundle",
		Args:  cobra.NoArgs,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	startupScript, err := fileutil.GetTomcatScriptPath("startup")

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	tomcatPath, err := fileutil.GetTomcatPath()

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	err = os.Setenv("CATALINA_HOME", tomcatPath)

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	err = procutil.SetCatalinaPid()

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	startupCmd := exec.Command(startupScript)
	startupCmd.Stdout = os.Stdout

	err = startupCmd.Run()

	if err != nil {
		printutil.Danger(err.Error())
		os.Exit(1)
	}

	fmt.Println("\nFollow the logs:")
	printutil.Info("lfr logs -f\n\n")
}
