/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package stat

type LbVserver struct {
	Name                                  string  `json:"name,omitempty" nitro:"permission=readwrite"`
	ClearState                            string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	SortBy                                string  `json:"sortby,omitempty" nitro:"permission=readwrite"`
	SortOrder                             string  `json:"sortorder,omitempty" nitro:"permission=readwrite"`
	CurrentClientConnections              string  `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	EstablishedConnections                string  `json:"establishedconn,omitempty" nitro:"permission=readonly"`
	TotalPacketsSent                      string  `json:"totalpktssent,omitempty" nitro:"permission=readonly"`
	LabelledConnections                   string  `json:"labelledconn,omitempty" nitro:"permission=readonly"`
	TotalHits                             string  `json:"tothits,omitempty" nitro:"permission=readonly"`
	TotalRequests                         string  `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	SpilloverThreshold                    string  `json:"sothreshold,omitempty" nitro:"permission=readonly"`
	CurrentMultipathTcpSubflowConnections string  `json:"cursubflowconn,omitempty" nitro:"permission=readonly"`
	ServerBusyErrorRate                   float64 `json:"svrbusyerrrate,omitempty" nitro:"permission=readonly"`
	SurgeCount                            string  `json:"surgecount,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                     float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponseDropped         string  `json:"invalidrequestresponsedropped,omitempty" nitro:"permission=readonly"`
	TotalResponses                        string  `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	CurrentPersistenceSessions            string  `json:"curpersistencesessions,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                      float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	TotalServerBusyError                  string  `json:"totalsvrbusyerr,omitempty" nitro:"permission=readonly"`
	AverageClientTimeToLastByte           string  `json:"avgcltttlb,omitempty" nitro:"permission=readonly"`
	Type                                  string  `json:"type,omitempty" nitro:"permission=readonly"`
	HitsRate                              float64 `json:"hitsrate,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueue75      string  `json:"totalconnreassemblyqueue75,omitempty" nitro:"permission=readonly"`
	RequestRetryCountExceeded             string  `json:"reqretrycountexceeded,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections              string  `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	TotalConnectionReassemblyQueueFlush   string  `json:"totalconnreassemblyqueueflush,omitempty" nitro:"permission=readonly"`
	ClientResponseTimeApdex               float64 `json:"cltresponsetimeapdex,omitempty" nitro:"permission=readonly"`
	TotalClientTimeToLastByteTransactions string  `json:"totcltttlbtransactions,omitempty" nitro:"permission=readonly"`
	CpuUsagePerMille                      string  `json:"cpuusagepm,omitempty" nitro:"permission=readonly"`
	PacketsReceivedRate                   float64 `json:"pktsrecvdrate,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                      string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	VserverSurgeCount                     string  `json:"vsvrsurgecount,omitempty" nitro:"permission=readonly"`
	PushLabel                             string  `json:"pushlabel,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions  string  `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	ResponsesRate                         float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	DeferredRequests                      string  `json:"deferredreq,omitempty" nitro:"permission=readonly"`
	TotalHttp2Requests                    string  `json:"totalh2requests,omitempty" nitro:"permission=readonly"`
	CurrentMultiPathTcpSessions           string  `json:"curmptcpsessions,omitempty" nitro:"permission=readonly"`
	TcpMaxOutOfOrderPackets               string  `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	TotalSpillovers                       string  `json:"totspillovers,omitempty" nitro:"permission=readonly"`
	ServiceSurgeCount                     string  `json:"svcsurgecount,omitempty" nitro:"permission=readonly"`
	Http2ResponsesRate                    float64 `json:"h2responsesrate,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                     string  `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponse                string  `json:"invalidrequestresponse,omitempty" nitro:"permission=readonly"`
	State                                 string  `json:"state,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePackets              string  `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	VserverLbHealth                       string  `json:"vslbhealth,omitempty" nitro:"permission=readonly"`
	DeferredRequestsRate                  float64 `json:"deferredreqrate,omitempty" nitro:"permission=readonly"`
	ActiveServices                        string  `json:"actsvcs,omitempty" nitro:"permission=readonly"`
	TotalPacketsReceived                  string  `json:"totalpktsrecvd,omitempty" nitro:"permission=readonly"`
	CurrentBackupPersistenceSessions      string  `json:"curbackuppersistencesessions,omitempty" nitro:"permission=readonly"`
	PacketsSentRate                       float64 `json:"pktssentrate,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions string  `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                    string  `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	TotalHttp2Responses                   string  `json:"totalh2responses,omitempty" nitro:"permission=readonly"`
	PrimaryPort                           int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	RequestsRate                          float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	RequestRetryCount                     string  `json:"reqretrycount,omitempty" nitro:"permission=readonly"`
	TotalVserverDownBackupHits            string  `json:"totvserverdownbackuphits,omitempty" nitro:"permission=readonly"`
	InactiveServices                      string  `json:"inactsvcs,omitempty" nitro:"permission=readonly"`
	Http2RequestsRate                     float64 `json:"h2requestsrate,omitempty" nitro:"permission=readonly"`
}

func (r LbVserver) GetTypeName() string {
	return "lbvserver"
}
