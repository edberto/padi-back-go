package condition

type ConditionDetailVM struct {
	Label       string `json:"label"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
	Solution    string `json:"solution"`
	Prevention  string `json:"prevention"`
}
