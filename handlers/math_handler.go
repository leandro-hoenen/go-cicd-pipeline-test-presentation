package handlers

import (
	"crypto/rand"
	"math/big"

	"github.com/gin-gonic/gin"
	"github.com/leandro-hoenen/go-cicd-pipeline-test/services"
)

type CompInput struct {
	ComponentOne int `json:"component_one" binding:"required"`
	ComponentTwo int `json:"component_two" binding:"required"`
}

type CompResult struct {
	InputComponents CompInput `json:"input_components"`
	Result          int       `json:"result"`
}

func HandleRandomAdd(c *gin.Context) {
	summandOne, errO := secureRandomInt(100)
	if errO != nil {
		c.JSON(400, gin.H{"error": "comp failed"})
		return
	}
	summandTwo, errT := secureRandomInt(100)
	if errT != nil {
		c.JSON(400, gin.H{"error": "comp failed"})
		return
	}

	sum := services.AddOperation(summandOne, summandTwo)

	input := CompInput{
		ComponentOne: summandOne,
		ComponentTwo: summandTwo,
	}

	r := CompResult{
		InputComponents: input,
		Result:          sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})
}

func HandleRandomSub(c *gin.Context) {
	minuend, errM := secureRandomInt(100)
	if errM != nil {
		c.JSON(400, gin.H{"error": "comp failed"})
		return
	}

	subtrahend, errS := secureRandomInt(100)
	if errS != nil {
		c.JSON(400, gin.H{"error": "comp failed"})
		return
	}

	difference := services.SubOperation(minuend, subtrahend)

	input := CompInput{
		ComponentOne: minuend,
		ComponentTwo: subtrahend,
	}

	r := CompResult{
		InputComponents: input,
		Result:          difference,
	}

	c.JSON(200, gin.H{
		"subResult": r,
	})
}

func HandleAdd(c *gin.Context) {
	var input CompInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	sum := services.AddOperation(input.ComponentOne, input.ComponentTwo)

	r := CompResult{
		InputComponents: input,
		Result:          sum,
	}

	c.JSON(200, gin.H{
		"addResult": r,
	})

}

func HandleSub(c *gin.Context) {
	var input CompInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON input",
		})
		return
	}

	difference := services.SubOperation(input.ComponentOne, input.ComponentTwo)

	r := CompResult{
		InputComponents: input,
		Result:          difference,
	}

	c.JSON(200, gin.H{
		"subResult": r,
	})
}

// With the help of ChatGPT
func secureRandomInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}
