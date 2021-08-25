package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// InterfaceNoFragmentsQueryResponse is returned by InterfaceNoFragmentsQuery on success.
type InterfaceNoFragmentsQueryResponse struct {
	Root InterfaceNoFragmentsQueryRootTopic `json:"root"`
}

// InterfaceNoFragmentsQueryRootTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRootTopic struct {
	// ID is documented in the Content interface.
	Id       testutil.ID                                         `json:"id"`
	Name     string                                              `json:"name"`
	Children []InterfaceNoFragmentsQueryRootTopicChildrenContent `json:"children"`
}

// InterfaceNoFragmentsQueryRootTopicChildrenArticle includes the requested fields of the GraphQL type Article.
type InterfaceNoFragmentsQueryRootTopicChildrenArticle struct {
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRootTopicChildrenContent includes the requested fields of the GraphQL type Content.
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type InterfaceNoFragmentsQueryRootTopicChildrenContent interface {
	implementsGraphQLInterfaceInterfaceNoFragmentsQueryRootTopicChildrenContent()
}

func (v InterfaceNoFragmentsQueryRootTopicChildrenArticle) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRootTopicChildrenContent() {
}
func (v InterfaceNoFragmentsQueryRootTopicChildrenVideo) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRootTopicChildrenContent() {
}
func (v InterfaceNoFragmentsQueryRootTopicChildrenTopic) implementsGraphQLInterfaceInterfaceNoFragmentsQueryRootTopicChildrenContent() {
}

// InterfaceNoFragmentsQueryRootTopicChildrenTopic includes the requested fields of the GraphQL type Topic.
type InterfaceNoFragmentsQueryRootTopicChildrenTopic struct {
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// InterfaceNoFragmentsQueryRootTopicChildrenVideo includes the requested fields of the GraphQL type Video.
type InterfaceNoFragmentsQueryRootTopicChildrenVideo struct {
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

func InterfaceNoFragmentsQuery(
	client graphql.Client,
) (*InterfaceNoFragmentsQueryResponse, error) {
	var retval InterfaceNoFragmentsQueryResponse
	err := client.MakeRequest(
		nil,
		"InterfaceNoFragmentsQuery",
		`
query InterfaceNoFragmentsQuery {
	root {
		id
		name
		children {
			id
			name
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

