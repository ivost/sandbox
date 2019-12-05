package report

type Warning struct {
	Tag  string `json:",omitempty"`
	Text string
}

type Data struct {
	Warnings []Warning `json:",omitempty"`
	Error    string    `json:",omitempty"`
}
