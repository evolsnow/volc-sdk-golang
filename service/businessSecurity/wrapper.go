package businessSecurity

import (
	"encoding/json"
	"fmt"
)

// Synchronous risk detection
// 风险识别实时接口
func (p *BusinessSecurity) RiskDetection(req *RiskDetectionRequest) (*RiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("RiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("RiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("RiskDetection", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
			}
			result := new(RiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("RiskDetection: fail to do request, %v", err)
	}
	result := new(RiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Asynchronous risk detection
// 风险识别异步接口
func (p *BusinessSecurity) AsyncRiskDetection(req *AsyncRiskDetectionRequest) (*AsyncRiskDetectionResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("AsyncRiskDetection", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("AsyncRiskDetection", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(AsyncRiskDetectionResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(AsyncRiskDetectionResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

// Risk result
// 风险识别结果获取接口
func (p *BusinessSecurity) RiskResult(req *RiskResultRequest) (*RiskResultResponse, error) {
	respBody, _, err := p.Client.Query("RiskResult", req.ToQuery())
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Query("RiskResult", req.ToQuery())
			if err != nil {
				return nil, fmt.Errorf("RiskResult: fail to do request, %v", err)
			}
			result := new(RiskResultResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
	}
	result := new(RiskResultResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) DataReport(req *DataReportRequest) (*DataReportResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("AsyncRiskDetectionRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("DataReport", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("DataReport", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("AsyncRiskDetection: fail to do request, %v", err)
			}
			result := new(DataReportResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("DataReport: fail to do request, %v", err)
	}
	result := new(DataReportResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) ElementVerify(req *ElementVerifyRequest) (*ElementVerifyResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerify", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerify", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) ElementVerifyV2(req *ElementVerifyRequest) (*ElementVerifyResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ElementVerifyRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("ElementVerifyV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("ElementVerifyV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("ElementVerify: fail to do request, %v", err)
			}
			result := new(ElementVerifyResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("ElementVerifyV2: fail to do request, %v", err)
	}
	result := new(ElementVerifyResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileStatus(req *MobileStatusRequest) (*MobileStatusResponse, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatus", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatus", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
			}
			result := new(MobileStatusResponse)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
	}
	result := new(MobileStatusResponse)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (p *BusinessSecurity) MobileStatusV2(req *MobileStatusRequest) (*MobileStatusResponseV2, error) {
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("MobileSecondSaleRequest: fail to marshal request, %v", err)
	}

	respBody, _, err := p.Client.Json("MobileStatusV2", nil, string(reqData))
	if err != nil {
		// Retry on error
		// 支持错误重试
		if p.Retry() {
			respBody, _, err = p.Client.Json("MobileStatusV2", nil, string(reqData))
			if err != nil {
				return nil, fmt.Errorf("MobileStatusV2: fail to do request, %v", err)
			}
			result := new(MobileStatusResponseV2)
			if err := UnmarshalResultInto(respBody, result); err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("MobileStatus: fail to do request, %v", err)
	}
	result := new(MobileStatusResponseV2)
	if err := UnmarshalResultInto(respBody, result); err != nil {
		return nil, err
	}
	return result, nil
}
