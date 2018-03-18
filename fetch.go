package ghstats

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vrde/ghstats/utils"
	"log"
	"net/http"
)

const IssueUrl = "https://api.github.com/repos/%s/issues"

// Fetch issues from a URL and return the next URL to follow for even moar
// issues.
func fetchIssues(ctx *Context, url string, issues *Issues) (error, string) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "token "+ctx.GitHubToken)
	resp, err := client.Do(req)

	if err != nil {
		return err, ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("GET <%s> failed with status: %d", url, resp.StatusCode)), ""
	}

	if err := json.NewDecoder(resp.Body).Decode(issues); err != nil {
		return err, ""
	}

	return nil, utils.NextLinkHeader(resp.Header.Get("Link"))
}

// Fetch all the issues from a GitHub repository, and send them to a channel.
func FetchIssues(ctx *Context, repository string, ch chan<- *Issues) {
	var err error
	url := fmt.Sprintf(IssueUrl, repository)

	for {
		i := &Issues{}

		log.Println("Fetching", url)
		err, url = fetchIssues(ctx, url, i)

		if err != nil {
			log.Println(err)
			close(ch)
			return
		}

		ch <- i

		if url == "" {
			close(ch)
			return
		}

	}
}
