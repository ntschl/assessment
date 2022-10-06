package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/questions")    // all questions
	r.GET("/question")     // random question
	r.GET("/question/:id") // specific Q
	r.Run("0.0.0.0:8080")
}

type question struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

var questionsMap = map[string]question{
	"1152881b-221d-437d-8e8d-aae7bbb24abc": {"1152881b-221d-437d-8e8d-aae7bbb24abc", "What is your favorite color and why?"},
	"15240739-1ad9-43ea-bf5a-66dce200dced": {"15240739-1ad9-43ea-bf5a-66dce200dced", "What is your dream vacation like?"},
	"266ab3b0-cedf-4450-b11f-2ae81107636d": {"266ab3b0-cedf-4450-b11f-2ae81107636d", "Who is your biggest role model and why?"},
	"b0e5e096-76cc-475f-b523-8820b4aa5c00": {"b0e5e096-76cc-475f-b523-8820b4aa5c00", "Where did you see yourself in 10 years when you were 16?"},
	"b5266b38-1992-46a0-9bc9-defcd20ee7d5": {"b5266b38-1992-46a0-9bc9-defcd20ee7d5", "Why did you choose a career in tech?"},
}
