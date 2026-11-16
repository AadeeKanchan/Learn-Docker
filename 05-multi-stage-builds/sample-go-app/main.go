package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

type Response struct {
	Message string `json:"message"`
	Arch    string `json:"architecture"`
	GoVer   string `json:"go_version"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := Response{
			Message: "Super lean multi-stage Go service is running!",
			Arch:    runtime.GOARCH,
			GoVer:   runtime.Version(),
		}
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		metrics := map[string]interface{}{
			"alloc_kb":       m.Alloc / 1024,
			"total_alloc_kb": m.TotalAlloc / 1024,
			"sys_kb":         m.Sys / 1024,
			"num_gc":         m.NumGC,
		}
		json.NewEncoder(w).Encode(metrics)
	})

	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
