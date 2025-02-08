package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// func (parser *Parser) ParseReleasesI() (bool, error) {
// 	// node, err := parser.api.MakeRequest(parser.url)
// 	// if err != nil {
// 	// 	fmt.Errorf("Can not make request: %s", err)
// 	// }

// 	data, err := os.ReadFile("output.txt")
// 	if err != nil {
// 		fmt.Errorf("%s", err)
// 	}

// 	buff := bytes.NewBuffer(data)

// 	html, err := htmlquery.Parse(buff)
// 	if err != nil {
// 		fmt.Errorf("%s", err)
// 	}

// 	allReleases, err := parser.parseReleases(html)
// 	if err != nil {
// 		fmt.Errorf("%s", err)
// 	}

// 	lastRelease := getHighestVersion(allReleases)
// 	ok, err := parser.apiVersions.GetVersion(lastRelease)
// 	if err != nil {
// 		fmt.Errorf("%s", err)
// 	}

// 	if !ok {
// 		desc, err := parser.parseDescription(html, "go1.23.0")
// 		if err != nil {
// 			fmt.Errorf("%s", err)
// 		}

// 		parser.apiVersions.AddVersion(models.Version{
// 			Version:        lastRelease,
// 			Description:    desc,
// 			Description_RU: "123456",
// 			ReleaseDate:    "0000-00-00",
// 		})
// 		fmt.Printf("Добавлена версия: %s. Описание: %s", lastRelease, desc)
// 		return true, nil
// 	}

// 	return false, nil
// }

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

func trimVersion(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return version
}

func getHighestVersion(versions []string) string {
	highest := versions[0]

	for _, version := range versions[1:] {
		if compareVersions(version, highest) < 0 {
			break
		}
		highest = version
	}

	return highest
}

func compareVersions(v1, v2 string) int {
	v1Parts := strings.Split(v1[2:], ".")
	v2Parts := strings.Split(v2[2:], ".")

	for i := 0; i < len(v1Parts) && i < len(v2Parts); i++ {
		if v1Parts[i] > v2Parts[i] {
			return 1
		} else if v1Parts[i] < v2Parts[i] {
			return -1
		}
	}
	return 0
}
