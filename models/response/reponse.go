package response

type ResultResponse struct {
	ResponseCode   string      `json:"responseCode"`
	ReponseMessage string      `json:"reponseMessage"`
	ResultData     interface{} `json:"resultData"`
}

func (r ResultResponse) FmtResponse() ResultResponse {
	if r.ResultData == nil {
		r.ResultData = make(map[string]string)
	}
	return r
}

func (r *ResultResponse) CreateSuccess() {
	// resSuc := map[string]string{""}
	r.ResultData = "Create Success"
}
