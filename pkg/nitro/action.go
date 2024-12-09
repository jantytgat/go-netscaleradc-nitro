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
