package groups

type GroupStoreData struct {
	GroupID string `json:"group_id"`
}

type GroupResponse struct {
	Data struct {
		AffectedItems []struct {
			Name      string `json:"name"`
			Count     int    `json:"count"`
			MergedSum string `json:"mergedSum"`
			ConfigSum string `json:"configSum"`
		} `json:"affected_items"`
	} `json:"data"`
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type GroupCreatedResponse struct {
	Message string `json:"message"`
	Error   int    `json:"error"`
}

type GroupCreateError struct {
	Title       string                 `json:"title"`
	Detail      string                 `json:"detail"`
	Remediation string                 `json:"remediation"`
	DapiErrors  map[string]interface{} `json:"dapi_errors"`
	Error       int                    `json:"error"`
}
