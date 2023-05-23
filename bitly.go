package dachsurl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Bitly struct {
	url   string
	group string
}

type Group struct {
	Guid     string `json:"guid"`
	IsActive bool   `json:"is_active"`
}

func (bitly *Bitly) Groups(config *Config) ([]*Group, error) {
	request, err := http.NewRequest("GET", bitly.buildUrl("groups"), nil)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	return parseGroups(data)
}

func parseGroups(data []byte) ([]*Group, error) {
	result := struct {
		Groups []*Group `json:"groups"`
	}{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	returnValue := []*Group{}
	for _, g := range result.Groups {
		if g.IsActive {
			returnValue = append(returnValue, g)
		}
	}
	return returnValue, nil
}

func NewBitly(group string) *Bitly {
	return &Bitly{url: "https://api-ssl.bitly.com/v4/", group: group}
}

func (bitly *Bitly) buildUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", bitly.url, endpoint)
}

func handleGroup(config *Config, bitly *Bitly) (string, error) {
	if bitly.group != "" {
		return bitly.group, nil
	}
	gs, err := bitly.Groups(config)
	if err != nil {
		return "", err
	}
	if len(gs) == 0 {
		return "", fmt.Errorf("no active groups")
	}
	return gs[0].Guid, nil
}

func (bitly *Bitly) List(config *Config) ([]*ShortenUrl, error) {
	group, err := handleGroup(config, bitly)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("GET", bitly.buildUrl(fmt.Sprintf("/groups/%s/bitlinks?size=20", group)), nil)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	return handleListResponse(data, group)
}

func handleListResponse(data []byte, group string) ([]*ShortenUrl, error) {
	result := struct {
		Links []*ShortenUrl `json:"links"`
	}{}
	err := json.Unmarshal(data, &result)
	return removeDeletedLinks(result.Links, group), err
}

func removeDeletedLinks(links []*ShortenUrl, group string) []*ShortenUrl {
	result := []*ShortenUrl{}
	for _, link := range links {
		if !link.IsDeleted {
			link.Group = group
			result = append(result, link)
		}
	}
	return result
}

func (bitly *Bitly) Shorten(config *Config, url string) (*ShortenUrl, error) {
	reader := strings.NewReader(fmt.Sprintf(`{"long_url": "%s", "group_guid": "%s"}`, url, bitly.group))
	request, err := http.NewRequest("POST", bitly.buildUrl("shorten"), reader)
	if err != nil {
		return nil, err
	}
	data, err := sendRequest(request, config)
	if err != nil {
		return nil, err
	}
	return handleShortenResponse(data)
}

func handleShortenResponse(data []byte) (*ShortenUrl, error) {
	result := &ShortenUrl{}
	fmt.Println("result:", string(data))
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	result.Group, err = findGroup(data)
	return result, err
}

func findGroup(data []byte) (string, error) {
	r := struct {
		References struct {
			Group string `json:"group"`
		} `json:"references"`
	}{}
	err := json.Unmarshal(data, &r)
	uri := r.References.Group
	index := strings.LastIndex(uri, "/")
	return uri[index+1:], err
}

func (bitly *Bitly) Delete(config *Config, shortenURL string) error {
	request, err := http.NewRequest("DELETE", bitly.buildUrl("bitlinks/"+strings.TrimPrefix(shortenURL, "https://")), nil)
	if err != nil {
		return err
	}
	_, err = sendRequest(request, config)
	return err
}

func (bitly *Bitly) QRCode(config *Config, shortenURL string) ([]byte, error) {
	return []byte{}, fmt.Errorf("not implement yet")
}

func sendRequest(request *http.Request, config *Config) ([]byte, error) {
	response, err := sendRequestImpl(request, config)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	return handleResponse(response)
}

func sendRequestImpl(request *http.Request, config *Config) (*http.Response, error) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Token))
	request.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	return client.Do(request)
}

func handleResponse(response *http.Response) ([]byte, error) {
	if response.StatusCode/100 != 2 {
		data, _ := io.ReadAll(response.Body)
		fmt.Println("response body:", string(data))
		return []byte{}, fmt.Errorf("response status code %d", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}
