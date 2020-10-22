package appsec

import (
	"context"
	"fmt"
	"net/http"
)

// SelectableHostnames represents a collection of SelectableHostnames
//
// See: SelectableHostnames.GetSelectableHostnames()
// API Docs: // appsec v1
//
// https://developer.akamai.com/api/cloud_security/application_security/v1.html

type (
	// SelectableHostnames  contains operations available on SelectableHostnames  resource
	// See: // appsec v1
	//
	// https://developer.akamai.com/api/cloud_security/application_security/v1.html#getselectablehostnames
	SelectableHostnames interface {
		GetSelectableHostnames(ctx context.Context, params GetSelectableHostnamesRequest) (*GetSelectableHostnamesResponse, error)
	}

	GetSelectableHostnamesRequest struct {
		ConfigID int `json:"configId"`
		Version  int `json:"version"`
	}

	GetSelectableHostnamesResponse struct {
		AvailableSet []struct {
			ActiveInProduction     bool   `json:"activeInProduction"`
			ActiveInStaging        bool   `json:"activeInStaging"`
			ArlInclusion           bool   `json:"arlInclusion"`
			Hostname               string `json:"hostname"`
			ConfigIDInProduction   int    `json:"configIdInProduction,omitempty"`
			ConfigNameInProduction string `json:"configNameInProduction,omitempty"`
		} `json:"availableSet"`
		ConfigID                int  `json:"configId"`
		ConfigVersion           int  `json:"configVersion"`
		ProtectARLInclusionHost bool `json:"protectARLInclusionHost"`
	}
)

func (p *appsec) GetSelectableHostnames(ctx context.Context, params GetSelectableHostnamesRequest) (*GetSelectableHostnamesResponse, error) {

	logger := p.Log(ctx)
	logger.Debug("GetSelectableHostnamess")

	var rval GetSelectableHostnamesResponse

	uri := fmt.Sprintf(
		"/appsec/v1/configs/%d/versions/%d/selectable-hostnames",
		params.ConfigID,
		params.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create getselectablehostnames request: %w", err)
	}

	resp, err := p.Exec(req, &rval)
	if err != nil {
		return nil, fmt.Errorf("getselectablehostnames request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, p.Error(resp)
	}

	return &rval, nil

}
