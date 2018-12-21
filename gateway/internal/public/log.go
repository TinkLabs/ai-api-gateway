package public

import (
	"encoding/json"
	"fmt"
)

type Message map[string]interface{}

func Logger(consolelog interface{}) []byte {
	output, _ := json.Marshal(consolelog)
	fmt.Println(string(output))
	return output
}
