package database

type Form struct {
	Service  string    `json:"service"`
	Name     string    `json:"name"`
	Info     string    `json:"info"`
	Contacts []Contact `json:"contacts"`
	FileInfo []File    `json:"file"`
}

type Contact struct {
	Kind string `json:"type"`
	Info string `json:"info"`
}

type File struct {
	Name string `json:"name"`
	Data Data   `json:"data"`
}

type Data struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}
