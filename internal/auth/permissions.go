package auth

var RoutePermissionMap = map[string]string{
	"/health":               "",
	"/users":                "lms.user.create",
	"/users/password-reset": "lms.user.edit",
}
