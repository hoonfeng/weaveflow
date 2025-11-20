package sdk

import (
    "fmt"
    "strings"
    "ifaceconf/internal/config"
)

func TsClient(interfaces []*config.InterfaceConfig) string {
    var b strings.Builder
    b.WriteString("export class ApiClient {\n")
    b.WriteString("  constructor(base, token){ this.base=base||''; this.token=token||'' }\n")
    b.WriteString("  async request(method, path, body){ const h={}; if(this.token){ h['Authorization']='Bearer '+this.token } ; let opt={ method, headers:h }; if(body){ h['Content-Type']='application/json'; opt.body=JSON.stringify(body) } ; const r=await fetch(this.base+path, opt); const t=await r.text(); try{ return {status:r.status, data:JSON.parse(t)} }catch{ return {status:r.status, data:t} } }\n")
    b.WriteString("}\n")
    for _, ic := range interfaces {
        name := sanitize(ic.Module + "_" + ic.Endpoint)
        b.WriteString(fmt.Sprintf("ApiClient.prototype.%s = async function(payload){ return this.request('%s','%s',payload) }\n", name, strings.ToUpper(ic.Method), ic.Path))
    }
    return b.String()
}

func PyClient(interfaces []*config.InterfaceConfig) string {
    var b strings.Builder
    b.WriteString("import requests, json, time, hmac, hashlib, base64\n\n")
    b.WriteString("class ApiClient:\n")
    b.WriteString("  def __init__(self, base='', token='', api_key=None, api_secret=None):\n    self.base=base\n    self.token=token\n    self.api_key=api_key\n    self.api_secret=api_secret\n")
    b.WriteString("  def request(self, method, path, body=None):\n    h={}\n    if self.token: h['Authorization']='Bearer '+self.token\n    url=self.base+path\n    if body is not None:\n      r=requests.request(method, url, headers={**h,'Content-Type':'application/json'}, data=json.dumps(body))\n    else:\n      r=requests.request(method, url, headers=h)\n    try:\n      return {'status': r.status_code, 'data': r.json()}\n    except Exception:\n      return {'status': r.status_code, 'data': r.text}\n")
    b.WriteString("  def request_signed(self, method, path, body=None):\n    if not (self.api_key and self.api_secret):\n      return self.request(method, path, body)\n    ts=time.strftime('%Y-%m-%dT%H:%M:%SZ', time.gmtime())\n    nonce='n'+str(int(time.time()*1000))\n    body_hex=hashlib.sha256(json.dumps(body or {}, separators=(',',':')).encode('utf-8')).hexdigest()\n    msg=method+'\\n'+path+'\\n'+ts+'\\n'+nonce+'\\n'+body_hex\n    sig=base64.urlsafe_b64encode(hmac.new(self.api_secret.encode('utf-8'), msg.encode('utf-8'), hashlib.sha256).digest()).decode('utf-8').rstrip('=')\n    h={'X-Api-Key':self.api_key,'X-Timestamp':ts,'X-Nonce':nonce,'X-Signature':sig}\n    url=self.base+path\n    if body is not None:\n      r=requests.request(method, url, headers={**h,'Content-Type':'application/json'}, data=json.dumps(body, separators=(',',':')) )\n    else:\n      r=requests.request(method, url, headers=h)\n    try:\n      return {'status': r.status_code, 'data': r.json()}\n    except Exception:\n      return {'status': r.status_code, 'data': r.text}\n")
    for _, ic := range interfaces {
        name := sanitize(ic.Module + "_" + ic.Endpoint)
        b.WriteString(fmt.Sprintf("  def %s(self, payload=None):\n    return self.request('%s','%s',payload)\n", name, strings.ToUpper(ic.Method), ic.Path))
        if strings.ToLower(ic.Auth) == "apikey" {
            b.WriteString(fmt.Sprintf("  def %s_signed(self, payload=None):\n    return self.request_signed('%s','%s',payload)\n", name, strings.ToUpper(ic.Method), ic.Path))
        }
    }
    return b.String()
}

func sanitize(s string) string {
    t := strings.Map(func(r rune) rune { if (r>='a'&&r<='z')||(r>='A'&&r<='Z')||(r>='0'&&r<='9')||r=='_' { return r } ; return '_' }, s)
    return t
}