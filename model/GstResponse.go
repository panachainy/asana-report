package model

type GstResponse struct {
	Data         []Gst `json:"data"`
	SumCompleted int
	SumTask      int
}

type Gst struct {
	ProjectName    string
	TotalCompleted int
	TotalTask      int
}
