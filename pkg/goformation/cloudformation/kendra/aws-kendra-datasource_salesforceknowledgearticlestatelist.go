package kendra

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// DataSource_SalesforceKnowledgeArticleStateList AWS CloudFormation Resource (AWS::Kendra::DataSource.SalesforceKnowledgeArticleStateList)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticlestatelist.html
type DataSource_SalesforceKnowledgeArticleStateList struct {

	// SalesforceKnowledgeArticleStateList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceknowledgearticlestatelist.html#cfn-kendra-datasource-salesforceknowledgearticlestatelist-salesforceknowledgearticlestatelist
	SalesforceKnowledgeArticleStateList *types.Value `json:"SalesforceKnowledgeArticleStateList,omitempty"`

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
func (r *DataSource_SalesforceKnowledgeArticleStateList) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.SalesforceKnowledgeArticleStateList"
}
