package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	queryParamApiKey = "key"

	queryParamCapability = "capability"
	capabilityWoff2      = "WOFF2"

	queryParamSort = "sort"
	sortPopularity = "popularity"
	sortTrending   = "trending"
)

type GetGoogleFontsResponse struct {
	Kind  string           `json:"kind"`
	Items []GoogleFontItem `json:"items"`
}

type GoogleFontItem struct {
	Family       string            `json:"family"`
	Variants     []string          `json:"variants"`
	Subsets      []string          `json:"subsets"`
	Version      string            `json:"version"`
	LastModified string            `json:"lastModified"`
	Files        map[string]string `json:"files"`
	Category     string            `json:"category"`
	Kind         string            `json:"kind"`
	Menu         string            `json:"menu"`
}

type GetGoogleFontsOptions struct {
	Sort GoogleFontsSort
}

type GoogleFontsSort struct {
	DescendingPopularity bool
	DescendingTrending   bool
}

func fetchGoogleFonts(ctx context.Context, opts GetGoogleFontsOptions) (*GetGoogleFontsResponse, error) {
	apiKey := env.GoogleFontsApiKey
	googleFontsUrl := url.URL{
		Scheme: "https",
		Host:   "www.googleapis.com",
		Path:   "/webfonts/v1/webfonts",
		RawQuery: url.Values{
			queryParamApiKey:     []string{apiKey},
			queryParamCapability: []string{capabilityWoff2},
		}.Encode(),
	}
	if opts.Sort.DescendingPopularity {
		googleFontsUrl.Query().Add(queryParamSort, sortPopularity)
	}
	if opts.Sort.DescendingTrending {
		googleFontsUrl.Query().Add(queryParamSort, sortTrending)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", googleFontsUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBuf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response GetGoogleFontsResponse
	if err := json.Unmarshal(respBuf, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
