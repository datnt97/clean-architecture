package service

import (
	"github.com/phungvandat/clean-architecture/service/user"
)

// Service struct
type Service struct {
	UserService user.Service
}
