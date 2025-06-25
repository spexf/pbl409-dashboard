package utils

func GetDeleteErrorMessage(dataMap map[string]interface{}) (string, bool) {
	data, ok := dataMap["data"].(map[string]interface{})
	if !ok {
		return "", false
	}
	failedItems, ok := data["failed_items"].([]interface{})
	if !ok || len(failedItems) == 0 {
		return "", false
	}
	firstItem, ok := failedItems[0].(map[string]interface{})
	if !ok {
		return "", false
	}
	errorField, ok := firstItem["error"].(map[string]interface{})
	if !ok {
		return "", false
	}
	message, ok := errorField["message"].(string)
	return message, ok
}
