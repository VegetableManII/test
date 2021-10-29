package questions

import (
	"github.com/graphql-go/graphql"
)

type graphiQL struct {
	// query
}

const queryAllRoles = `query {
    allRoles(filter:JSON,pager:{},keyword:""){
        total
        nodes{
            id
            name
        }
    }
}`

const mutationAddRoles = `mutation {\n createRole(input:{
	  id:1,
	  name:"qiqiqi",
	  perms:[],
	}) {
	  id
	  name
	}
  }
`

type Role struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Desc   string `json:"desc"`
	//ExpireDate time.Time `json:"expireDate"`
	//Perms      []Perms   `json:"perms"`
	Active bool `json:"active"`
	//Category   string    `json:"category"`
}

var roles = []Role{
	{
		ID:     1,
		Name:   "业务方1",
		NameEn: "group1",
		Desc:   "验证码业务方",
		Active: true,
	},
	{
		ID:     2,
		Name:   "业务方2",
		NameEn: "group2",
		Desc:   "验证码业务方",
		Active: true,
	},
	{
		ID:     3,
		Name:   "业务方3",
		NameEn: "group3",
		Desc:   "验证码业务方",
		Active: true,
	},
}
var roleType = graphql.NewObject(graphql.ObjectConfig{})
var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createRole": &graphql.Field{
			Type:        roleType,
			Description: "用角色来区分不同业务方",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"name_en": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"desc": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"expireDate": &graphql.ArgumentConfig{
					Type: graphql.DateTime,
				},
				"active": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
			},
		},
	},
},
)

func UseGraphQL() {
	// json.NewEncoder()
}
