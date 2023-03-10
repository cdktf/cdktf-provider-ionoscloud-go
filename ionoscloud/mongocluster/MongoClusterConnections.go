package mongocluster


type MongoClusterConnections struct {
	// The list of IPs and subnet for your cluster.
	//
	// Note the following unavailable IP ranges:
	//        10.233.64.0/18
	//        10.233.0.0/18
	//        10.233.114.0/24
	// example: [192.168.1.100/24, 192.168.1.101/24]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_cluster#cidr_list MongoCluster#cidr_list}
	CidrList *[]*string `field:"required" json:"cidrList" yaml:"cidrList"`
	// The datacenter to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_cluster#datacenter_id MongoCluster#datacenter_id}
	DatacenterId *string `field:"required" json:"datacenterId" yaml:"datacenterId"`
	// The LAN to connect your cluster to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/ionoscloud/r/mongo_cluster#lan_id MongoCluster#lan_id}
	LanId *string `field:"required" json:"lanId" yaml:"lanId"`
}

