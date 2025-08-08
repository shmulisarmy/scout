package apiglue

type StoreUpdateMessage struct {
	Type    string `json:"type" default:"store-update"`
	Path    string `json:"path"`
	NewData any    `json:"new_data"`
}

type StoreDeleteMessage struct {
	Type string `json:"type" default:"store-delete"`
	Path string `json:"path"`
}

type StoreAppendMessage struct {
	Type    string `json:"type" default:"store-append"`
	Path    string `json:"path"`
	NewData any    `json:"new_data"`
}

type StoreArrayFilterDeleteMessage struct {
	Type  string `json:"type" default:"store-array-filter-delete"`
	Path  string `json:"path"`
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type StoreArrayIndexDeleteMessage struct {
	Type  string `json:"type" default:"store-array-index-delete"`
	Path  string `json:"path"`
	Index int    `json:"index"`
}

// //
type MutableUpdateMessage struct {
	Type    string `json:"type" default:"mutable-update"`
	Key     string `json:"key"`
	Path    string `json:"path"`
	NewData any    `json:"new_data"`
}

type MutableDeleteMessage struct {
	Type string `json:"type" default:"mutable-delete"`
	Key  string `json:"key"`
	Path string `json:"path"`
}

type MutableAppendMessage struct {
	Type    string `json:"type" default:"mutable-append"`
	Key     string `json:"key"`
	Path    string `json:"path"`
	NewData any    `json:"new_data"`
}
