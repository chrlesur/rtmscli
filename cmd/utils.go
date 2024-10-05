package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func formatOutput(data interface{}, format string) (string, error) {
	// Si data est de type []byte, essayons de le décoder en JSON
	if byteData, ok := data.([]byte); ok {
		var jsonData interface{}
		err := json.Unmarshal(byteData, &jsonData)
		if err == nil {
			// Si le décodage réussit, utilisez les données décodées
			data = jsonData
		}
		// Si le décodage échoue, on continue avec les données brutes
	}

	switch strings.ToLower(format) {
	case "json":
		return formatJSON(data)
	case "text":
		return formatText(data)
	default:
		return "", fmt.Errorf("format non pris en charge : %s", format)
	}
}

func formatJSON(data interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("erreur lors de la conversion en JSON : %w", err)
	}
	return string(jsonBytes), nil
}

func formatText(data interface{}) (string, error) {
	return formatTextRecursive(data, 0), nil
}

func formatTextRecursive(v interface{}, indent int) string {
	indentStr := strings.Repeat("  ", indent)

	switch reflect.TypeOf(v).Kind() {
	case reflect.Map:
		var builder strings.Builder
		val := reflect.ValueOf(v)
		for _, key := range val.MapKeys() {
			builder.WriteString(fmt.Sprintf("%s%v: %s\n", indentStr, key, formatTextRecursive(val.MapIndex(key).Interface(), indent+1)))
		}
		return builder.String()
	case reflect.Slice, reflect.Array:
		var builder strings.Builder
		val := reflect.ValueOf(v)
		for i := 0; i < val.Len(); i++ {
			builder.WriteString(fmt.Sprintf("%s- %s\n", indentStr, formatTextRecursive(val.Index(i).Interface(), indent+1)))
		}
		return builder.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func updateListCommand(cmd *cobra.Command, endpoint string, paramsFunc func() map[string]string) {
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		params := paramsFunc()

		// Ajout du filtre s'il est spécifié
		if filter != "" {
			params["filter"] = filter
		}

		// Utilisation de StreamData pour récupérer les données
		dataChan, errChan := client.StreamData(endpoint, params, batchSize)

		count := 0
		var data []interface{}

		// Canal pour signaler l'arrêt de la collecte de données
		done := make(chan struct{})

		// Goroutine pour collecter les données
		go func() {
			defer close(done)
			for item := range dataChan {
				data = append(data, item)
				count++
				if limit > 0 && count >= limit {
					return
				}
			}
		}()

		// Attente de la fin de la collecte ou d'une erreur
		select {
		case <-done:
			// La collecte est terminée normalement
		case err := <-errChan:
			if err != nil {
				return fmt.Errorf("erreur lors de la récupération des données : %w", err)
			}
		}

		// Vérification si des données ont été récupérées
		if len(data) == 0 {
			fmt.Println("Aucune donnée n'a été trouvée.")
			return nil
		}

		// Formatage de la sortie
		output, err := formatOutput(data, outputFormat)
		if err != nil {
			return fmt.Errorf("erreur lors du formatage de la sortie : %w", err)
		}

		// Affichage de la sortie
		fmt.Fprintln(os.Stdout, output)

		return nil
	}

	// Ajout des flags communs
	cmd.Flags().IntVar(&limit, "limit", 0, "Limite le nombre de résultats retournés")
	cmd.Flags().IntVar(&batchSize, "batch-size", 100, "Nombre d'éléments à récupérer par lot")
	cmd.Flags().StringVar(&filter, "filter", "", "Filtre les résultats (format dépendant de la commande)")
	cmd.Flags().StringVar(&outputFormat, "output-format", "json", "Format de sortie (json, text)")
}

func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, ",")
}

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
