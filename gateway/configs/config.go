package configs

import (
    "encoding/json"
    "os"
    "log"
)

type Configuration map[string]interface{}
var cache Configuration

func Parms(key string) interface{} {

    env := os.Getenv("ENV")
    if cache[key] == nil {
        jsonStr, err := os.Open("configs/"+env+".config.json")
        if err != nil {
            log.Fatal("File error: %v\n", err)
        }
        json.NewDecoder(jsonStr).Decode(&cache)
    }

    switch key {
    case "NAME_HOST":
        return formatterHost(env, cache[key])
    default:
        if value, ok := cache[key]; ok {
            return value
        }
        return nil
	}
}

func formatterHost(env string, val interface{}) interface{} {
    if env == "prod" {
        val = "hk." + val.(string)
    }
    return val
}