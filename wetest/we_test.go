/**
  @author: xuejian.he
  @since: 2023/3/30
  @desc: //TODO
**/
package wetest

import (
	"github.com/g8rswimmer/go-sfdc"
	"github.com/g8rswimmer/go-sfdc/credentials"
	"github.com/g8rswimmer/go-sfdc/session"
	"github.com/g8rswimmer/go-sfdc/sobject"
	"net/http"
	"testing"
)

func TestNewPasswordLogin(t *testing.T) {
	//loginUri     := "https://login.salesforce.com/services/oauth2/token"
	testLoginUri := "https://test.salesforce.com"
	cred := credentials.PasswordCredentials{
		URL:          testLoginUri,
		Username:     "yuan.jiang@anyue.net.fullbox",
		Password:     "EasyTM_2018zxZRQgByvV1P5hzPUsqrnM8si",
		ClientID:     "3MVG9Y6d_Btp4xp6bypxI7YjObo.F2N9b4EwQt4GgCkh_lQmYPCFW_yp_hKeOAjERqpMt4.GUZYgJOwUVhH1K",
		ClientSecret: "1216742033804660031",
	}

	creds, err := credentials.NewPasswordCredentials(cred)
	if err != nil {
		t.Logf("登录错误：%v", err)
		return
	}

	config := sfdc.Configuration{
		Credentials: creds,
		Client:      http.DefaultClient,
		Version:     56,
	}

	sessionx, err := session.Open(config)
	if err != nil {
		t.Logf("登录错误：%v", err)
		return
	}

	resource, err := sobject.NewResources(sessionx)
	if err != nil {
		t.Logf("登录错误：%v", err)
		return
	}

	records, err := resource.Query(&mockQuery{
		sobject: "Account",
		id:      "00190000021BQJY",
	})

	t.Logf("查询错误：%v", err)
	t.Logf("查询结果：%v", records)
}

type mockQuery struct {
	sobject string
	id      string
	fields  []string
}

func (mock *mockQuery) SObject() string {
	return mock.sobject
}
func (mock *mockQuery) ID() string {
	return mock.id
}
func (mock *mockQuery) Fields() []string {
	return mock.fields
}
