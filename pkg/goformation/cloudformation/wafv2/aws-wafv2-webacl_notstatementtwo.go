package wafv2

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// WebACL_NotStatementTwo AWS CloudFormation Resource (AWS::WAFv2::WebACL.NotStatementTwo)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-notstatementtwo.html
type WebACL_NotStatementTwo struct {

	// Statement AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-notstatementtwo.html#cfn-wafv2-webacl-notstatementtwo-statement
	Statement *WebACL_StatementThree `json:"Statement,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *WebACL_NotStatementTwo) AWSCloudFormationType() string {
	return "AWS::WAFv2::WebACL.NotStatementTwo"
}
