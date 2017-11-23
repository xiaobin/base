package go_base

//设备信息
type Device struct {
	TenancyCode  string `bson:"tenancyCode" json:"tenancyCode"`
	GroupCode    string `bson:"groupCode" json:"groupCode"`
	GroupName    string `bson:"groupName" json:"groupName"`
	TenancyType  string `bson:"tenancyType" json:"tenancyType"`
	TenantId     string `bson:"tenantId" json:"tenantId"`
	TenantName   string `bson:"tenantName" json:"tenantName"`
	DeviceUserId string `bson:"deviceUserId" json:"deviceUserId"`
	DeviceId     string `bson:"deviceId" json:"deviceId"`
	VersionId    string `bson:"versionId" json:"versionId"`
	RoomId       string `bson:"roomId" json:"roomId"`
	ServerIp     string `bson:"serverIp" json:"serverIp"`
}
