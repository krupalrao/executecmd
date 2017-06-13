# tibco-log
This activity allows your flogo application to execute external shell, batch or python script.


## Installation

```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/exec
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "scriptType",
      "type": "string",
      "required": true,
      "allowed" : ["Shell", "Batch", "Python"]
    },
    {
      "name": "filePath",
      "type": "string",
	  "required": true,
      "value": ""
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| scriptType   | The type of script - Shell, batch or python |         
| filePath  | Script file path |
| result | The execution result of the script  |


## Configuration Examples
### Simple
Configure a task to execute a script:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "tibco-exec",
  "name": "Execute Script",
  "attributes": [
    { "name": "scriptType", "value": "Python" }
	{ "name": "filePath", "value": "./test.py" }
  ]
}
```
