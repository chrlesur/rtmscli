package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func formatOutput(data interface{}, format string) (string, error) {
	if data == nil {
		return "No data available", nil
	}

	// If data is of type []byte, try to decode it as JSON
	if byteData, ok := data.([]byte); ok {
		var jsonData interface{}
		err := json.Unmarshal(byteData, &jsonData)
		if err == nil {
			// If decoding is successful, use the decoded data
			data = jsonData
		}
		// If decoding fails, continue with raw data
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
		return "", fmt.Errorf("unsupported format: %s", format)
	}
}

func formatJSON(data interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error converting to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

func formatText(data interface{}) (string, error) {
	if data == nil {
		return "No data available", nil
	}

	var builder strings.Builder

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(data)
		for i := 0; i < value.Len(); i++ {
			builder.WriteString(fmt.Sprintf("Item %d:\n", i+1))
			builder.WriteString(strings.Repeat("=", 40) + "\n")
			builder.WriteString(formatTextItem(value.Index(i).Interface(), 0))
			builder.WriteString("\n")
		}
	default:
		builder.WriteString(formatTextItem(data, 0))
	}

	return builder.String(), nil
}

func formatTextItem(v interface{}, indent int) string {
	var builder strings.Builder
	indentStr := strings.Repeat("  ", indent)

	switch reflect.TypeOf(v).Kind() {
	case reflect.Map:
		val := reflect.ValueOf(v)
		maxKeyLength := 0
		for _, key := range val.MapKeys() {
			if len(key.String()) > maxKeyLength {
				maxKeyLength = len(key.String())
			}
		}
		for _, key := range val.MapKeys() {
			value := val.MapIndex(key)
			builder.WriteString(fmt.Sprintf("%s%-*s : ", indentStr, maxKeyLength, key))
			if value.Kind() == reflect.Map || value.Kind() == reflect.Slice {
				builder.WriteString("\n")
				builder.WriteString(formatTextItem(value.Interface(), indent+1))
			} else {
				builder.WriteString(formatValue(value.Interface()))
				builder.WriteString("\n")
			}
		}
	case reflect.Slice:
		val := reflect.ValueOf(v)
		if val.Len() == 0 {
			builder.WriteString("(empty)\n")
		} else {
			for i := 0; i < val.Len(); i++ {
				builder.WriteString(fmt.Sprintf("%s- ", indentStr))
				if val.Index(i).Kind() == reflect.Map || val.Index(i).Kind() == reflect.Slice {
					builder.WriteString("\n")
					builder.WriteString(formatTextItem(val.Index(i).Interface(), indent+1))
				} else {
					builder.WriteString(formatValue(val.Index(i).Interface()))
					builder.WriteString("\n")
				}
			}
		}
	default:
		builder.WriteString(formatValue(v))
		builder.WriteString("\n")
	}

	return builder.String()
}

func formatValue(v interface{}) string {
	if v == nil {
		return "(null)"
	}
	return fmt.Sprintf("%v", v)
}

func formatHTML(data interface{}) (string, error) {
	if data == nil {
		return "<p>No data available</p>", nil
	}

	var builder strings.Builder
	builder.WriteString("<html><head><style>")
	builder.WriteString("body { font-family: Arial, sans-serif; background-color: #f0f0f0; margin: 0; padding: 20px; }")
	builder.WriteString("table { border-collapse: separate; border-spacing: 0; width: 100%; background-color: white; box-shadow: 0 1px 3px rgba(0,0,0,0.2); border-radius: 6px; overflow: hidden; }")
	builder.WriteString("th, td { padding: 15px; text-align: left; }")
	builder.WriteString("th { background-color: #4CAF50; color: white; text-transform: uppercase; font-weight: bold; }")
	builder.WriteString("td { border-top: 1px solid #ddd; }")
	builder.WriteString("tr:nth-child(even) { background-color: #f8f8f8; }")
	builder.WriteString("tr:hover { background-color: #f1f1f1; }")
	builder.WriteString(".null { color: #999; font-style: italic; }")
	builder.WriteString(".empty { color: #999; font-style: italic; }")
	builder.WriteString("</style></head><body>")

	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		value := reflect.ValueOf(data)
		if value.Len() > 0 {
			builder.WriteString("<table>")
			// Table header
			firstItem := value.Index(0).Interface()
			if mapItem, ok := firstItem.(map[string]interface{}); ok {
				builder.WriteString("<tr>")
				for _, key := range getSortedKeys(mapItem) {
					builder.WriteString(fmt.Sprintf("<th>%s</th>", html.EscapeString(key)))
				}
				builder.WriteString("</tr>")
			}
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
		for _, key := range getSortedKeys(mapItem) {
			builder.WriteString("<td>")
			value := mapItem[key]
			switch v := value.(type) {
			case nil:
				builder.WriteString("<span class=\"null\">null</span>")
			case []interface{}:
				if len(v) == 0 {
					builder.WriteString("<span class=\"empty\">empty</span>")
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

func getSortedKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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
			builder.WriteString("### Item " + strconv.Itoa(i+1) + "\n\n")
			builder.WriteString(formatMarkdownItem(value.Index(i).Interface(), 0))
			builder.WriteString("\n")
		}
	default:
		builder.WriteString(formatMarkdownItem(data, 0))
	}

	return builder.String(), nil
}

func formatMarkdownItem(item interface{}, depth int) string {
	indent := strings.Repeat("  ", depth)
	var builder strings.Builder

	switch v := reflect.ValueOf(item); v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			value := v.MapIndex(key)
			builder.WriteString(fmt.Sprintf("%s- **%v**: ", indent, key))
			if value.Kind() == reflect.Map || value.Kind() == reflect.Slice {
				builder.WriteString("\n")
				builder.WriteString(formatMarkdownItem(value.Interface(), depth+1))
			} else {
				builder.WriteString(fmt.Sprintf("%v\n", value))
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			builder.WriteString(fmt.Sprintf("%s- ", indent))
			if v.Index(i).Kind() == reflect.Map || v.Index(i).Kind() == reflect.Slice {
				builder.WriteString("\n")
				builder.WriteString(formatMarkdownItem(v.Index(i).Interface(), depth+1))
			} else {
				builder.WriteString(fmt.Sprintf("%v\n", v.Index(i)))
			}
		}
	default:
		builder.WriteString(fmt.Sprintf("%v\n", item))
	}

	return builder.String()
}

func updateListCommand(cmd *cobra.Command, endpoint string, paramsFunc func() map[string]string) {
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		params := paramsFunc()

		// Add filter if specified
		if filter != "" {
			params["filter"] = filter
		}

		// Use StreamData to fetch data
		dataChan, errChan := client.StreamData(endpoint, params, batchSize)

		count := 0
		var data []interface{}

		// Channel to signal when data collection is done
		done := make(chan struct{})

		// Goroutine to collect data
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

		// Wait for data collection or error
		select {
		case <-done:
			// Data collection finished naturally
		case err := <-errChan:
			if err != nil {
				return fmt.Errorf("error fetching data: %w", err)
			}
		}

		// Check if any data was fetched
		if len(data) == 0 {
			fmt.Println("No data found.")
			return nil
		}

		// Format output
		output, err := formatOutput(data, outputFormat)
		if err != nil {
			return fmt.Errorf("error formatting output: %w", err)
		}

		// Print output
		fmt.Fprintln(os.Stdout, output)

		return nil
	}

	// Add common flags
	cmd.Flags().IntVar(&limit, "limit", 0, "Limit the number of results returned")
	cmd.Flags().IntVar(&batchSize, "batch-size", 100, "Number of items to fetch per batch")
	cmd.Flags().StringVar(&filter, "filter", "", "Filter results (format depends on the command)")
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
