package model

type AstResponse struct {
	Data                []Ast `json:"data"`
	SumCompleted        int
	SumTask             int
	SumSubTask          int
	SumSubTaskCompleted int
}

type Ast struct {
	ProjectName           string
	TotalCompleted        int
	TotalTask             int
	TotalSubTask          int
	TotalSubTaskCompleted int
}
