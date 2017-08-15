package resources

// AWS::S3::Bucket.VersioningConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html
type AWSS3BucketVersioningConfiguration struct {

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-versioningconfig.html#cfn-s3-bucket-versioningconfig-status
	Status string `json:"Status"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketVersioningConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.VersioningConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketVersioningConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}