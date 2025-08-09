package stat

var ServiceGroupMemberFieldNames = struct {
	AverageServerTimeToFirstByte          string
	ClearStats                            string
	CurrentClientConnections              string
	CurrentLoad                           string
	CurrentReusePool                      string
	CurrentServerConnections              string
	FrustratingTimeToLastByteTransactions string
	HttpMaxHeaderSizePacketsCount         string
	Ip                                    string
	MaxClients                            string
	MaxHttpHeaderFieldLengthExceededCount string
	MaxOutOfOrderPacketsCount             string
	Name                                  string
	Port                                  string
	PrimaryIpAddress                      string
	PrimaryPort                           string
	RequestBytesRate                      string
	RequestsRate                          string
	ResponseBytesRate                     string
	ResponsesRate                         string
	ServerEstablishedConnections          string
	ServerName                            string
	ServiceGroupMemberBindingOrder        string
	ServiceOrder                          string
	ServiceType                           string
	State                                 string
	SurgeCount                            string
	ToleratingTimeToLastByteTransactions  string
	TotalConnectionReassemblyQueueFlush   string
	TotalConnectionReassemblyQueueFlush75 string
	TotalRequestBytes                     string
	TotalRequests                         string
	TotalResponseBytes                    string
	TotalResponses                        string
	TotalServerTimeToLastByteTransactions string
}{
	AverageServerTimeToFirstByte:          "avgsvrttfb",
	ClearStats:                            "clearstats",
	CurrentClientConnections:              "curclntconnections",
	CurrentLoad:                           "curload",
	CurrentReusePool:                      "curreusepool",
	CurrentServerConnections:              "cursrvrconnections",
	FrustratingTimeToLastByteTransactions: "frustratingttlbtransactions",
	HttpMaxHeaderSizePacketsCount:         "httpmaxhdrszpkts",
	Ip:                                    "ip",
	MaxClients:                            "maxclients",
	MaxHttpHeaderFieldLengthExceededCount: "httpmaxhdrfldlenpkts",
	MaxOutOfOrderPacketsCount:             "tcpmaxooopkts",
	Name:                                  "servicegroupname",
	Port:                                  "port",
	PrimaryIpAddress:                      "primaryipaddress",
	PrimaryPort:                           "primaryport",
	RequestBytesRate:                      "requestbytesrate",
	RequestsRate:                          "requestsrate",
	ResponseBytesRate:                     "responsebytesrate",
	ResponsesRate:                         "responsesrate",
	ServerEstablishedConnections:          "svrestablishedconn",
	ServerName:                            "servername",
	ServiceGroupMemberBindingOrder:        "svcgrpmemberbindingorder",
	ServiceOrder:                          "serviceorder",
	ServiceType:                           "servicetype",
	State:                                 "state",
	SurgeCount:                            "surgecount",
	ToleratingTimeToLastByteTransactions:  "toleratingttlbtransactions",
	TotalConnectionReassemblyQueueFlush:   "totalconnreassemblyqueueflush",
	TotalConnectionReassemblyQueueFlush75: "totalconnreassemblyqueue75",
	TotalRequestBytes:                     "totalrequestbytes",
	TotalRequests:                         "totalrequests",
	TotalResponseBytes:                    "totalresponsebytes",
	TotalResponses:                        "totalresponses",
	TotalServerTimeToLastByteTransactions: "totsvrttlbtransactions",
}

type ServiceGroupMember struct {
	AverageServerTimeToFirstByte          float64 `json:"avgsvrttfb,omitempty" nitro:"permission=readonly"`
	ClearStats                            string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentClientConnections              float64 `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	CurrentLoad                           float64 `json:"curload,omitempty" nitro:"permission=readonly"`
	CurrentReusePool                      float64 `json:"curreusepool,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections              float64 `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions float64 `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePacketsCount         float64 `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	Ip                                    string  `json:"ip,omitempty" nitro:"permission=readwrite"`
	MaxClients                            float64 `json:"maxclients,omitempty" nitro:"permission=readonly"`
	MaxHttpHeaderFieldLengthExceededCount float64 `json:"httpmaxhdrfldlenpkts,omitempty" nitro:"permission=readonly"`
	MaxOutOfOrderPacketsCount             float64 `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	Name                                  string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	Port                                  int     `json:"port,omitempty" nitro:"permission=readwrite"`
	PrimaryIpAddress                      string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PrimaryPort                           int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                      float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	RequestsRate                          float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                     float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	ResponsesRate                         float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	ServerEstablishedConnections          float64 `json:"svrestablishedconn,omitempty" nitro:"permission=readonly"`
	ServerName                            string  `json:"servername,omitempty" nitro:"permission=readwrite"`
	ServiceGroupMemberBindingOrder        float64 `json:"svcgrpmemberbindingorder,omitempty" nitro:"permission=readonly"`
	ServiceOrder                          float64 `json:"serviceorder,omitempty" nitro:"permission=readonly"`
	ServiceType                           string  `json:"servicetype,omitempty" nitro:"permission=readonly"`
	State                                 float64 `json:"state,omitempty" nitro:"permission=readonly"`
	SurgeCount                            float64 `json:"surgecount,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions  float64 `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush   float64 `json:"totalconnreassemblyqueueflush,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush75 float64 `json:"totalconnreassemblyqueue75,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                     float64 `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	TotalRequests                         float64 `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                    float64 `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	TotalResponses                        float64 `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	TotalServerTimeToLastByteTransactions float64 `json:"totsvrttlbtransactions,omitempty" nitro:"permission=readonly"`
}

func (r ServiceGroupMember) GetTypeName() string {
	return "servicegroupmember"
}
