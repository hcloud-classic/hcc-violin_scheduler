package http

import (
	"encoding/json"
	"errors"
	violinSchedulerData "hcc/violin-scheduler/data"
	"hcc/violin-scheduler/lib/config"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DoHTTPRequest : Send http request to other modules with GraphQL query string.
func DoHTTPRequest(moduleName string, needData bool, data interface{}, query string) (interface{}, error) {
	client := &http.Client{Timeout: time.Duration(config.Flute.RequestTimeoutMs) * time.Millisecond}
	var url = "http://"
	switch moduleName {
	case "flute":
		url += config.Flute.ServerAddress + ":" + strconv.Itoa(int(config.Flute.ServerPort))
		break

	default:
		return nil, errors.New("unknown module name")
	}
	url += "/graphql?query=" + queryURLEncoder(query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// Check response
		respBody, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			result := string(respBody)

			if strings.Contains(result, "errors") {
				return nil, errors.New(result)
			}

			if needData {
				if data == nil {
					return nil, errors.New("needData marked as true but data is nil")
				}

				switch moduleName {
				case "flute":
					listNodeData := data.(violinSchedulerData.ListNodeData)
					err = json.Unmarshal([]byte(result), &(listNodeData))
					// fmt.Println("listNodeData: ", listNodeData)

					if err != nil {
						return nil, err
					}
					return listNodeData, nil

				default:
					return nil, errors.New("data is not supported for " + moduleName + " module")
				}
			}

			return result, nil
		}

		return nil, err
	}

	return nil, errors.New("http response returned error code")
}
