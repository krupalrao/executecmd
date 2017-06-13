package exec

import (
	"os/exec"
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-exec")


const (
	ivScriptType   = "scriptType"
	ivFilePath  = "filePath"

	ovResult = "result"
)

func init() {
	log.SetLogLevel(logger.InfoLevel)
}

// ExecActivity is an Activity that is used to execute shell, batch, python scripts
// inputs : {scriptType, filePath}
// outputs: {result}
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

	scriptType, _ := context.GetInput(ivScriptType).(string)
	filePath, _ := context.GetInput(ivFilePath).(string)

	log.Info(scriptType)
	
	cmd := exec.Command(filePath)
	out, err := cmd.Output()

	if err != nil {
		log.Debug(err.Error())
		context.SetOutput(ovResult, err.Error())
	} else {
		log.Info(string(out))
		context.SetOutput(ovResult, string(out))
	}
	
	return true, nil
}
