package cmd

import (
	"encoding/json"
	"fmt"
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

	markdown := ""
	markdown += jsonToMarkdownRecursive(jsonData, 0)

	return markdown, nil
}

func jsonToMarkdownRecursive(data interface{}, depth int) string {
	markdown := ""
	indent := strings.Repeat("  ", depth)

	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			markdown += fmt.Sprintf("%s- **%s**: ", indent, key)
			markdown += jsonToMarkdownRecursive(value, depth+1)
			markdown += "\n"
		}
	case []interface{}:
		markdown += "\n"
		for _, item := range v {
			markdown += fmt.Sprintf("%s- ", indent)
			markdown += jsonToMarkdownRecursive(item, depth+1)
			markdown += "\n"
		}
	default:
		markdown += fmt.Sprintf("%v", v)
	}

	return markdown
}

func intSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strings.Join(strSlice, ",")
}
