package auth

import "net/http"

var RoutePermissionMap = map[string]string{
	"/health":                       "",
	"/users":                        "lms.user.create",
	"/users/change-password":        "lms.user.edit",
	"/users/reset-password":         "lms.user.edit",
	"/users/request-password-reset": "lms.user.edit",
	"/users/{id}":                   "lms.user.delete",
	"/users/me":                     "",
}

var PermissionMap = map[string]map[string]string{
	http.MethodGet: {
		"/health":   "",
		"/users":    "lms.user.create",
		"/users/me": "lms.user.read",
		"/class":    "lms.class.read",
	},
	http.MethodPost: {
		"/users":                 "lms.user.create",
		"/users/change-password": "lms.user.edit",
		"/users/reset-password":  "lms.user.edit",
	},
	http.MethodDelete: {
		"/users": "lms.user.delete",
	},
}
