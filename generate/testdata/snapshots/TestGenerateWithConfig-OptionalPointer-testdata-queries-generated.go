// Code generated by github.com/triarius/genqlient, DO NOT EDIT.

package queries

import (
	"context"

	"github.com/triarius/genqlient/graphql"
)

// ListInputQueryResponse is returned by ListInputQuery on success.
type ListInputQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *ListInputQueryUser `json:"user"`
}

// GetUser returns ListInputQueryResponse.User, and is useful for accessing the field via an interface.
func (v *ListInputQueryResponse) GetUser() *ListInputQueryUser { return v.User }

// ListInputQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type ListInputQueryUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id string `json:"id"`
}

// GetId returns ListInputQueryUser.Id, and is useful for accessing the field via an interface.
func (v *ListInputQueryUser) GetId() string { return v.Id }

// QueryWithSlicesResponse is returned by QueryWithSlices on success.
type QueryWithSlicesResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User *QueryWithSlicesUser `json:"user"`
}

// GetUser returns QueryWithSlicesResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesResponse) GetUser() *QueryWithSlicesUser { return v.User }

// QueryWithSlicesUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithSlicesUser struct {
	Emails                []string  `json:"emails"`
	EmailsOrNull          []string  `json:"emailsOrNull"`
	EmailsWithNulls       []*string `json:"emailsWithNulls"`
	EmailsWithNullsOrNull []*string `json:"emailsWithNullsOrNull"`
}

// GetEmails returns QueryWithSlicesUser.Emails, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmails() []string { return v.Emails }

// GetEmailsOrNull returns QueryWithSlicesUser.EmailsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsOrNull() []string { return v.EmailsOrNull }

// GetEmailsWithNulls returns QueryWithSlicesUser.EmailsWithNulls, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNulls() []*string { return v.EmailsWithNulls }

// GetEmailsWithNullsOrNull returns QueryWithSlicesUser.EmailsWithNullsOrNull, and is useful for accessing the field via an interface.
func (v *QueryWithSlicesUser) GetEmailsWithNullsOrNull() []*string { return v.EmailsWithNullsOrNull }

// __ListInputQueryInput is used internally by genqlient
type __ListInputQueryInput struct {
	Names []*string `json:"names"`
}

// GetNames returns __ListInputQueryInput.Names, and is useful for accessing the field via an interface.
func (v *__ListInputQueryInput) GetNames() []*string { return v.Names }

// The query or mutation executed by ListInputQuery.
const ListInputQuery_Operation = `
query ListInputQuery ($names: [String]) {
	user(query: {names:$names}) {
		id
	}
}
`

func ListInputQuery(
	ctx context.Context,
	client graphql.Client,
	names []*string,
) (*ListInputQueryResponse, error) {
	req := &graphql.Request{
		OpName: "ListInputQuery",
		Query:  ListInputQuery_Operation,
		Variables: &__ListInputQueryInput{
			Names: names,
		},
	}
	var err error

	var data ListInputQueryResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by QueryWithSlices.
const QueryWithSlices_Operation = `
query QueryWithSlices {
	user {
		emails
		emailsOrNull
		emailsWithNulls
		emailsWithNullsOrNull
	}
}
`

func QueryWithSlices(
	ctx context.Context,
	client graphql.Client,
) (*QueryWithSlicesResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithSlices",
		Query:  QueryWithSlices_Operation,
	}
	var err error

	var data QueryWithSlicesResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

