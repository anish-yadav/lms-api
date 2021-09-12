package auth

var RoutePermissionMap = map[string]string{
	"/health":                       "",
	"/users":                        "lms.user.create",
	"/users/change-password":        "lms.user.edit",
	"/users/reset-password":         "lms.user.edit",
	"/users/request-password-reset": "lms.user.edit",
	"/users/{id}":                   "lms.user.delete",
}
