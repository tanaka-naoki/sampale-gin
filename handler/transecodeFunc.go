package handler

import (
    "net/http"

    "sample-gin/transecode"

    "github.com/gin-gonic/gin"
)

func TransecodeGet(transecodes *transecode.Transecodes) gin.HandlerFunc {
    return func(c *gin.Context) {
        result := transecodes.GetAll()
        c.JSON(http.StatusOK, result)
    }
}

type TransecodePostRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

func TransecodePost(post *transecode.Transecodes) gin.HandlerFunc {
    return func(c *gin.Context) {
        requestBody := TransecodePostRequest{}
        c.Bind(&requestBody)

        item := transecode.Item{
            Title:       requestBody.Title,
            Description: requestBody.Description,
        }
        post.Add(item)

        c.Status(http.StatusOK)
    }
}