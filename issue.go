// Package ghstats provides functions to extract issues from GitHub and the
// related types.
package ghstats

import (
	"time"
)

type Issues struct {
	OrgId  int
	RepoId int
	Name   string
	buffer IssuesBuffer
}

// A GitHub Issue
type Issue struct {
	Id          int
	Number      int
	State       string
	Title       string
	Body        string
	User        User
	PullRequest *PullRequest `json:"pull_request"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	ClosedAt    time.Time    `json:"closed_at"`
}

type IssuesBuffer []Issue

// A GitHub pull request
type PullRequest struct {
	Url string
}

func (i *Issues) Table() Table {
	return Table{"issues", []Column{
		Column{"id", "INTEGER PRIMARY KEY"},
		Column{"org_id", "INTEGER"},
		Column{"repo_id", "INTEGER"},
		Column{"number", "INTEGER"},
		Column{"state", "TEXT"},
		Column{"title", "TEXT"},
		Column{"body", "TEXT"},
		Column{"user_id", "INTEGER"},
		Column{"pr_url", "TEXT"},
		Column{"created_at", "DATETIME"},
		Column{"updated_at", "DATETIME"},
		Column{"closed_at", "DATETIME"}}}

}

func (i *Issues) Values() []interface{} {
	items := i.buffer
	l := len(i.Table().Columns)
	v := make([]interface{}, l*len(items))

	for j, x := range items {
		o := j * l
		v[o+0] = x.Id
		v[o+1] = i.OrgId
		v[o+2] = i.RepoId
		v[o+3] = x.Number
		v[o+4] = x.State
		v[o+5] = x.Title
		v[o+6] = x.Body
		v[o+7] = x.User.Id
		v[o+8] = x.PullRequest.Url
		v[o+9] = x.CreatedAt
		v[o+10] = x.UpdatedAt
		v[o+11] = x.ClosedAt
	}
	return v
}

func (i *Issues) NewBuffer() Iterable {
	i.buffer = IssuesBuffer{}
	return &i.buffer
}

func (i *IssuesBuffer) Items() []interface{} {
	items := make([]interface{}, len(*i))
	for j, v := range *i {
		items[j] = v
	}
	return items
}
