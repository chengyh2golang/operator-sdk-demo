// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"operator-sdk-demo/pkg/apis/app/v1alpha1.Demo":       schema_pkg_apis_app_v1alpha1_Demo(ref),
		"operator-sdk-demo/pkg/apis/app/v1alpha1.DemoSpec":   schema_pkg_apis_app_v1alpha1_DemoSpec(ref),
		"operator-sdk-demo/pkg/apis/app/v1alpha1.DemoStatus": schema_pkg_apis_app_v1alpha1_DemoStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Demo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Demo is the Schema for the demos API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("operator-sdk-demo/pkg/apis/app/v1alpha1.DemoSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("operator-sdk-demo/pkg/apis/app/v1alpha1.DemoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "operator-sdk-demo/pkg/apis/app/v1alpha1.DemoSpec", "operator-sdk-demo/pkg/apis/app/v1alpha1.DemoStatus"},
	}
}

func schema_pkg_apis_app_v1alpha1_DemoSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DemoSpec defines the desired state of Demo",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_app_v1alpha1_DemoStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DemoStatus defines the observed state of Demo",
				Type:        []string{"object"},
			},
		},
	}
}