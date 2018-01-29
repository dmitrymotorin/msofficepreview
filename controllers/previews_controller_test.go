package controllers

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

func TestCreatePreviewsController(t *testing.T) {
  controller := CreatePreviewsController(previewsGenerator)

  assert.NotNil(t, controller)
}

func TestCreateMethod(t *testing.T) {
  controller := CreatePreviewsController(previewGenerator)
  w := httptest.NewRecorder()
  ctx, eng := gin.CreateTestContext(w)

  controller.Create(ctx)

  assert.Equal(t, w.Result().Status, 200)
}

