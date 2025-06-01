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
	GoogleFonts []GoogleFont `json:"fonts"`
}

type GoogleFont struct {
	Family string `json:"family"`
	URL    string `json:"url"`
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

	log.Println("Successfully fetched fonts:", len(resp.GoogleFonts), "fonts found")
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

	var popularFontFamilies []GoogleFont
	for _, item := range resp.Items {
		popularFontFamilies = append(popularFontFamilies, GoogleFont{
			Family: item.Family,
			URL:    buildGoogleFontURL(item.Family),
		})
	}

	return &GetTrendingGoogleFontsResponse{
		GoogleFonts: popularFontFamilies,
	}, nil
}

func buildGoogleFontURL(fontFamily string) string {
	// TODO
	return "https://fonts.google.com/specimen/" + fontFamily
}
