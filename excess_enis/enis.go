package enis

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2")

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

func (sess session.Session) AnalyzeENIs() {
	svc := ec2.New(sess, configObject)

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

	fmt.Print(result.NetworkInterfaces)

}
