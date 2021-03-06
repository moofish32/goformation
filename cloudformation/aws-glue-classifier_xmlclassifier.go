package cloudformation

// AWSGlueClassifier_XMLClassifier AWS CloudFormation Resource (AWS::Glue::Classifier.XMLClassifier)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html
type AWSGlueClassifier_XMLClassifier struct {

	// Classification AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-classification
	Classification string `json:"Classification,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-name
	Name string `json:"Name,omitempty"`

	// RowTag AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-xmlclassifier.html#cfn-glue-classifier-xmlclassifier-rowtag
	RowTag string `json:"RowTag,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueClassifier_XMLClassifier) AWSCloudFormationType() string {
	return "AWS::Glue::Classifier.XMLClassifier"
}
