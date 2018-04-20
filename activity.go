package executecmd

import (
	"os/exec"
	"runtime"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-executecmd")

// ExecActivity is an Activity that is used to execute command
type ExecActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ExecActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ExecActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - executes the script
func (a *ExecActivity) Eval(context activity.Context) (done bool, err error) {

	command, _ := context.GetInput("command").(string)

	log.Info("command:", command)
	result := ""
	var err error
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", command)
		out, err := cmd.Output()
		result := string(out)
	} else {
		cmd := exec.Command("bash", "-c", command)
		out, err := cmd.Output()
		result := string(out)
	}

	if err != nil {
		log.Debug(err.Error())
		context.SetOutput("result", err.Error())
	} else {
		log.Info(result)
		context.SetOutput("result", result)
	}

	return true, nil
}
