package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func RegisterCustomValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("goal_time_format", GoalTimeValidator)
	}
}

func GoalTimeValidator(fl validator.FieldLevel) bool {
	time := fl.Field().String()
	match, _ := regexp.MatchString(`^\d{1,2}:[0-5]\d$`, time)
	return match
}
