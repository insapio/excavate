package models

// ReceiveMessage struct
type ReceiveMessage struct {
	UpdateID int `json:"update_id"`
}

// Structures Below for programming reference

// ENI struct
type ENI []struct {
	AvailabilityZone   string               `json:"AvailabilityZone"`
	Description        string               `json:"Description"`
	Groups             []Groups             `json:"Groups"`
	InterfaceType      string               `json:"InterfaceType"`
	MacAddress         string               `json:"MacAddress"`
	NetworkInterfaceID string               `json:"NetworkInterfaceId"`
	OwnerID            string               `json:"OwnerId"`
	PrivateDNSName     string               `json:"PrivateDnsName"`
	PrivateIPAddress   string               `json:"PrivateIpAddress"`
	PrivateIPAddresses []PrivateIPAddresses `json:"PrivateIpAddresses"`
	RequesterID        *string              `locationName:"requesterId" type:"string"`
	RequesterManaged   bool                 `json:"RequesterManaged"`
	SourceDestCheck    bool                 `json:"SourceDestCheck"`
	Status             string               `json:"Status"`
	SubnetID           string               `json:"SubnetId"`
	TagSet             []*Tag               `locationName:"tagSet" locationNameList:"item" type:"list"`
	VpcID              string               `json:"VpcId"`
}

// Groups struct
type Groups struct {
	GroupID   string `json:"GroupId"`
	GroupName string `json:"GroupName"`
}

// PrivateIPAddresses struct
type PrivateIPAddresses struct {
	Primary          bool   `json:"Primary"`
	PrivateDNSName   string `json:"PrivateDnsName"`
	PrivateIPAddress string `json:"PrivateIpAddress"`
}

// Tag struct
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	//
	// Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode
	// characters. May not begin with aws:.
	Key *string `locationName:"key" type:"string"`

	// The value of the tag.
	//
	// Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode
	// characters.
	Value *string `locationName:"value" type:"string"`
}
