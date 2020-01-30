package methods

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli"
)

//Getbasic sends a simple GET request to the url with any potential parameters like Tokens or Basic Auth
func Getbasic(c *cli.Context) error {
	var url = c.String("url")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if c.String("token") != "" {
		var bearer = "Bearer " + c.String("token")
		req.Header.Add("Authorization", bearer)
	}
	if c.String("u") != "" && c.String("p") != "" {
		un := c.String("u")
		pw := c.String("p")
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()
	s := formatresp(resp)
	fmt.Println(s)
	return nil
}
func getsend(c []Colls, ind int, method string) (string, error) {
	var url = c[0].Request[ind].URL + c[0].Request[ind].Path
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	if c[0].Request[ind].Token != "" {
		var bearer = "Bearer " + c[0].Request[ind].Token
		req.Header.Add("Authorization", bearer)
	}
	if c[0].Request[ind].User != "" && c[0].Request[ind].Pass != "" {
		un := c[0].Request[ind].User
		pw := c[0].Request[ind].Pass
		req.Header.Add("Authorization", "Basic "+basicAuth(un, pw))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()
	s := fmt.Sprintf("\nStatus:\t\t%s\n\nStatusCode:\t%d\n", resp.Status, resp.StatusCode)
	return s, nil
}
