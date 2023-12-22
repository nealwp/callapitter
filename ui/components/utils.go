package ui

import (
    "bytes"
    "encoding/json"
)

func PrettyPrintJSON(input string) (string) {
    var pretty bytes.Buffer
    if err := json.Indent(&pretty, []byte(input), "", "    "); err != nil {
        return input  
    }
    return pretty.String()
}
