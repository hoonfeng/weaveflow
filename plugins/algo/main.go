package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "fmt"
    "strings"
    "golang.org/x/crypto/pbkdf2"
    sdk "github.com/hoonfeng/goproc/sdk"
)

func toInt(v any, def int) int { switch t := v.(type) { case int: return t; case int64: return int(t); case float64: return int(t); case string: var n int; _, _ = fmt.Sscan(t, &n); if n!=0 { return n }; return def; default: return def } }
func toBytes(s string) []byte { return []byte(s) }
func sha256Hex(b []byte) string { h := sha256.Sum256(b); return hex.EncodeToString(h[:]) }
func normalizeKey(key []byte, want int) []byte {
    if want == 0 { if l := len(key); l == 16 || l == 24 || l == 32 { return key } ; want = 32 }
    if len(key) == want { return key }
    if len(key) > want { return key[:want] }
    d := make([]byte, want)
    copy(d, key)
    i := len(key)
    for i < want { d[i] = byte(i); i++ }
    return d
}
func normalizeIV(iv []byte) []byte { if len(iv) >= aes.BlockSize { return iv[:aes.BlockSize] } ; out := make([]byte, aes.BlockSize); copy(out, iv); for i := len(iv); i < aes.BlockSize; i++ { out[i] = byte(i) } ; return out }
func pkcs7Pad(b []byte) []byte { pad := aes.BlockSize - (len(b) % aes.BlockSize); if pad == 0 { pad = aes.BlockSize } ; out := make([]byte, len(b)+pad); copy(out, b); for i:=len(b); i<len(out); i++ { out[i] = byte(pad) } ; return out }
func pkcs7Unpad(b []byte) ([]byte, error) { if len(b) == 0 { return nil, errors.New("empty") } ; p := int(b[len(b)-1]); if p <= 0 || p > aes.BlockSize || p > len(b) { return nil, errors.New("badpad") } ; for i := len(b)-p; i < len(b); i++ { if int(b[i]) != p { return nil, errors.New("badpad") } } ; return b[:len(b)-p], nil }

func fnHash(params map[string]interface{}) (interface{}, error) {
    alg := strings.ToLower(fmt.Sprint(params["algorithm"]))
    text := fmt.Sprint(params["text"])
    if s, ok := params["salt"]; ok { text = text + fmt.Sprint(s) }
    switch alg {
    case "sha256": s := sha256.Sum256([]byte(text)); return hex.EncodeToString(s[:]), nil
    case "sha512": s := sha512.Sum512([]byte(text)); return hex.EncodeToString(s[:]), nil
    case "md5": s := md5.Sum([]byte(text)); return hex.EncodeToString(s[:]), nil
    default:
        return nil, fmt.Errorf("unsupported algorithm: %s", alg)
    }
}

func fnHMAC(params map[string]interface{}) (interface{}, error) {
    alg := strings.ToLower(fmt.Sprint(params["algorithm"]))
    text := fmt.Sprint(params["text"])
    key := []byte(fmt.Sprint(params["key"]))
    want := toInt(params["length"], 0)
    if want > 0 { key = normalizeKey(key, want) }
    var mac hashFn
    switch alg { case "sha256": mac = func() cipherHash { return cipherHash{hmac.New(sha256.New, key)} } ; case "sha512": mac = func() cipherHash { return cipherHash{hmac.New(sha512.New, key)} } ; case "md5": mac = func() cipherHash { return cipherHash{hmac.New(md5.New, key)} } ; default: return nil, fmt.Errorf("unsupported algorithm: %s", alg) }
    h := mac().h
    _, _ = h.Write([]byte(text))
    return hex.EncodeToString(h.Sum(nil)), nil
}

type cipherHash struct{ h hash }
type hash interface{ Write([]byte) (int, error); Sum([]byte) []byte }
type hashFn func() cipherHash

func fnRandom(params map[string]interface{}) (interface{}, error) {
    n := toInt(params["bytes"], 16)
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil { return nil, err }
    return hex.EncodeToString(b), nil
}

func fnPBKDF2(params map[string]interface{}) (interface{}, error) {
    pass := fmt.Sprint(params["password"])
    salt := []byte(fmt.Sprint(params["salt"]))
    iter := toInt(params["iterations"], 10000)
    klen := toInt(params["keyLen"], 32)
    alg := strings.ToLower(fmt.Sprint(params["algorithm"]))
    var dk []byte
    switch alg { case "sha512": dk = pbkdf2.Key([]byte(pass), salt, iter, klen, sha512.New) ; default: dk = pbkdf2.Key([]byte(pass), salt, iter, klen, sha256.New) }
    return hex.EncodeToString(dk), nil
}

func fnAESEncrypt(params map[string]interface{}) (interface{}, error) {
    key := normalizeKey([]byte(fmt.Sprint(params["key"])), toInt(params["keyLen"], 0))
    iv := normalizeIV([]byte(fmt.Sprint(params["iv"])) )
    mode := strings.ToLower(fmt.Sprint(params["mode"]))
    enc := strings.ToLower(fmt.Sprint(params["encoding"]))
    pt := []byte(fmt.Sprint(params["plaintext"]))
    b, err := aes.NewCipher(key); if err != nil { return nil, err }
    var ct []byte
    switch mode { case "ctr": ct = make([]byte, len(pt)); cipher.NewCTR(b, iv).XORKeyStream(ct, pt) ; default: p := pkcs7Pad(pt); ct = make([]byte, len(p)); cipher.NewCBCEncrypter(b, iv).CryptBlocks(ct, p) }
    switch enc { case "base64": return base64.StdEncoding.EncodeToString(ct), nil ; default: return hex.EncodeToString(ct), nil }
}

func fnAESDecrypt(params map[string]interface{}) (interface{}, error) {
    key := normalizeKey([]byte(fmt.Sprint(params["key"])), toInt(params["keyLen"], 0))
    iv := normalizeIV([]byte(fmt.Sprint(params["iv"])) )
    mode := strings.ToLower(fmt.Sprint(params["mode"]))
    enc := strings.ToLower(fmt.Sprint(params["encoding"]))
    ctStr := fmt.Sprint(params["ciphertext"]) ; var ct []byte ; var err error
    switch enc { case "base64": ct, err = base64.StdEncoding.DecodeString(ctStr) ; default: ct, err = hex.DecodeString(ctStr) } ; if err != nil { return nil, err }
    b, err := aes.NewCipher(key); if err != nil { return nil, err }
    var pt []byte
    switch mode { case "ctr": pt = make([]byte, len(ct)); cipher.NewCTR(b, iv).XORKeyStream(pt, ct) ; default: pt = make([]byte, len(ct)); cipher.NewCBCDecrypter(b, iv).CryptBlocks(pt, ct); pt, err = pkcs7Unpad(pt); if err != nil { return nil, err } }
    return string(pt), nil
}

func main() {
    sdk.RegisterFunction("hash", fnHash)
    sdk.RegisterFunction("hmac", fnHMAC)
    sdk.RegisterFunction("random", fnRandom)
    sdk.RegisterFunction("pbkdf2", fnPBKDF2)
    sdk.RegisterFunction("aes_encrypt", fnAESEncrypt)
    sdk.RegisterFunction("aes_decrypt", fnAESDecrypt)
    _ = sdk.Start()
    sdk.Wait()
}