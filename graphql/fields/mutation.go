package fields

import (
	"github.com/graphql-go/graphql"
	"github.com/jon20/graphql-sample/models"
)

func SetMutation() graphql.Fields {
	mutateQuery := graphql.Fields{
		"SampleMutate": &graphql.Field{
			Type:    SampleMutateType(),
			Args:    SampleMutateArgs(),
			Resolve: SampleMutateResolve,
		},
	}
	return mutateQuery
}

func SampleMutateType() *graphql.Object {
	sampleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SampleMutate",
		Fields: graphql.Fields{
			"Message": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	return sampleType
}

func SampleMutateArgs() map[string]*graphql.ArgumentConfig {
	sampleArgs := graphql.FieldConfigArgument{
		"Username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"Age": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
	return sampleArgs
}

func SampleMutateResolve(params graphql.ResolveParams) (interface{}, error) {
	resp := &models.SampleMutateResp{
		Message: "Hello",
	}
	return resp, nil
}
