package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSSageMakerEndpoint AWS CloudFormation Resource (AWS::SageMaker::Endpoint)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html
type AWSSageMakerEndpoint struct {

	// EndpointConfigName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointconfigname
	EndpointConfigName string `json:"EndpointConfigName,omitempty"`

	// EndpointName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointname
	EndpointName string `json:"EndpointName,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-tags
	Tags []Tag `json:"Tags,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSageMakerEndpoint) AWSCloudFormationType() string {
	return "AWS::SageMaker::Endpoint"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSSageMakerEndpoint) MarshalJSON() ([]byte, error) {
	type Properties AWSSageMakerEndpoint
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSSageMakerEndpoint) UnmarshalJSON(b []byte) error {
	type Properties AWSSageMakerEndpoint
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSSageMakerEndpoint(*res.Properties)
	}

	return nil
}

// GetAllAWSSageMakerEndpointResources retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
func (t *Template) GetAllAWSSageMakerEndpointResources() map[string]AWSSageMakerEndpoint {
	results := map[string]AWSSageMakerEndpoint{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSSageMakerEndpoint:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SageMaker::Endpoint" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSSageMakerEndpoint
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSSageMakerEndpointWithName retrieves all AWSSageMakerEndpoint items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSageMakerEndpointWithName(name string) (AWSSageMakerEndpoint, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSSageMakerEndpoint:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SageMaker::Endpoint" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSSageMakerEndpoint
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSSageMakerEndpoint{}, errors.New("resource not found")
}