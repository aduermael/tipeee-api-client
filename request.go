package tipeee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) authenticate(req *http.Request) {
	if c.accessToken != "" {
		// req.Header.Add("access-token", c.accessToken)
	}
}

func (c *Client) authenticateInPath(path string) string {
	if c.accessToken != "" {
		if strings.Contains(path, "?") == false {
			path = path + "?"
		} else {
			path = path + "&"
		}
	}
	return path + "access_token=" + c.accessToken
}

func (c *Client) apiget(route string, res interface{}) error {

	path := "https://api.tipeee.com/v2.0" + route

	path = c.authenticateInPath(path)

	fmt.Println("GET", path)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	c.authenticate(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("STATUS CODE:", resp.StatusCode)
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil && jsonBytes != nil {
			fmt.Println(string(jsonBytes))
		}

		return fmt.Errorf("error %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) apigetJsonBytes(route string) ([]byte, error) {

	path := "https://api.tipeee.com/v2.0" + route

	path = c.authenticateInPath(path)

	fmt.Println("GET", path)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	c.authenticate(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("STATUS CODE:", resp.StatusCode)
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err == nil && jsonBytes != nil {
			fmt.Println(string(jsonBytes))
		}

		return nil, fmt.Errorf("error %d", resp.StatusCode)
	}

	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}
