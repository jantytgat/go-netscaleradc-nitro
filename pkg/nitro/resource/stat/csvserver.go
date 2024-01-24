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

type CsVserver struct {
	Name                                  string  `json:"name,omitempty" nitro:"permission=readwrite"`
	ClearStats                            string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentClientConnections              string  `json:"curclntconnections,omitempty" nitro:"permission=readonly"`
	EstablishedConnections                string  `json:"establishedconn,omitempty" nitro:"permission=readonly"`
	TotalPacketsSent                      string  `json:"totalpktssent,omitempty" nitro:"permission=readonly"`
	LabeledConnections                    string  `json:"labelledconn,omitempty" nitro:"permission=readonly"`
	TotalHits                             string  `json:"tothits,omitempty" nitro:"permission=readonly"`
	TotalRequests                         string  `json:"totalrequests,omitempty" nitro:"permission=readonly"`
	SpilloverThreshold                    string  `json:"sothreshold,omitempty" nitro:"permission=readonly"`
	CurrentMultiPathTcpSubflowConnections string  `json:"cursubflowconn,omitempty" nitro:"permission=readonly"`
	ResponseBytesRate                     float64 `json:"responsebytesrate,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponseDropped         string  `json:"invalidrequestresponsedropped,omitempty" nitro:"permission=readonly"`
	TotalResponses                        string  `json:"totalresponses,omitempty" nitro:"permission=readonly"`
	CurrentPersistenceSessions            string  `json:"curpersistencesessions,omitempty" nitro:"permission=readonly"`
	RequestBytesRate                      float64 `json:"requestbytesrate,omitempty" nitro:"permission=readonly"`
	AverageClientTimeToLastByte           string  `json:"avgcltttlb,omitempty" nitro:"permission=readonly"`
	Type                                  string  `json:"type,omitempty" nitro:"permission=readonly"`
	HitsRate                              float64 `json:"hitsrate,omitempty" nitro:"permission=readonly"`
	CurrentServerConnections              string  `json:"cursrvrconnections,omitempty" nitro:"permission=readonly"`
	ClientResponseTimeApdex               float64 `json:"cltresponsetimeapdex,omitempty" nitro:"permission=readonly"`
	TotalClientTimeToLastByteTransactions string  `json:"totcltttlbtransactions,omitempty" nitro:"permission=readonly"`
	PacketsReceivedRate                   float64 `json:"pktsrecvdrate,omitempty" nitro:"permission=readonly"`
	PrimaryIpAddress                      string  `json:"primaryipaddress,omitempty" nitro:"permission=readonly"`
	PushLabel                             string  `json:"pushlabel,omitempty" nitro:"permission=readonly"`
	ToleratingTimeToLastByteTransactions  string  `json:"toleratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	ResponseRate                          float64 `json:"responsesrate,omitempty" nitro:"permission=readonly"`
	DeferredRequests                      string  `json:"deferredreq,omitempty" nitro:"permission=readonly"`
	CurrentMultipathTcpSessions           string  `json:"curmptcpsessions,omitempty" nitro:"permission=readonly"`
	TcpMaxOutOfOrderPackets               string  `json:"tcpmaxooopkts,omitempty" nitro:"permission=readonly"`
	TotalSpillovers                       string  `json:"totspillovers,omitempty" nitro:"permission=readonly"`
	TotalRequestBytes                     string  `json:"totalrequestbytes,omitempty" nitro:"permission=readonly"`
	InvalidRequestResponse                string  `json:"invalidrequestresponse,omitempty" nitro:"permission=readonly"`
	State                                 string  `json:"state,omitempty" nitro:"permission=readonly"`
	HttpMaxHeaderSizePackets              string  `json:"httpmaxhdrszpkts,omitempty" nitro:"permission=readonly"`
	DeferredRequestsRate                  float64 `json:"deferredreqrate,omitempty" nitro:"permission=readonly"`
	TotalPacketsReceiver                  string  `json:"totalpktsrecvd,omitempty" nitro:"permission=readonly"`
	CurrentBackupPersistenceSessions      string  `json:"curbackuppersistencesessions,omitempty" nitro:"permission=readonly"`
	PacketsSentRate                       float64 `json:"pktssentrate,omitempty" nitro:"permission=readonly"`
	FrustratingTimeToLastByteTransactions string  `json:"frustratingttlbtransactions,omitempty" nitro:"permission=readonly"`
	TotalResponseBytes                    string  `json:"totalresponsebytes,omitempty" nitro:"permission=readonly"`
	PrimaryPort                           int     `json:"primaryport,omitempty" nitro:"permission=readonly"`
	RequestsRate                          float64 `json:"requestsrate,omitempty" nitro:"permission=readonly"`
	TotalVserverDownBackupHits            string  `json:"totvserverdownbackuphits,omitempty" nitro:"permission=readonly"`
}

func (r CsVserver) GetTypeName() string {
	return "csvserver"
}
