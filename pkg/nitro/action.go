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

package nitro

var (
	ActionNone    = Action{""}
	ActionClear   = Action{"clear"}
	ActionCount   = Action{"count"}
	ActionCreate  = Action{"create"}
	ActionDiff    = Action{"diff"}
	ActionDisable = Action{"disable"}
	ActionEnable  = Action{"enable"}
	ActionForce   = Action{"force"}
	ActionKill    = Action{"kill"}
	ActionLink    = Action{"link"}
	ActionRename  = Action{"rename"}
	ActionRestore = Action{"restore"}
	ActionSave    = Action{"save"}
	ActionSync    = Action{"sync"}
	ActionUnlink  = Action{"unlink"}
	ActionUnset   = Action{"unset"}
	ActionUpdate  = Action{"update"}
)

type Action struct {
	string
}
