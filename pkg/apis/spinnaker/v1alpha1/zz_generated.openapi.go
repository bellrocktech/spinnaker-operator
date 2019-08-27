// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/spinnaker/v1alpha1.ExposeConfig":                 schema_pkg_apis_spinnaker_v1alpha1_ExposeConfig(ref),
		"./pkg/apis/spinnaker/v1alpha1.ExposeConfigService":          schema_pkg_apis_spinnaker_v1alpha1_ExposeConfigService(ref),
		"./pkg/apis/spinnaker/v1alpha1.ExposeConfigServiceOverrides": schema_pkg_apis_spinnaker_v1alpha1_ExposeConfigServiceOverrides(ref),
		"./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSource":          schema_pkg_apis_spinnaker_v1alpha1_SpinnakerFileSource(ref),
		"./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceStatus":    schema_pkg_apis_spinnaker_v1alpha1_SpinnakerFileSourceStatus(ref),
		"./pkg/apis/spinnaker/v1alpha1.SpinnakerService":             schema_pkg_apis_spinnaker_v1alpha1_SpinnakerService(ref),
		"./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceSpec":         schema_pkg_apis_spinnaker_v1alpha1_SpinnakerServiceSpec(ref),
		"./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceStatus":       schema_pkg_apis_spinnaker_v1alpha1_SpinnakerServiceStatus(ref),
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_ExposeConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExposeConfig represents the configuration for exposing Spinnaker",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/spinnaker/v1alpha1.ExposeConfigService"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.ExposeConfigService"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_ExposeConfigService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExposeConfigService represents the configuration for exposing Spinnaker using k8s services",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"overrides": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/spinnaker/v1alpha1.ExposeConfigServiceOverrides"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.ExposeConfigServiceOverrides"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_ExposeConfigServiceOverrides(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ExposeConfigServiceOverrides represents expose configurations of type service, overriden by specific services",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_SpinnakerFileSource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpinnakerFileSource represents a source for Spinnaker files",
				Properties: map[string]spec.Schema{
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "Config map reference if Spinnaker config stored in a configMap",
							Ref:         ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReference"),
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Config map reference if Spinnaker config stored in a secret",
							Ref:         ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReference"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_SpinnakerFileSourceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpinnakerFileSourceStatus represents a source for Spinnaker files",
				Properties: map[string]spec.Schema{
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "Config map reference if Spinnaker config stored in a configMap",
							Ref:         ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReferenceStatus"),
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Config map reference if Spinnaker config stored in a secret",
							Ref:         ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReferenceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceReferenceStatus"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_SpinnakerService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpinnakerService is the Schema for the spinnakerservices API",
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
							Ref: ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceSpec", "./pkg/apis/spinnaker/v1alpha1.SpinnakerServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_SpinnakerServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpinnakerServiceSpec defines the desired state of SpinnakerService",
				Properties: map[string]spec.Schema{
					"spinnakerConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSource"),
						},
					},
					"expose": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/spinnaker/v1alpha1.ExposeConfig"),
						},
					},
				},
				Required: []string{"spinnakerConfig"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.ExposeConfig", "./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSource"},
	}
}

func schema_pkg_apis_spinnaker_v1alpha1_SpinnakerServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SpinnakerServiceStatus defines the observed state of SpinnakerService",
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Current deployed version of Spinnaker",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastConfigurationTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the configuration was updated",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"halConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "Spinnaker Halyard configuration current configured",
							Ref:         ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceStatus"),
						},
					},
					"services": {
						SchemaProps: spec.SchemaProps{
							Description: "Services deployment information",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/spinnaker/v1alpha1.SpinnakerDeploymentStatus"),
									},
								},
							},
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "Indicates when all services are ready",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"uiUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "Exposed Deck URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "Exposed Gate URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/spinnaker/v1alpha1.SpinnakerDeploymentStatus", "./pkg/apis/spinnaker/v1alpha1.SpinnakerFileSourceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
