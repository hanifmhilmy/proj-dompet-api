package helpers

const (
	projDompet = "proj-dompet"
)

type (
	ResponseWrapper struct {
		Header HeaderResponse
		Result ResultResponse
	}

	HeaderResponse struct {
		Reason   string
		Messages []string
		Code     int
	}

	ResultResponse struct {
		Data interface{}
	}
)
