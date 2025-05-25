package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type GetTrendingGoogleFontsResponse struct {
	Fonts []string `json:"fonts"`
}

func handleGetPopularGoogleFonts(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for popular Google Fonts")

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := getPopularGoogleFonts(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch trending fonts", http.StatusInternalServerError)
		log.Println("Error fetching trending fonts:", err)
		return
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("Error encoding response:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

	log.Println("Successfully fetched fonts:", len(resp.Fonts), "fonts found")
}

func getPopularGoogleFonts(ctx context.Context) (*GetTrendingGoogleFontsResponse, error) {
	resp, err := fetchGoogleFonts(ctx, GetGoogleFontsOptions{
		Sort: GoogleFontsSort{
			DescendingPopularity: true,
		},
	})
	if err != nil {
		return nil, err
	}

	var trendingFonts []string
	for _, item := range resp.Items {
		trendingFonts = append(trendingFonts, item.Family)
	}
	return &GetTrendingGoogleFontsResponse{
		Fonts: trendingFonts,
	}, nil
}
