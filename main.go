package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Uploads a file to S3 given a bucket and object key. Also takes a duration
// value to terminate the update if it doesn't complete within that time.
//
// The AWS Region needs to be provided in the AWS shared config or on the
// environment variable as `AWS_REGION`. Credentials also must be provided
// Will default to shared config file, but can load from environment if provided.
//
// Usage:
//   # Upload myfile.txt to myBucket/myKey. Must complete within 10 minutes or will fail
//   go run withContext.go -b mybucket -k myKey -d 10m < myfile.txt

// 🎉 Sprint 0 || Connect to AWS
// 🎉 Sprint 0 || shell credentials for now
// Sprint 2 || ^^ assume credentials
// 🎉 Sprint 0 || pull relevant data for First Check
// Sprint 1 || ^^ pull test config from API
// Sprint 2 || ^^ pull test config from API (for specific user)
// 👉 Sprint 1 || perform transforms/analyses (print)
/*
	- Check for tag signifying there is an exception for this
	- Get number of secondary IPs for this ENI
	- Get SG for ENI
		if - check if this is an EKS ENI SG
			- If so, get EKS cluster details
				- Check EKS cluster for exception tags
					- 🙅‍ submit recommmendation to set KEEP_WARM_TARGET on this cluster

*/
// Sprint 2 || ^^ pull transforms/analyses from API based on config (print)
// Sprint 2 || post results to API for UI display
// Sprint 2 || create API

func main() {

	// First Check: excess ips (ENIs) EKS reserved hidden ENIs.

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.
	svc := ec2.New(sess, &aws.Config{
		Region: aws.String("us-east-1"),
	})

	input := &ec2.DescribeNetworkInterfacesInput{}

	result, err := svc.DescribeNetworkInterfaces(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	/*
		func ENISecondaryIPs() int {
			total := 0
			for _, item := range basket.Items {
				total += item.Price
			}
			return total
		}

		fmt.Printf("Primary IP: %t\t| Security Group: %s |  \n",
			*result.NetworkInterfaces[0].PrivateIpAddresses[0].Primary,
			*result.NetworkInterfaces[0].Groups[0].GroupId)
	*/

	fmt.Print(result.NetworkInterfaces)
}
