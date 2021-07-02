package model

type AstResponse struct {
	Data         []Ast `json:"data"`
	SumCompleted int
	SumTask      int
}

type Ast struct {
	ProjectName    string
	TotalCompleted int
	TotalTask      int
}
