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

type HaNode struct {
	ClearStats                    string  `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	CurrentStatus                 string  `json:"hacurstatus,omitempty" nitro:"permission=readonly"`
	CurrentState                  string  `json:"hacurstate,omitempty" nitro:"permission=readonly"`
	CurrentMasterState            string  `json:"hacurmasterstate,omitempty" nitro:"permission=readonly"`
	TransitionTime                string  `json:"transtime,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPackets      string  `json:"hatotpktrx,omitempty" nitro:"permission=readonly"`
	HeartbeatReceivedPacketRate   float64 `json:"hapktrxrate,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPackets          string  `json:"hatotpkttx,omitempty" nitro:"permission=readonly"`
	HeartbeatSentPacketRate       float64 `json:"hapkttxrate,omitempty" nitro:"permission=readonly"`
	PropagationTimeoutCounter     string  `json:"haerrproptimeout,omitempty" nitro:"permission=readonly"`
	SynchronizationFailureCounter string  `json:"haerrsyncfailure,omitempty" nitro:"permission=readonly"`
}

func (r HaNode) GetTypeName() string {
	return "hanode"
}
