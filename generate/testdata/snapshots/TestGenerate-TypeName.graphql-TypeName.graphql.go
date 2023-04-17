// Code generated by github.com/triarius/genqlient, DO NOT EDIT.

package test

import (
	"github.com/triarius/genqlient/graphql"
	"github.com/triarius/genqlient/internal/testutil"
)

// TypeNameQueryResponse is returned by TypeNameQuery on success.
type TypeNameQueryResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User TypeNameQueryUser `json:"user"`
}

// GetUser returns TypeNameQueryResponse.User, and is useful for accessing the field via an interface.
func (v *TypeNameQueryResponse) GetUser() TypeNameQueryUser { return v.User }

// TypeNameQueryUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type TypeNameQueryUser struct {
	Typename string `json:"__typename"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	Id testutil.ID `json:"id"`
}

// GetTypename returns TypeNameQueryUser.Typename, and is useful for accessing the field via an interface.
func (v *TypeNameQueryUser) GetTypename() string { return v.Typename }

// GetId returns TypeNameQueryUser.Id, and is useful for accessing the field via an interface.
func (v *TypeNameQueryUser) GetId() testutil.ID { return v.Id }

// The query or mutation executed by TypeNameQuery.
const TypeNameQuery_Operation = `
query TypeNameQuery {
	user {
		__typename
		id
	}
}
`

func TypeNameQuery(
	client graphql.Client,
) (*TypeNameQueryResponse, error) {
	req := &graphql.Request{
		OpName: "TypeNameQuery",
		Query:  TypeNameQuery_Operation,
	}
	var err error

	var data TypeNameQueryResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
	)

	return &data, err
}

