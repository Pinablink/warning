package warning

import (
	"errors"
	"fmt"
	"g2ssms/send"

	"github.com/Pinablink/mTools/mtime"
	uuid "github.com/satori/go.uuid"
)

//
type Warning struct {
	WarningUUID uuid.UUID
	urlService  string
	acao        send.SSAcao
	login       string
	token       string
	num         string
}

//
func NewWarning() *Warning {
	u1, _ := uuid.NewV4()
	return &Warning{WarningUUID: u1}
}

//
func (warning *Warning) ConfigSMSWarning(strUrlService, strLogin, strToken, strNum string) error {

	var allNotEmpty bool = (len(strUrlService) > 0 && len(strLogin) > 0 && len(strToken) > 0 && len(strNum) > 0)

	if allNotEmpty {

		warning.urlService = strUrlService
		warning.login = strLogin
		warning.token = strToken
		warning.num = strNum
		warning.acao = send.SendSms

	} else {
		return errors.New("Parametros de Configuração Incompleto")
	}

	return nil
}

//
func (warning *Warning) SendSMSMessage(refMessage string, errProcess bool) (string, error) {

	if len(refMessage) > 0 {
		var strStatus string = "OK"
		smsMsg := "%s %s %s %s"

		obSendSMS := &send.SSendSMS{}
		obSendSMS.UrlService = warning.urlService
		obSendSMS.Acao = warning.acao
		obSendSMS.Login = warning.login
		obSendSMS.Token = warning.token
		obSendSMS.Numero = warning.num

		if errProcess {
			strStatus = "NOT"
		}

		obSendSMS.Msg = send.SMsg{
			Msg: fmt.Sprintf(smsMsg, refMessage, mtime.DateTimeCurrentDefault(), strStatus, "OK"),
		}

		idResponse, err := obSendSMS.SendSMS()

		return idResponse, err
	}

	return "", nil
}
