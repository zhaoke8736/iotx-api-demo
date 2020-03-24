package rule

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/iot"
)

//StartRuleSample 查询指定规则的详细信息，RuleId在创建Rule时返回
func StartRuleSample() {
	client, err := iot.NewClientWithAccessKey(endpoint, accessKeyID, accessKeySecret)

	request := iot.CreateStartRuleRequest()
	request.AcceptFormat = "json"
	request.RuleId = "427058"

	response, err := client.StartRule(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}