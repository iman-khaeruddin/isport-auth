package dto

type BaseResponseList struct {
	PreviousPage interface{} `json:"previousPage"`
	CurrentPage  interface{} `json:"currentPage"`
	NextPage     interface{} `json:"nextPage"`
	Total        int64       `json:"total"`
	PerPage      interface{} `json:"perPage"`
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
}

type BaseResponse struct {
	Data         interface{} `json:"data"`
	Success      bool        `json:"success"`
	MessageTitle string      `json:"messageTitle"`
	Message      string      `json:"message"`
}

func DefaultErrorBaseResponseList(err error) BaseResponseList {
	return BaseResponseList{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}
}

func DefaultErrorBaseResponse(err error) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}
}

func DefaultBaseResponseWithError(err error) (BaseResponse, error) {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      err.Error(),
	}, err
}

func DefaultErrorBaseResponseWithMessage(message string) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      false,
		MessageTitle: "",
		Message:      message,
	}
}

func DefaultSuccessResponseWithMessage(msg string) BaseResponse {
	return BaseResponse{
		Data:         nil,
		Success:      true,
		MessageTitle: "Success",
		Message:      msg,
	}
}

func NewSuccessResponse(data any, messageTitle string) BaseResponse {
	return BaseResponse{
		Success:      true,
		Data:         data,
		MessageTitle: messageTitle,
		Message:      "Success",
	}
}
