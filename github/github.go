package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xpetit/x/v3"
)

type Config struct {
	Token string
	Org   string
}

type Repo struct {
	Name          string
	Archived      bool
	Disabled      bool
	Fork          bool
	IsTemplate    bool      `json:"is_template"`
	DefaultBranch string    `json:"default_branch"`
	PushedAt      time.Time `json:"pushed_at"`
}

func (c Config) FetchRepo(name string) (repo Repo, _ error) {
	return repo, c.request("repos", name, nil, &repo)
}

func (c Config) FetchAllRepositories() (repos []Repo, _ error) {
	for page := 1; ; page++ {
		tmp, err := c.fetchRepositories(page)
		if err != nil {
			return nil, err
		} else if len(tmp) == 0 {
			return
		}
		repos = append(repos, tmp...)
	}
}

func (c Config) fetchRepositories(page int) (repos []Repo, _ error) {
	return repos, c.request("orgs", "repos", url.Values{
		"per_page": {"100"},
		"page":     {strconv.Itoa(page)},
	}, &repos)
}

func (c Config) request(k, v string, values url.Values, data any) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%s/%s/%s", k, c.Org, v), nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = values.Encode()
	req.Header.Set("Authorization", "Bearer "+c.Token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer x.Closing(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(data)
}
