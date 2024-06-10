package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"path/filepath"
	"testing"
	"worframe/share/constant"
	"worframe/share/utils"
)

func casbinInit(modelName string, policyName string) (*casbin.Enforcer, error) {
	workdir, err := utils.FindWorkDir()
	if err != nil {
		panic(err)
	}
	ms := constant.CasbinModel[modelName]
	m, err := model.NewModelFromString(ms)

	strategyPath := filepath.Join(workdir, "test", "casbin", "data", "casbin_test")
	policyFile := filepath.Join(strategyPath, policyName+"_policy.csv")
	return casbin.NewEnforcer(m, policyFile)
}

type cas struct {
	*casbin.Enforcer
}

func (cas) casbinTest(t testing.T, sub, obj, act string) {

}
