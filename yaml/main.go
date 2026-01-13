package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func main() {
	root := &yaml.Node{
		Kind: yaml.DocumentNode,
		Content: []*yaml.Node{
			{
				Kind: yaml.MappingNode,
				Content: []*yaml.Node{
					{
						Kind:  yaml.ScalarNode,
						Value: "name",
						HeadComment: "Application name",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "my-app",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "port",
						HeadComment: "Server port",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "8080",
					},
				},
			},
		},
	}

	out, _ := yaml.Marshal(root)
	fmt.Println(string(out))
}
