import aws from "aws-sdk";

var ec2 = new aws.EC2({ region: "us-east-1" });

var params = {};
ec2.describeNetworkInterfaces(params, function(err, data) {
  if (err) console.log(err, err.stack);
  // an error occurred
  else {
    data.NetworkInterfaces.map(eni => {
      eni.PrivateIpAddresses.map(address => {
        console.log(address);
      });
    });
  }
});
