package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/xpetit/x/v4"
)

type Config struct {
	Token    string
	Username string

	kind string // organization or username
}

func (c Config) request(k, v string, values url.Values, data any) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%s/%s/%s", k, c.Username, v), nil)
	if err != nil {
		return err
	}
	req.URL.RawQuery = values.Encode()
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

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

type Repo struct {
	PushedAt      time.Time `json:"pushed_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Language      string    `json:"language"`
	DefaultBranch string    `json:"default_branch"`
	Archived      bool      `json:"archived"`
	Disabled      bool      `json:"disabled"`
	Fork          bool      `json:"fork"`
	IsTemplate    bool      `json:"is_template"`
}

func (c Config) GetRepo(name string) (Repo, error) {
	var repo Repo
	if err := c.request("repos", name, nil, &repo); err != nil {
		return Repo{}, err
	}
	return repo, nil
}

func (c Config) listRepositories(page int) ([]Repo, error) {
	var repos []Repo

	if c.kind == "" {
		if err1 := c.request("orgs", "repos", url.Values{"per_page": {"1"}}, &repos); err1 == nil {
			c.kind = "orgs"
		} else if err2 := c.request("users", "repos", url.Values{"per_page": {"1"}}, &repos); err2 == nil {
			c.kind = "users"
		} else {
			return nil, errors.Join(err1, err2)
		}
	}

	if err := c.request(c.kind, "repos", url.Values{
		"per_page": {"100"},
		"page":     {strconv.Itoa(page)},
	}, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (c Config) ListRepositories() ([]Repo, error) {
	var repos []Repo
	for page := 1; ; page++ {
		tmp, err := c.listRepositories(page)
		if err != nil {
			return nil, err
		} else if len(tmp) == 0 {
			return repos, nil
		}
		repos = append(repos, tmp...)
	}
}
