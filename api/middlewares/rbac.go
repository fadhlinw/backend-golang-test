package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fadhlinw/backend-golang-test/domains"
	"github.com/fadhlinw/backend-golang-test/lib"
)

// RBACMiddleware middleware for rbac
type RBACMiddleware struct {
	service domains.RBACService
	logger  lib.Logger
}

// NewRBACMiddleware creates new rbac middleware
func NewRBACMiddleware(
	logger lib.Logger,
	service domains.RBACService,
) RBACMiddleware {
	return RBACMiddleware{
		service: service,
		logger:  logger,
	}
}

// Setup sets up rbac middleware
func (m RBACMiddleware) Setup() {}

// Handler handles middleware functionality
func (m RBACMiddleware) Handler(slug string, action string) gin.HandlerFunc {
	return func(c *gin.Context) {

		for key, values := range c.Request.Header {
			for _, value := range values {
				m.logger.Debugf("Header: %s = %s", key, value)
			}
		}

		m.logger.Info("Checking permission for role_id: ", c.Request.Header.Get("role_id"))
		roleId, err := strconv.Atoi(c.Request.Header.Get("role_id"))
		if err != nil {
			m.logger.Error("Error parsing role_id")
			abortErrorResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
			return
		}

		m.logger.Debug("role_id: ", roleId)
		result, _ := m.service.HasPermission(roleId, slug, action)
		if result {
			c.Next()
			return
		}

		abortErrorResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden))
		c.Next()
	}

}
