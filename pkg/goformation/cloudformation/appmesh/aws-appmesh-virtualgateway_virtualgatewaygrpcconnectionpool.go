package appmesh

import (
	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/types"

	"github.com/weaveworks/eksctl/pkg/goformation/cloudformation/policies"
)

// VirtualGateway_VirtualGatewayGrpcConnectionPool AWS CloudFormation Resource (AWS::AppMesh::VirtualGateway.VirtualGatewayGrpcConnectionPool)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool.html
type VirtualGateway_VirtualGatewayGrpcConnectionPool struct {

	// MaxRequests AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool.html#cfn-appmesh-virtualgateway-virtualgatewaygrpcconnectionpool-maxrequests
	MaxRequests *types.Value `json:"MaxRequests"`

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
func (r *VirtualGateway_VirtualGatewayGrpcConnectionPool) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualGateway.VirtualGatewayGrpcConnectionPool"
}
