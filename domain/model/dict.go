package model

type Dict struct {
	ID      int      `json:"id"`
	JpName  string   `json:"jp_name"`
	EngName string   `json:"eng_name"`
	Body    string   `json:"body"`
	Tags    []string `json:"tags"`
}
