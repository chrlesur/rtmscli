package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func formatOutput(data []byte) (string, error) {
	switch outputFormat {
	case "json":
		return string(data), nil
	case "html":
		return jsonToHTML(data)
	case "markdown":
		return jsonToMarkdown(data)
	default:
		return "", fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func jsonToHTML(data []byte) (string, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}

	html := "<table>\n"
	html += jsonToHTMLRecursive(jsonData, 0)
	html += "</table>"

	return html, nil
}

func jsonToHTMLRecursive(data interface{}, depth int) string {
	html := ""
	indent := strings.Repeat("  ", depth)

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			html += fmt.Sprintf("%s<tr><th>%s</th><td>", indent, key)
			html += jsonToHTMLRecursive(value, depth+1)
			html += "</td></tr>\n"
		}
	case []interface{}:
		html += "<ul>\n"
		for _, item := range v {
			html += fmt.Sprintf("%s<li>", indent)
			html += jsonToHTMLRecursive(item, depth+1)
			html += "</li>\n"
		}
		html += indent + "</ul>\n"
	default:
		html += fmt.Sprintf("%v", v)
	}

	return html
}

func jsonToMarkdown(data []byte) (string, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}

	return jsonToMarkdownTable(jsonData), nil
}

func jsonToMarkdownTable(data interface{}) string {
	switch v := data.(type) {
	case map[string]interface{}:
		return mapToMarkdownTable(v)
	case []interface{}:
		return sliceToMarkdownTable(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func mapToMarkdownTable(m map[string]interface{}) string {
	var sb strings.Builder
	sb.WriteString("| Key | Value |\n|-----|-------|\n")

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := m[k]
		sb.WriteString(fmt.Sprintf("| %s | ", k))
		switch innerV := v.(type) {
		case map[string]interface{}:
			sb.WriteString(mapToMarkdownTable(innerV))
		case []interface{}:
			sb.WriteString(sliceToMarkdownTable(innerV))
		default:
			sb.WriteString(fmt.Sprintf("%v", innerV))
		}
		sb.WriteString(" |\n")
	}
	return sb.String()
}

func sliceToMarkdownTable(s []interface{}) string {
	if len(s) == 0 {
		return "Empty array"
	}

	var sb strings.Builder
	sb.WriteString("| Index | Value |\n|-------|-------|\n")

	for i, v := range s {
		sb.WriteString(fmt.Sprintf("| %d | ", i))
		switch innerV := v.(type) {
		case map[string]interface{}:
			sb.WriteString(mapToMarkdownTable(innerV))
		case []interface{}:
			sb.WriteString(sliceToMarkdownTable(innerV))
		default:
			sb.WriteString(fmt.Sprintf("%v", innerV))
		}
		sb.WriteString(" |\n")
	}
	return sb.String()
}

func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, ",")
}
