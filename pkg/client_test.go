package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStreamData(t *testing.T) {
	t.Run("Pagination normale", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			offset := r.URL.Query().Get("offset")
			switch offset {
			case "0":
				json.NewEncoder(w).Encode(PaginatedResponse{
					Data:       []interface{}{1, 2, 3},
					Pagination: struct{ Total, Offset, Limit int }{9, 0, 3},
				})
			case "3":
				json.NewEncoder(w).Encode(PaginatedResponse{
					Data:       []interface{}{4, 5, 6},
					Pagination: struct{ Total, Offset, Limit int }{9, 3, 3},
				})
			case "6":
				json.NewEncoder(w).Encode(PaginatedResponse{
					Data:       []interface{}{7, 8, 9},
					Pagination: struct{ Total, Offset, Limit int }{9, 6, 3},
				})
			}
		}))
		defer server.Close()

		client, _ := NewRTMSClient("fake-api-key", server.URL)
		dataChan, errChan := client.StreamData("/test", nil, 3)

		var results []interface{}
		for data := range dataChan {
			results = append(results, data)
		}

		if err := <-errChan; err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		if len(results) != 9 {
			t.Errorf("Nombre de résultats incorrect. Attendu : 9, Obtenu : %d", len(results))
		}
	})

	t.Run("Différentes tailles de lot", func(t *testing.T) {
		for _, batchSize := range []int{10, 50, 100} {
			t.Run(fmt.Sprintf("Taille de lot %d", batchSize), func(t *testing.T) {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					json.NewEncoder(w).Encode(PaginatedResponse{
						Data:       make([]interface{}, batchSize),
						Pagination: struct{ Total, Offset, Limit int }{batchSize, 0, batchSize},
					})
				}))
				defer server.Close()

				client, _ := NewRTMSClient("fake-api-key", server.URL)
				dataChan, errChan := client.StreamData("/test", nil, batchSize)

				var results []interface{}
				for data := range dataChan {
					results = append(results, data)
				}

				if err := <-errChan; err != nil {
					t.Errorf("Erreur inattendue : %v", err)
				}

				if len(results) != batchSize {
					t.Errorf("Nombre de résultats incorrect. Attendu : %d, Obtenu : %d", batchSize, len(results))
				}
			})
		}
	})

	t.Run("Erreur API", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erreur interne du serveur"))
		}))
		defer server.Close()

		client, _ := NewRTMSClient("fake-api-key", server.URL)
		dataChan, errChan := client.StreamData("/test", nil, 10)

		select {
		case <-dataChan:
			t.Error("Données reçues alors qu'une erreur était attendue")
		case err := <-errChan:
			if err == nil {
				t.Error("Erreur attendue, mais aucune erreur reçue")
			}
		case <-time.After(time.Second):
			t.Error("Timeout en attendant l'erreur")
		}
	})

	t.Run("Réponse JSON invalide", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("JSON invalide"))
		}))
		defer server.Close()

		client, _ := NewRTMSClient("fake-api-key", server.URL)
		dataChan, errChan := client.StreamData("/test", nil, 10)

		select {
		case <-dataChan:
			t.Error("Données reçues alors qu'une erreur était attendue")
		case err := <-errChan:
			if err == nil {
				t.Error("Erreur attendue, mais aucune erreur reçue")
			}
		case <-time.After(time.Second):
			t.Error("Timeout en attendant l'erreur")
		}
	})

	t.Run("Zéro résultat", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(PaginatedResponse{
				Data:       []interface{}{},
				Pagination: struct{ Total, Offset, Limit int }{0, 0, 10},
			})
		}))
		defer server.Close()

		client, _ := NewRTMSClient("fake-api-key", server.URL)
		dataChan, errChan := client.StreamData("/test", nil, 10)

		var results []interface{}
		for data := range dataChan {
			results = append(results, data)
		}

		if err := <-errChan; err != nil {
			t.Errorf("Erreur inattendue : %v", err)
		}

		if len(results) != 0 {
			t.Errorf("Nombre de résultats incorrect. Attendu : 0, Obtenu : %d", len(results))
		}
	})
}
