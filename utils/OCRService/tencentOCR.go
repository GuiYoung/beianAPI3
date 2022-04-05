package OCRService

import (
	"beianAPI/utils"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// 因为腾讯限制额度，所以后期拉了个开源项目在自己服务器上部署了一下OCR服务，所以这段代码没啥用

func TencentOCRService(base64 string) (string, error) {
	config := &utils.Conf.TOCR

	credential := common.NewCredential(
		config.SecretID,
		config.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.EndPoint
	client, _ := ocr.NewClient(credential, config.Region, cpf)

	request := ocr.NewGeneralHandwritingOCRRequest()

	request.ImageBase64 = common.StringPtr(base64)

	response, err := client.GeneralHandwritingOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return "", err
	}
	if err != nil {
		panic(err)
	}
	return *response.Response.TextDetections[0].DetectedText, nil
}
