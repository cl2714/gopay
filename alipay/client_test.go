package alipay

import (
	"fmt"
	"os"
	"testing"

	"github.com/iGoogle-ink/gopay"
)

var (
	client          *Client
	appid           = "2016091200494382"
	aliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey      = "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
)

func TestMain(m *testing.M) {

	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client = NewClient(appid, privateKey, false)
	// 配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
	// SetReturnUrl("https://www.gopay.ink").
	// SetNotifyUrl("https://www.gopay.ink")

	// err := client.SetCertSnByPath("cert/appCertPublicKey.crt", "cert/alipayRootCert.crt", "cert/alipayCertPublicKey_RSA2.crt")
	// if err != nil {
	//	fmt.Println("SetCertSnByPath:", err)
	//	return
	// }

	os.Exit(m.Run())
}

func TestClient_TradePrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单")
	bm.Set("out_trade_no", gopay.GetRandomString(32))
	bm.Set("total_amount", "100")

	// 创建订单
	aliRsp, err := client.TradePrecreate(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.QrCode:", aliRsp.Response.QrCode)
	fmt.Println("aliRsp.OutTradeNo:", aliRsp.Response.OutTradeNo)
}

func TestClient_TradeCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "创建订单")
	bm.Set("buyer_id", "2088802095984694")
	bm.Set("out_trade_no", "GZ201901301040355709")
	bm.Set("total_amount", "0.01")

	// 创建订单
	aliRsp, err := client.TradeCreate(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.TradeNo:", aliRsp.Response.TradeNo)
}

func TestClient_TradeAppPay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "测试APP支付")
	bm.Set("out_trade_no", "GZ201901301040355706100469")
	bm.Set("total_amount", "1.00")

	// 手机APP支付参数请求
	payParam, err := client.TradeAppPay(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payParam:", payParam)
}

func TestClient_TradeCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 撤销支付订单
	aliRsp, err := client.TradeCancel(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_TradeClose(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 条码支付
	aliRsp, err := client.TradeClose(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_TradePay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "条码支付")
	bm.Set("scene", "bar_code")
	bm.Set("auth_code", "286248566432274952")
	bm.Set("out_trade_no", "GZ201909081743431443")
	bm.Set("total_amount", "0.01")
	bm.Set("timeout_express", "2m")

	// 条码支付
	aliRsp, err := client.TradePay(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		fmt.Println("err:::", err)
	}
	fmt.Println("同步返回验签：", ok)
}

func TestClient_TradeQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")

	// 查询订单
	aliRsp, err := client.TradeQuery(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_TradeWapPay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "手机网站测试支付")
	bm.Set("out_trade_no", "GZ201909081743431443")
	bm.Set("quit_url", "https://www.gopay.ink")
	bm.Set("total_amount", "100.00")
	bm.Set("product_code", "QUICK_WAP_WAY")

	// 手机网站支付请求
	payUrl, err := client.TradeWapPay(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestClient_TradePagePay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "网站测试支付")
	bm.Set("out_trade_no", "GZ201909081743431443")
	bm.Set("total_amount", "88.88")
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	// 电脑网站支付请求
	payUrl, err := client.TradePagePay(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestClient_TradeRefund(t *testing.T) {
	// 请求参数
	body := make(gopay.BodyMap)
	body.Set("out_trade_no", "GZ201909081743431443")
	body.Set("refund_amount", "5")
	body.Set("refund_reason", "测试退款")

	// 发起退款请求
	aliRsp, err := client.TradeRefund(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_TradePageRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")
	bm.Set("refund_amount", "5")
	bm.Set("out_request_no", gopay.GetRandomString(32))

	// 发起退款请求
	aliRsp, err := client.TradePageRefund(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_SystemOauthToken(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "3a06216ac8f84b8c93507bb9774bWX11")

	// 发起请求
	aliRsp, err := client.SystemOauthToken(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp:", aliRsp.Response.AccessToken)
	fmt.Println("aliRsp:", aliRsp.SignData)
}

func TestClient_TradeOrderSettle(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_request_no", "201907301518083384")
	bm.Set("trade_no", "2019072522001484690549776067")

	var listParams []OpenApiRoyaltyDetailInfoPojo
	listParams = append(listParams, OpenApiRoyaltyDetailInfoPojo{"transfer", "2088802095984694", "userId", "userId", "2088102363632794", "0.01", "分账给2088102363632794"})

	bm.Set("royalty_parameters", listParams)
	// fmt.Println("listParams:", bm.Get("royalty_parameters"))

	// 发起交易结算接口
	aliRsp, err := client.TradeOrderSettle(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "866185490c4e40efa9f71efea6766X02")

	// 发起请求
	aliRsp, err := client.OpenAuthTokenApp(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_TradeFastPayRefundQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", "GZ201909081743431443")
	bm.Set("out_request_no", "GZ201909081743431443")

	// 发起退款查询请求
	aliRsp, err := client.TradeFastPayRefundQuery(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_FundTransToaccountTransfer(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", gopay.GetRandomString(32))
	bm.Set("payee_type", "ALIPAY_LOGONID")
	bm.Set("payee_account", "otmdfd2378@sandbox.com")
	bm.Set("amount", "1000")
	bm.Set("payer_show_name", "发钱人名字")
	bm.Set("payee_real_name", "沙箱环境")
	bm.Set("remark", "转账测试")

	// 转账
	aliRsp, err := client.FundTransToaccountTransfer(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_UserCertifyOpenInit(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("outer_order_no", "ZGYD201809132323000001234")
	// 认证场景码：FACE：多因子人脸认证，CERT_PHOTO：多因子证照认证，CERT_PHOTO_FACE ：多因子证照和人脸认证，SMART_FACE：多因子快捷认证
	bm.Set("biz_code", "FACE")
	// 需要验证的身份信息参数，格式为json
	identity := make(map[string]string)
	identity["identity_type"] = "CERT_INFO"
	identity["cert_type"] = "IDENTITY_CARD"
	identity["cert_name"] = "张三"
	identity["cert_no"] = "310123199012301234"
	bm.Set("identity_param", identity)
	// 商户个性化配置，格式为json
	merchant := make(map[string]string)
	merchant["return_url"] = "https://www.gopay.ink"
	bm.Set("merchant_config", merchant)

	// 发起请求
	aliRsp, err := client.UserCertifyOpenInit(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_UserCertifyOpenCertify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "53827f9d085b3ce43938c6e5915b4729")

	// 发起请求
	certifyUrl, err := client.UserCertifyOpenCertify(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("certifyUrl:", certifyUrl)
}

func TestClient_UserCertifyOpenQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "OC201809253000000393900404029253")

	// 发起请求
	aliRsp, err := client.UserCertifyOpenQuery(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.Response.Passed:", aliRsp.Response.Passed)
}

func TestClient_UserInfoAuth(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 接口权限值，目前只支持auth_user和auth_base两个值。具体说明看文档介绍
	bm.Set("scopes", []string{"auth_user"})
	bm.Set("state", "init")

	// 发起请求
	aliRsp, err := client.UserInfoAuth(bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestClient_UserInfoShare(t *testing.T) {
	// 发起请求
	aliRsp, err := client.UserInfoShare()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSign(aliPayPublicKey, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		fmt.Println("VerifySign-err:", err)
		return
	}
	fmt.Println("ok:", ok)
}

func TestClient_ZhimaCreditScoreGet(t *testing.T) {
	// 请求参数
	body := make(gopay.BodyMap)
	body.Set("transaction_id", gopay.GetRandomString(48))
	body.Set("product_code", "w1010100100000000001")

	// 芝麻分
	aliRsp, err := client.ZhimaCreditScoreGet(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

// =================================================

func TestSyncVerifySign(t *testing.T) {
	signData := `{"code":"10000","msg":"Success","buyer_logon_id":"854***@qq.com","buyer_pay_amount":"0.01","buyer_user_id":"2088102363632794","fund_bill_list":[{"amount":"0.01","fund_channel":"PCREDIT"}],"gmt_payment":"2019-08-29 20:14:05","invoice_amount":"0.01","out_trade_no":"GZ201901301040361012","point_amount":"0.00","receipt_amount":"0.01","total_amount":"0.01","trade_no":"2019082922001432790585537960"}`
	sign := "bk3SzX0CZRI811IJioS2XKQHcgMixUT8mYyGQj+vcOAQas7GIYi6LpykqqSc3m7+yvqoG0TdX/c2JjYnpw/J53JxtC2IC4vsLuIPIgghVo5qafsfSxEJ22w20RZDatI2dYqFVcj8Jp+4aesQ8zMMNw7cX9NLyk7kw3DecYeyQp+zrZMueZPqLh88Z+54G+e6QuSU++0ouqQVd4PkpPqy6YI+8MdMUX4Ve0jOQxMmYH8BC6n5ZsTH/uEaLEtzYVZdSw/xdSQ7K1SH73aEH8XbRYx6rL7RkKksrdvhezX+ThDjQ+fTWjvNFrGcg3fmqXRy2elvoalu+BQmqlkWWjEJYA=="
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"
	pKey := FormatPublicKey(aliPayPublicKey)
	err := verifySign(signData, sign, "RSA2", pKey)
	if err != nil {
		fmt.Println("err:", err)
	}
}

func TestVerifySign(t *testing.T) {
	// 测试，假数据，无法验签通过

	publicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	//bm := make(gopay.BodyMap)
	//bm.Set("sign", "f19WZ3rko3cVpSG3uEEJF0eb4DuZVLt4/GXnNw9qg8iHUsJLkav0V91R5SSTDhW5lgkn3Xhq7TkFRJiDXdVXMu3XUlsONArp3Iu4tXagYJWt9jbcnc2/l29VYDXPLNcs7BXEWFEaCZLutQY2U82AumEwSc1XBKtsLC4mVX3M3f/ExFQHWklJEBHArYBGe4535uFRlsT2fk6WVuX8CuYZatCrVF1o02xMS5aD29eICPkmin/h87OcTbE1syktyCU1WVgcypagUdGGPTF0SVDFf7FRov7+w7fiCGGGL10tNlK/MLzcewtN2dyGF6RLUX3m+HQ7sNEk2wylRXLNUFig==")
	//bm.Set("seller_email", "imonkey@100tal.com")
	//bm.Set("sign_type", "RSA2")
	//bm.Set("total_amount", "0.02")
	//bm.Set("buyer_id", "2088812847201551")
	//bm.Set("invoice_amount", "0.02")
	//bm.Set("fund_bill_list", `[{"amount":"0.02","fundChannel":"PCREDIT"}]`)
	//bm.Set("trade_no", "2020010222001401551430614892")
	//bm.Set("receipt_amount", "0.02")
	//bm.Set("buyer_pay_amount", "0.02")
	//bm.Set("notify_time", "2020-01-02 16:18:21")
	//bm.Set("subject", "商品")
	//bm.Set("auth_app_id", "2015102700040153")
	//bm.Set("charset", "utf-8")
	//bm.Set("point_amount", "0.00")
	//bm.Set("notify_type", "trade_status_sync")
	//bm.Set("out_trade_no", "1086209247658383466")
	//bm.Set("gmt_payment", "2020-01-02 16:18:21")
	//bm.Set("trade_status", "TRADE_SUCCESS")
	//bm.Set("version", "1.0")
	//bm.Set("buyer_logon_id", "185****2920")
	//bm.Set("gmt_create", "2020-01-02 16:18:21")
	//bm.Set("app_id", "2015102700040153")
	//bm.Set("seller_id", "2088631240818980")
	//bm.Set("notify_id", "2020010200222161821001551453140885")

	req := new(NotifyRequest)
	req.GmtCreate = "2020-01-02 16:18:21"
	req.Charset = "utf-8"
	req.SellerEmail = "imonkey@100tal.com"
	req.Subject = "商品"
	req.Sign = "f19WZ3rko3cVpSG3uEEJF0eb4DuZVLt4/GXnNw9qg8iHUsJLkav0V91R5SSTDhW5lgkn3Xhq7TkFRJiDXdVXMu3XUlsONArp3Iu4tXagYJWt9jbcnc2/l29VYDXPLNcs7BXEWFEaCZLutQY2U82AumEwSc1XBKtsLC4mVX3M3f/ExFQHWklJEBHArYBGe4535uFRlsT2fk6WVuX8CuYZatCrVF1o02xMS5aD29eICPkmin/h87OcTbE1syktyCU1WVgcypagUdGGPTF0SVDFf7FRov7+w7fiCGGGL10tNlK/MLzcewtN2dyGF6RLUX3m+HQ7sNEk2wylRXLNUFig=="
	req.BuyerId = "2088812847201551"
	req.InvoiceAmount = "0.02"
	req.NotifyId = "2020010200222161821001551453140885"
	infos := []*FundBillListInfo{&FundBillListInfo{
		Amount:      "0.02",
		FundChannel: "PCREDIT",
	}}
	req.FundBillList = infos
	req.NotifyType = "trade_status_sync"
	req.TradeStatus = "TRADE_SUCCESS"
	req.ReceiptAmount = "0.02"
	req.AppId = "2015102700040153"
	req.BuyerPayAmount = "0.02"
	req.SignType = "RSA2"
	req.SellerId = "2088631240818980"
	req.GmtPayment = "2020-01-02 16:18:21"
	req.NotifyTime = "2020-01-02 16:18:21"
	req.Version = "1.0"
	req.OutTradeNo = "1086209247658383466"
	req.TotalAmount = "0.02"
	req.TradeNo = "2020010222001401551430614892"
	req.AuthAppId = "2015102700040153"
	req.BuyerLogonId = "185****2920"
	req.PointAmount = "0.00"

	ok, err := VerifySign(publicKey, req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("OK:", ok)
}

func TestVerifySignWithCert(t *testing.T) {
	// 测试，假数据，无法验签通过
	bm := make(gopay.BodyMap)
	bm.Set("sign", "kPbQIjX+xQc8F0/A6/AocELIjhhZnGbcBN6G4MM/HmfWL4ZiHM6fWl5NQhzXJusaklZ1LFuMo+lHQUELAYeugH8LYFvxnNajOvZhuxNFbN2LhF0l/KL8ANtj8oyPM4NN7Qft2kWJTDJUpQOzCzNnV9hDxh5AaT9FPqRS6ZKxnzM=")
	bm.Set("sign_type", "RSA2")
	bm.Set("total_amount", "2.00")
	bm.Set("buyer_id", "2088102116773037")
	bm.Set("body", "大乐透2.1")
	bm.Set("trade_no", "2016071921001003030200089909")
	bm.Set("refund_fee", "0.00")
	bm.Set("notify_time", "2016-07-19 14:10:49")
	bm.Set("subject", "大乐透2.1")
	bm.Set("charset", "utf-8")
	bm.Set("notify_type", "trade_status_sync")
	bm.Set("out_trade_no", "0719141034-6418")
	bm.Set("gmt_close", "2016-07-19 14:10:46")
	bm.Set("gmt_payment", "2016-07-19 14:10:47")
	bm.Set("trade_status", "TRADE_SUCCESS")
	bm.Set("version", "1.0")
	bm.Set("gmt_create", "2016-07-19 14:10:44")
	bm.Set("app_id", "2015102700040153")
	bm.Set("seller_id", "2088102119685838")
	bm.Set("notify_id", "4a91b7a78a503640467525113fb7d8bg8e")

	ok, err := VerifySignWithCert("/cert/alipayCertPublicKey_RSA2.crt", bm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("OK:", ok)
}

func TestGetCertSN(t *testing.T) {
	sn, err := GetCertSN("cert/alipayCertPublicKey_RSA2.crt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 04afd423ea5bd6f5c5482854ed73278c
	fmt.Println("alipayCertPublicKey_RSA2:", sn)

	sn, err = GetCertSN("cert/appCertPublicKey.crt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 4498aaa8ab0c8986c15c41b36186db7d
	fmt.Println("appCertPublicKey:", sn)

	sn, err = GetRootCertSN("cert/alipayRootCert.crt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6
	fmt.Println("alipay_root_cert_sn:", sn)
}

func TestDecryptOpenDataToBodyMap(t *testing.T) {
	data := "MkvuiIZsGOC8S038cu/JIpoRKnF+ZFjoIRGf5d/K4+ctYjCtb/eEkwgrdB5TeH/93bxff1Ylb+SE+UGStlpvcg=="
	key := "TDftre9FpItr46e9BVNJcw=="
	bm, err := DecryptOpenDataToBodyMap(data, key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("bm:", bm)
}
