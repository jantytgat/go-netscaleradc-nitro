package stat

var CsVserverFieldNames = struct {
	AverageClientTimeToLastByte           string
	ClearStats                            string
	ClientResponseTimeApdex               string
	CurrentBackupPersistenceSessions      string
	CurrentClientConnections              string
	CurrentMultipathTcpSessions           string
	CurrentMultiPathTcpSubflowConnections string
	CurrentPersistenceSessions            string
	CurrentServerConnections              string
	DeferredRequests                      string
	DeferredRequestsRate                  string
	EstablishedConnections                string
	FrustratingTimeToLastByteTransactions string
	HitsRate                              string
	HttpMaxHeaderSizePackets              string
	InvalidRequestResponse                string
	InvalidRequestResponseDropped         string
	LabeledConnections                    string
	Name                                  string
	PacketsReceivedRate                   string
	PacketsSentRate                       string
	PrimaryIpAddress                      string
	PrimaryPort                           string
	PushLabel                             string
	RequestBytesRate                      string
	RequestsRate                          string
	ResponseBytesRate                     string
	ResponseRate                          string
	SpilloverThreshold                    string
	State                                 string
	TcpMaxOutOfOrderPackets               string
	ToleratingTimeToLastByteTransactions  string
	TotalClientTimeToLastByteTransactions string
	TotalHits                             string
	TotalPacketsReceiver                  string
	TotalPacketsSent                      string
	TotalRequestBytes                     string
	TotalRequests                         string
	TotalResponseBytes                    string
	TotalResponses                        string
	TotalSpillovers                       string
	TotalVserverDownBackupHits            string
	Type                                  string
}{
	AverageClientTimeToLastByte:           "avgcltttlb",
	ClearStats:                            "clearstats",
	ClientResponseTimeApdex:               "cltresponsetimeapdex",
	CurrentBackupPersistenceSessions:      "curbackuppersistencesessions",
	CurrentClientConnections:              "curclntconnections",
	CurrentMultipathTcpSessions:           "curmptcpsessions",
	CurrentMultiPathTcpSubflowConnections: "cursubflowconn",
	CurrentPersistenceSessions:            "curpersistencesessions",
	CurrentServerConnections:              "cursrvrconnections",
	DeferredRequests:                      "deferredreq",
	DeferredRequestsRate:                  "deferredreqrate",
	EstablishedConnections:                "establishedconn",
	FrustratingTimeToLastByteTransactions: "frustratingttlbtransactions",
	HitsRate:                              "hitsrate",
	HttpMaxHeaderSizePackets:              "httpmaxhdrszpkts",
	InvalidRequestResponse:                "invalidrequestresponse",
	InvalidRequestResponseDropped:         "invalidrequestresponsedropped",
	LabeledConnections:                    "labelledconn",
	Name:                                  "name",
	PacketsReceivedRate:                   "pktsrecvdrate",
	PacketsSentRate:                       "pktssentrate",
	PrimaryIpAddress:                      "primaryipaddress",
	PrimaryPort:                           "primaryport",
	PushLabel:                             "pushlabel",
	RequestBytesRate:                      "requestbytesrate",
	RequestsRate:                          "requestsrate",
	ResponseBytesRate:                     "responsebytesrate",
	ResponseRate:                          "responsesrate",
	SpilloverThreshold:                    "sothreshold",
	State:                                 "state",
	TcpMaxOutOfOrderPackets:               "tcpmaxooopkts",
	ToleratingTimeToLastByteTransactions:  "toleratingttlbtransactions",
	TotalClientTimeToLastByteTransactions: "totcltttlbtransactions",
	TotalHits:                             "tothits",
	TotalPacketsReceiver:                  "totalpktsrecvd",
	TotalPacketsSent:                      "totalpktssent",
	TotalRequestBytes:                     "totalrequestbytes",
	TotalRequests:                         "totalrequests",
	TotalResponseBytes:                    "totalresponsebytes",
	TotalResponses:                        "totalresponses",
	TotalSpillovers:                       "totspillovers",
	TotalVserverDownBackupHits:            "totvserverdownbackuphits",
	Type:                                  "type",
}

type CsVserver struct {
	AverageClientTimeToLastByte           string  `json:"avgcltttlb,omitempty" nitro:"permission=readonly"`
	ClearStats                            string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	ClientResponseTimeApdex               float64 `json:"cltresponsetimeapdex,omitempty" nitro:"permission=readonly"`
	CurrentBackupPersistenceSessions      string  `json:"curbackuppersistencesessions,omitempty" nitro:"permission=readonly"`
	CurrentClientConnections              string  `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	CurrentMultipathTcpSessions           string  `json:"curmptcpsessions,omitempty" nitro:"permission=readonly"`
	CurrentMultiPathTcpSubflowConnections string  `json:"cursubflowconn,omitempty" nitro:"permission=readonly"`
	CurrentPersistenceSessions            string  `json:"curpersistencesessions,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections              string  `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	DeferredRequests                      string  `json:"deferredreq,omitempty" nitro:"permission=readonly"`
	DeferredRequestsRate                  float64 `json:"deferredreqrate,omitempty" nitro:"permission=readonly"`
	EstablishedConnections                string  `json:"establishedconn,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions string  `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	HitsRate                              float64 `json:"hitsrate,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePackets              string  `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponse                string  `json:"invalidrequestresponse,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponseDropped         string  `json:"invalidrequestresponsedropped,omitempty" nitro:"permission=readonly"`
	LabeledConnections                    string  `json:"labelledconn,omitempty" nitro:"permission=readonly"`
	Name                                  string  `json:"name,omitempty" nitro:"permission=readwrite"`
	PacketsReceivedRate                   float64 `json:"pktsrecvdrate,omitempty" nitro:"permission=readonly"`
	PacketsSentRate                       float64 `json:"pktssentrate,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                      string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PrimaryPort                           int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	PushLabel                             string  `json:"pushlabel,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                      float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	RequestsRate                          float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                     float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	ResponseRate                          float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	SpilloverThreshold                    string  `json:"sothreshold,omitempty" nitro:"permission=readonly"`
	State                                 string  `json:"state,omitempty" nitro:"permission=readonly"`
	TcpMaxOutOfOrderPackets               string  `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions  string  `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalClientTimeToLastByteTransactions string  `json:"totcltttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalHits                             string  `json:"tothits,omitempty" nitro:"permission=readonly"`
	TotalPacketsReceiver                  string  `json:"totalpktsrecvd,omitempty" nitro:"permission=readonly"`
	TotalPacketsSent                      string  `json:"totalpktssent,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                     string  `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	TotalRequests                         string  `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                    string  `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	TotalResponses                        string  `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	TotalSpillovers                       string  `json:"totspillovers,omitempty" nitro:"permission=readonly"`
	TotalVserverDownBackupHits            string  `json:"totvserverdownbackuphits,omitempty" nitro:"permission=readonly"`
	Type                                  string  `json:"type,omitempty" nitro:"permission=readonly"`
}

func (r CsVserver) GetTypeName() string {
	return "csvserver"
}
