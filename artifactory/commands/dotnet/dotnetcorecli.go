package dotnet

import (
	"github.com/jfrog/build-info-go/build/utils/dotnet"
	gofrogcmd "github.com/jfrog/gofrog/io"
	"github.com/jfrog/gofrog/version"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

const minDotnetSdkCoreVersionForAddSource = "3.1.200"

type DotnetCoreCliCommand struct {
	*DotnetCommand
}

func NewDotnetCoreCliCommand() *DotnetCoreCliCommand {
	dotnetCoreCliCmd := DotnetCoreCliCommand{&DotnetCommand{}}
	dotnetCoreCliCmd.SetToolchainType(dotnet.DotnetCore)
	return &dotnetCoreCliCmd
}

func (dccc *DotnetCoreCliCommand) Run() (err error) {
	dccc.useNugetAddSource, err = isDotnetVersionAboveMin()
	if err != nil {
		return err
	}
	return dccc.Exec()
}

func isDotnetVersionAboveMin() (bool, error) {
	// Run dotnet --version
	versionCmd, err := dotnet.NewToolchainCmd(dotnet.DotnetCore)
	if err != nil {
		return false, err
	}
	versionCmd.CommandFlags = []string{"--version"}

	output, err := gofrogcmd.RunCmdOutput(versionCmd)
	if err != nil {
		return false, err
	}

	dotNetSdkCoreVersion := version.NewVersion(output)
	log.Debug("using .NET SDK Core", output)
	return dotNetSdkCoreVersion.AtLeast(minDotnetSdkCoreVersionForAddSource), err
}
