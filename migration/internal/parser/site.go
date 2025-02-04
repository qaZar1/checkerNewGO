package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func (parser *Parser) parseReleases(html *html.Node) ([]string, error) {
	const regexp = `//*[@id][matches(@id, 'go([\d]+\.?){3}$')]`
	nodes, err := htmlquery.QueryAll(html, regexp)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(nodes))
	for _, node := range nodes {
		if node == nil {
			continue
		}
		list = append(list, htmlquery.SelectAttr(node, "id"))
	}

	return list, nil
}

func (parser *Parser) parseDescription(html *html.Node, version string) (string, error) {
	const re = `//*[@id][matches(@id, '%s')]`
	node, err := htmlquery.Query(html, fmt.Sprintf(re, version))
	if err != nil {
		return "", err
	}

	if node != nil {
		desc := htmlquery.InnerText(node)
		cleanDesc := strings.ReplaceAll(desc, "\n", " ")
		cleanDesc = strings.Join(strings.Fields(cleanDesc), " ")

		re := regexp.MustCompile(`^.+\(.+\)`)
		findString := re.FindString(cleanDesc)
		cleanDesc = strings.ReplaceAll(cleanDesc, findString, "")
		cleanDesc = strings.Trim(cleanDesc, " ")

		if cleanDesc != "" {
			cleanDesc = strings.ToUpper(string(cleanDesc[0])) + cleanDesc[1:]
			return cleanDesc, nil
		}
	}

	return fmt.Sprintf("https://go.dev/doc/devel/release#%s", version), nil
}

func (parser *Parser) parseDateRelease(html *html.Node, version string) (string, error) {
	const match = `//*[@id][matches(@id, '%s')]`
	const regexpStr = `([\d]+-?){3}`

	re := regexp.MustCompile(regexpStr)
	list, err := htmlquery.Query(html, fmt.Sprintf(match, version))
	if err != nil {
		return "", err
	}

	if list != nil {
		str := htmlquery.InnerText(list)
		cleanStr := re.FindString(str)
		return cleanStr, nil
	}

	return "0000-00-00", nil
}
