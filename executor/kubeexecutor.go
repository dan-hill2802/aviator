package executor

import (
	"os/exec"
	"reflect"

	"github.com/JulzDiverse/aviator"
	"github.com/pkg/errors"
	"github.com/starkandwayne/goutils/ansi"
)

type KubeExecutor struct{}

func (e KubeExecutor) Command(cfg interface{}) (*exec.Cmd, error) {
	kube, ok := cfg.(aviator.Kube)
	if !ok {
		return &exec.Cmd{}, errors.New(ansi.Sprintf("@R{Type Assertion failed! Cannot assert %s to %s}", reflect.TypeOf(cfg), "aviator.Kube"))
	}

	apply := kube.Apply

	args := []string{
		"apply", "-f", apply.File,
	}

	if apply.Force {
		args = append(args, "--force")
	}

	if apply.DryRun {
		args = append(args, "--dry-run")
	}

	if apply.Overwrite {
		args = append(args, "--overwrite")
	}

	if apply.Recursive {
		args = append(args, "--recursive")
	}

	if apply.Output != "" {
		args = append(args, "--output", apply.Output)
	}

	return exec.Command("kubectl", args...), nil
}

func (e KubeExecutor) Execute(cmd *exec.Cmd, _ interface{}) error {
	return execCmd(cmd)
}