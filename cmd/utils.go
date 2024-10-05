package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func formatOutput(data interface{}, format string) (string, error) {
	if data == nil {
		return "No data available", nil
	}

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
	case "html":
		return formatHTML(data)
	case "markdown":
		return formatMarkdown(data)
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
	if data == nil {
		return "nil", nil
	}
	return formatTextRecursive(data, 0), nil
}

func formatTextRecursive(v interface{}, indent int) string {
	if v == nil {
		return "nil"
	}

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

func formatHTML(data interface{}) (string, error) {
	if data == nil {
		return "<p>No data available</p>", nil
	}

	var builder strings.Builder
	builder.WriteString("<html><head><style>")
	builder.WriteString("body { font-family: Arial, sans-serif; }")
	builder.WriteString("table { border-collapse: collapse; width: 100%; }")
	builder.WriteString("th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }")
	builder.WriteString("th { background-color: #f2f2f2; }")
	builder.WriteString("</style></head><body>")

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(data)
		if value.Len() > 0 {
			builder.WriteString("<table>")
			// Table header
			builder.WriteString("<tr>")
			firstItem := value.Index(0).Interface()
			if mapItem, ok := firstItem.(map[string]interface{}); ok {
				for key := range mapItem {
					builder.WriteString(fmt.Sprintf("<th>%s</th>", html.EscapeString(key)))
				}
			}
			builder.WriteString("</tr>")
			// Table rows
			for i := 0; i < value.Len(); i++ {
				builder.WriteString(formatHTMLItem(value.Index(i).Interface()))
			}
			builder.WriteString("</table>")
		}
	default:
		builder.WriteString(formatHTMLItem(data))
	}

	builder.WriteString("</body></html>")
	return builder.String(), nil
}

func formatHTMLItem(item interface{}) string {
	var builder strings.Builder
	if mapItem, ok := item.(map[string]interface{}); ok {
		builder.WriteString("<tr>")
		for _, value := range mapItem {
			builder.WriteString("<td>")
			switch v := value.(type) {
			case nil:
				builder.WriteString("<em>null</em>")
			case []interface{}:
				if len(v) == 0 {
					builder.WriteString("<em>empty</em>")
				} else {
					builder.WriteString("<ul>")
					for _, subItem := range v {
						builder.WriteString(fmt.Sprintf("<li>%s</li>", html.EscapeString(fmt.Sprintf("%v", subItem))))
					}
					builder.WriteString("</ul>")
				}
			default:
				builder.WriteString(html.EscapeString(fmt.Sprintf("%v", v)))
			}
			builder.WriteString("</td>")
		}
		builder.WriteString("</tr>")
	} else {
		builder.WriteString(fmt.Sprintf("<p>%s</p>", html.EscapeString(fmt.Sprintf("%v", item))))
	}
	return builder.String()
}

func formatMarkdown(data interface{}) (string, error) {
	if data == nil {
		return "No data available", nil
	}

	var builder strings.Builder

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(data)
		for i := 0; i < value.Len(); i++ {
			builder.WriteString(formatMarkdownItem(value.Index(i).Interface(), 0))
			builder.WriteString("\n\n")
		}
	default:
		builder.WriteString(formatMarkdownItem(data, 0))
	}

	return builder.String(), nil
}

func formatMarkdownItem(item interface{}, depth int) string {
	if item == nil {
		return "nil"
	}

	var builder strings.Builder
	indent := strings.Repeat("  ", depth)

	switch v := reflect.ValueOf(item); v.Kind() {
	case reflect.Map:
		if v.IsNil() {
			return "nil"
		}
		for _, key := range v.MapKeys() {
			builder.WriteString(fmt.Sprintf("%s- **%v**: ", indent, key))
			value := v.MapIndex(key)
			if value.IsValid() {
				builder.WriteString(formatMarkdownItem(value.Interface(), depth+1))
			} else {
				builder.WriteString("nil")
			}
			builder.WriteString("\n")
		}
	case reflect.Slice, reflect.Array:
		if v.IsNil() {
			return "nil"
		}
		for i := 0; i < v.Len(); i++ {
			builder.WriteString(fmt.Sprintf("%s- ", indent))
			value := v.Index(i)
			if value.IsValid() {
				builder.WriteString(formatMarkdownItem(value.Interface(), depth+1))
			} else {
				builder.WriteString("nil")
			}
			builder.WriteString("\n")
		}
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return formatMarkdownItem(v.Elem().Interface(), depth)
	default:
		return fmt.Sprintf("%v", item)
	}

	return builder.String()
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
