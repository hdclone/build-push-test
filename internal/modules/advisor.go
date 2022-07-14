package modules

import (
	"broadcaster/internal/advisor"
	"net/http"
	"time"
)

func Advisor() advisor.Interface {
	return Register("advisor", func(s string) (Module, error) {
		advisorConfig := Config().Advisor
		if advisorConfig.Mock != nil {
			return advisor.NewAdvisorMock(advisorConfig.Mock), nil
		}
		return advisor.NewAdvisor(advisorConfig.Url, &http.Client{
			Timeout: time.Second * 5,
		}), nil
	}).(advisor.Interface)
}
