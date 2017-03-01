package core

type (
	APIResult struct {
		Success bool        `json:"success"`
		Result  interface{} `json:"result"`
		Error   *APIError   `json:"error"`
	}

	QueryResult struct {
		TotalCount int         `json:"totalCount"`
		Items      interface{} `json:"items"`
	}

	APIError struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details"`
	}
)

type (
	APIParam struct {
		SkipCount      int
		MaxResultCount int
		Fields         string
		Sort           string
		SortAsc        string
		SortDesc       string
	}
)
