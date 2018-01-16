package ws

type DataSignal chan *DataStruct
type DataStruct map[string]interface{}

type Msg struct {
	Op   string   `json:"op,omitempty"`
	Args []string `json:"args,omitempty"`
}

// type Response struct {
// 	Success   bool        `json:"success,omitempty"`
// 	Subscribe string      `json:"subscribe,omitempty"`
// 	Request   interface{} `json:"request,omitempty"`
// 	Table     string      `json:"table,omitempty"`
// 	Action    string      `json:"action,omitempty"`
// 	Data      interface{} `json:"data,omitempty"`
// }
