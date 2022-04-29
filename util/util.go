package util

import (
	"encoding/json"
	"fmt"
)

func GetInt(field string, data map[string]interface{}, defVal ...int) int {

	ret := 0

	itemData, ok := data[field]

	if ok {

		ret, _ = itemData.(int)
	} else if len(defVal) > 0 {

		ret = defVal[0]
	}

	return ret
}

func GetString(field string, data map[string]interface{}, defVal ...string) string {

	ret := ""

	itemData, ok := data[field]

	if ok {

		ret, _ = itemData.(string)
	} else if len(defVal) > 0 {

		ret = defVal[0]
	}

	return ret
}

func GetBool(field string, data map[string]interface{}, defVal ...bool) bool {

	ret := false

	itemData, ok := data[field]

	if ok {

		ret, _ = itemData.(bool)
	} else if len(defVal) > 0 {

		ret = defVal[0]
	}

	return ret
}

func GetFloat(field string, data map[string]interface{}, defVal ...float64) float64 {

	var ret float64 = 0

	itemData, ok := data[field]

	if ok {

		ret, ok = itemData.(float64)

		if !ok {

			var ret32 float32 = 0

			ret32, ok = itemData.(float32)

			if ok {
				ret = float64(ret32)
			}
		}
	} else if len(defVal) > 0 {

		ret = defVal[0]
	}

	return ret
}

func GetMap(field string, data map[string]interface{}) map[string]interface{} {

	var ret map[string]interface{}

	itemData, ok := data[field]

	if ok {

		ret, _ = itemData.(map[string]interface{})
	}

	return ret
}

func GetArr(field string, data map[string]interface{}) []interface{} {

	var ret []interface{}

	itemData, ok := data[field]

	if ok {

		ret, _ = itemData.([]interface{})
	}

	return ret
}

func PrettyPrint(obj interface{}) {

	b, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Println(string(b))
}
