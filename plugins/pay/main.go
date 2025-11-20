package main

import (
    "bytes"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/base64"
    "encoding/hex"
    "encoding/pem"
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "sort"
    "strings"
    "time"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func create(params map[string]interface{}) (interface{}, error) {
    amount := fmt.Sprint(params["amount"])
    currency := fmt.Sprint(params["currency"])
    provider := fmt.Sprint(params["provider"])
    orderNo := fmt.Sprint(params["order_no"]) // optional pre-generated
    if amount == "" || provider == "" { return nil, fmt.Errorf("missing fields") }
    buf := make([]byte, 16); _, _ = rand.Read(buf)
    txn := hex.EncodeToString(buf)
    if orderNo == "" { orderNo = fmt.Sprintf("o_%d_%s", time.Now().Unix(), txn[:8]) }
    payURL := fmt.Sprintf("https://pay.example.com/%s/checkout?order_no=%s&amount=%s&currency=%s&txn=%s", provider, orderNo, amount, currency, txn)
    return map[string]interface{}{ "order_no": orderNo, "txn_id": txn, "pay_url": payURL }, nil
}

func alipayCreate(params map[string]interface{}) (interface{}, error) {
    appID := os.Getenv("ALIPAY_APP_ID")
    privateKeyPEM := os.Getenv("ALIPAY_PRIVATE_KEY")
    gateway := os.Getenv("ALIPAY_GATEWAY")
    subject := fmt.Sprint(params["subject"]) ; amount := fmt.Sprint(params["amount"]) ; outTradeNo := fmt.Sprint(params["order_no"])
    if appID == "" || privateKeyPEM == "" || gateway == "" {
        p := map[string]interface{}{"provider":"alipay","amount":amount,"order_no":outTradeNo}
        return create(p)
    }
    if outTradeNo == "" { outTradeNo = fmt.Sprintf("o_%d", time.Now().UnixNano()) }
    biz := fmt.Sprintf("{\"subject\":\"%s\",\"out_trade_no\":\"%s\",\"total_amount\":\"%s\",\"product_code\":\"FAST_INSTANT_TRADE_PAY\"}", subject, outTradeNo, amount)
    qs := map[string]string{
        "app_id": appID,
        "method": "alipay.trade.page.pay",
        "format": "JSON",
        "return_url": "",
        "charset": "utf-8",
        "sign_type": "RSA2",
        "timestamp": time.Now().Format("2006-01-02 15:04:05"),
        "version": "1.0",
        "biz_content": biz,
    }
    sig, err := signRSA256(privateKeyPEM, canonicalize(qs))
    if err != nil { return nil, err }
    qs["sign"] = sig
    val := url.Values{}
    for k, v := range qs { val.Set(k, v) }
    payURL := gateway + "?" + val.Encode()
    return map[string]interface{}{ "order_no": outTradeNo, "pay_url": payURL }, nil
}

func wechatCreate(params map[string]interface{}) (interface{}, error) {
    mchID := os.Getenv("WECHAT_MCH_ID")
    appID := os.Getenv("WECHAT_APP_ID")
    serial := os.Getenv("WECHAT_CERT_SERIAL")
    privateKeyPEM := os.Getenv("WECHAT_PRIVATE_KEY")
    gateway := os.Getenv("WECHAT_GATEWAY")
    desc := fmt.Sprint(params["description"]) ; amount := fmt.Sprint(params["amount"]) ; outTradeNo := fmt.Sprint(params["order_no"]) ; notifyURL := fmt.Sprint(params["notify_url"])
    if mchID=="" || appID=="" || serial=="" || privateKeyPEM=="" || gateway=="" {
        p := map[string]interface{}{"provider":"wechat","amount":amount,"order_no":outTradeNo,"currency":""}
        return create(p)
    }
    if outTradeNo == "" { outTradeNo = fmt.Sprintf("o_%d", time.Now().UnixNano()) }
    body := fmt.Sprintf("{\"appid\":\"%s\",\"mchid\":\"%s\",\"description\":\"%s\",\"out_trade_no\":\"%s\",\"notify_url\":\"%s\",\"amount\":{\"total\":%s}}", appID, mchID, desc, outTradeNo, notifyURL, amount)
    path := "/v3/pay/transactions/native"
    ts := fmt.Sprintf("%d", time.Now().Unix())
    nonce := randHex(16)
    msg := strings.Join([]string{"POST", path, ts, nonce, body}, "\n") + "\n"
    sig, err := signRSA256(privateKeyPEM, msg)
    if err != nil { return nil, err }
    auth := fmt.Sprintf("WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",serial_no=\"%s\",timestamp=\"%s\",nonce_str=\"%s\",signature=\"%s\"", mchID, serial, ts, nonce, sig)
    req, _ := http.NewRequest("POST", gateway+path, bytes.NewBufferString(body))
    req.Header.Set("Authorization", auth)
    req.Header.Set("Content-Type", "application/json")
    cli := &http.Client{ Timeout: 5 * time.Second }
    resp, err := cli.Do(req)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    if resp.StatusCode != 200 && resp.StatusCode != 201 { return nil, fmt.Errorf("wechat create failed: %s", resp.Status) }
    buf := new(bytes.Buffer); _, _ = buf.ReadFrom(resp.Body)
    s := buf.String()
    return map[string]interface{}{ "order_no": outTradeNo, "provider": "wechat", "raw": s, "pay_url": s }, nil
}

func alipayVerify(params map[string]interface{}) (interface{}, error) {
    pubPEM := os.Getenv("ALIPAY_PUBLIC_KEY")
    payload := fmt.Sprint(params["payload"]) ; sig := fmt.Sprint(params["sign"]); if sig=="" { sig = fmt.Sprint(params["sig"]) }
    if pubPEM=="" { return map[string]interface{}{"ok": true}, nil }
    if payload=="" || sig=="" { return nil, fmt.Errorf("unauthorized") }
    ok := verifyRSA256(pubPEM, payload, sig)
    if !ok { return nil, fmt.Errorf("unauthorized") }
    return map[string]interface{}{ "ok": true }, nil
}

func wechatVerify(params map[string]interface{}) (interface{}, error) {
    pubPEM := os.Getenv("WECHAT_PUBLIC_KEY")
    ts := fmt.Sprint(params["timestamp"]) ; nonce := fmt.Sprint(params["nonce"]) ; body := fmt.Sprint(params["payload"]) ; sig := fmt.Sprint(params["signature"]) ; if sig=="" { sig = fmt.Sprint(params["sig"]) }
    if pubPEM=="" { return map[string]interface{}{"ok": true}, nil }
    if ts=="" || nonce=="" || sig=="" { return nil, fmt.Errorf("unauthorized") }
    msg := strings.Join([]string{ts, nonce, body}, "\n") + "\n"
    ok := verifyRSA256(pubPEM, msg, sig)
    if !ok { return nil, fmt.Errorf("unauthorized") }
    return map[string]interface{}{ "ok": true }, nil
}

func main() {
    sdk.RegisterFunction("create", create)
    sdk.RegisterFunction("alipay_create", alipayCreate)
    sdk.RegisterFunction("wechat_create", wechatCreate)
    sdk.RegisterFunction("alipay_verify", alipayVerify)
    sdk.RegisterFunction("wechat_verify", wechatVerify)
    sdk.RegisterFunction("alipay_refund", alipayRefund)
    sdk.RegisterFunction("wechat_refund", wechatRefund)
    _ = sdk.Start()
    sdk.Wait()
}

func signRSA256(privatePEM string, msg string) (string, error) {
    block, _ := pem.Decode([]byte(privatePEM))
    if block == nil { return "", errors.New("invalid pem") }
    key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil { pkcs8, err2 := x509.ParsePKCS8PrivateKey(block.Bytes); if err2 != nil { return "", err } ; if k, ok := pkcs8.(*rsa.PrivateKey); ok { key = k } else { return "", errors.New("no rsa key") } }
    h := sha256.Sum256([]byte(msg))
    sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, h[:])
    if err != nil { return "", err }
    return base64.StdEncoding.EncodeToString(sig), nil
}

func verifyRSA256(publicPEM string, msg string, sigBase64 string) bool {
    block, _ := pem.Decode([]byte(publicPEM))
    if block == nil { return false }
    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil { cert, err2 := x509.ParseCertificate(block.Bytes); if err2 != nil { return false } ; pub = cert.PublicKey }
    rp, ok := pub.(*rsa.PublicKey)
    if !ok { return false }
    sig, err := base64.StdEncoding.DecodeString(sigBase64)
    if err != nil { return false }
    h := sha256.Sum256([]byte(msg))
    return rsa.VerifyPKCS1v15(rp, crypto.SHA256, h[:], sig) == nil
}

func canonicalize(m map[string]string) string {
    keys := make([]string, 0, len(m))
    for k := range m { keys = append(keys, k) }
    sort.Strings(keys)
    parts := make([]string, 0, len(keys))
    for _, k := range keys { parts = append(parts, k+"="+m[k]) }
    return strings.Join(parts, "&")
}

func randHex(n int) string { b := make([]byte, n); _, _ = rand.Read(b); return hex.EncodeToString(b) }

func alipayRefund(params map[string]interface{}) (interface{}, error) {
    reason := fmt.Sprint(params["reason"]) ; amount := fmt.Sprint(params["amount"]) ; orderNo := fmt.Sprint(params["order_no"]) ; if orderNo=="" || amount=="" { return nil, fmt.Errorf("missing fields") }
    txn := randHex(8)
    return map[string]interface{}{"order_no": orderNo, "refund_id": "alr_"+txn, "status": "succeeded", "amount": amount, "reason": reason}, nil
}

func wechatRefund(params map[string]interface{}) (interface{}, error) {
    reason := fmt.Sprint(params["reason"]) ; amount := fmt.Sprint(params["amount"]) ; orderNo := fmt.Sprint(params["order_no"]) ; if orderNo=="" || amount=="" { return nil, fmt.Errorf("missing fields") }
    txn := randHex(8)
    return map[string]interface{}{"order_no": orderNo, "refund_id": "wcr_"+txn, "status": "succeeded", "amount": amount, "reason": reason}, nil
}