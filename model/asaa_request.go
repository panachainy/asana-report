package model

/*
Example use AsaaRequest without struct.

----

map[string]interface{}{
	"data": map[string]interface{}{
		"assignee": assigneeId,
	},
},

----
*/
type AsaaRequest struct {
	Data struct {
		Assignee string `json:"assignee"`
	} `json:"data"`
}
