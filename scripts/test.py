import requests, json, time, hmac, hashlib, base64, sys, os, tempfile, subprocess

def b64url(b):
    return base64.urlsafe_b64encode(b).rstrip(b"=")

def gen_jwt(secret="devsecret", roles=("admin",), exp_minutes=30):
    header = {"alg":"HS256","typ":"JWT"}
    exp = int(time.time() + exp_minutes*60)
    payload = {"roles":list(roles),"exp":exp}
    h_b64 = b64url(json.dumps(header,separators=(",",":")).encode("utf-8"))
    p_b64 = b64url(json.dumps(payload,separators=(",",":")).encode("utf-8"))
    unsigned = h_b64 + b"." + p_b64
    sig = hmac.new(secret.encode("utf-8"), unsigned, hashlib.sha256).digest()
    token = (unsigned + b"." + b64url(sig)).decode("utf-8")
    return token

def post_json(url, body, headers=None):
    return requests.post(url, headers={"Content-Type":"application/json",**(headers or {})}, data=json.dumps(body))

def main():
    base = sys.argv[1] if len(sys.argv)>1 else "http://localhost:8080"
    failed = 0
    print("[SKIP] docs endpoints (removed from framework)")
    print("[GET] /debug/plugins")
    r = requests.get(base+"/debug/plugins"); print(r.status_code); print(r.text)
    print("[POST] /api/text/reverse")
    r = post_json(base+"/api/text/reverse", {"text":"Hello World"}); print(r.status_code); print(r.text); failed += 0 if r.status_code==200 and r.json().get('code')==0 else 1
    print("[POST] /api/files/upload")
    fd, path = tempfile.mkstemp(suffix=".txt"); os.write(fd, b"hello"); os.close(fd)
    with open(path,"rb") as f:
        r = requests.post(base+"/api/files/upload", files={"files":(os.path.basename(path), f)})
    print(r.status_code); print(r.text); failed += 0 if r.status_code in (200,201) and r.json().get('code')==0 else 1
    os.remove(path)
    print("[POST] /api/vector/demo")
    r = post_json(base+"/api/vector/demo", {"id":"u1","vec":[0.1,0.2,0.3]}); print(r.status_code); print(r.text); failed += 0 if r.status_code==200 and r.json().get('code')==0 else 1
    qdrant = os.environ.get("QDRANT_ENDPOINT")
    if qdrant:
        print("[POST] /api/vector/qdrant/init")
        r = post_json(base+"/api/vector/qdrant/init", {"size":3,"metric":"Cosine","collection":"default"}); print(r.status_code); print(r.text)
        print("[POST] /api/vector/qdrant/upsert_batch")
        items = [{"id":"q1","vec":[0.1,0.2,0.3],"meta":{"k":"v"}},{"id":"q2","vec":[0.2,0.1,0.3],"meta":{}}]
        r = post_json(base+"/api/vector/qdrant/upsert_batch", {"collection":"default","items":items}); print(r.status_code); print(r.text)
        print("[POST] /api/vector/qdrant/search (no options)")
        r = post_json(base+"/api/vector/qdrant/search", {"vec":[0.1,0.2,0.3],"topK":2}); print(r.status_code); print(r.text);
        print("[POST] /api/vector/qdrant/search (with options.filter)")
        opts = {"filter":{"must":[{"key":"k","match":{"value":"v"}}]}}
        r = post_json(base+"/api/vector/qdrant/search", {"vec":[0.1,0.2,0.3],"topK":10,"options":opts}); print(r.status_code); print(r.text)
    print("[GET] /api/secure/ping (jwt)")
    token = gen_jwt(); r = requests.get(base+"/api/secure/ping", headers={"Authorization":"Bearer "+token}); print(r.status_code); print(r.text); failed += 0 if r.status_code==200 and r.json().get('code')==0 else 1
    print("[POST] /admin/migrate (jwt)")
    r = requests.post(base+"/admin/migrate", headers={"Authorization":"Bearer "+token}); print(r.status_code); print(r.text)
    # build plugins if missing (Windows paths)
    try:
        if os.name == 'nt':
            builds = [
                ("plugins/auth", "plugins/auth.exe"),
                ("plugins/sig", "plugins/sig.exe"),
                ("plugins/quota", "plugins/quota.exe"),
                ("plugins/metrics", "plugins/metrics.exe"),
            ]
            for src, out in builds:
                if not os.path.exists(out):
                    print("[BUILD]", out)
                    subprocess.check_call(["go","build","-o",out,"./"+src])
    except Exception as e:
        print("build plugins failed/skipped:", e)
    api_key = os.environ.get("API_KEY") or "ak_test_123"
    api_secret = os.environ.get("API_SECRET") or "sk_test_456"
    if api_key and api_secret:
        print("[POST] /api/service/compute (apikey+hmac)")
        ts = time.strftime('%Y-%m-%dT%H:%M:%SZ', time.gmtime())
        nonce = "n"+str(int(time.time()*1000))
        body = {"x":2,"y":3}
        body_hex = hashlib.sha256(json.dumps(body,separators=(',',':')).encode('utf-8')).hexdigest()
        msg = "POST\n/api/service/compute\n"+ts+"\n"+nonce+"\n"+body_hex
        sig = b64url(hmac.new(api_secret.encode('utf-8'), msg.encode('utf-8'), hashlib.sha256).digest()).decode('utf-8')
        tenant_id = os.environ.get("TENANT_ID") or "1"
        headers = {"Content-Type":"application/json","X-Api-Key":api_key,"X-Timestamp":ts,"X-Nonce":nonce,"X-Signature":sig, "X-Tenant-Id": tenant_id}
        r = requests.post(base+"/api/service/compute", headers=headers, data=json.dumps(body, separators=(',',':'))); print(r.status_code); print(r.text)
    print("[POST] /admin/reload")
    rr = requests.post(base+"/admin/reload"); print(rr.status_code); print(rr.text)
    print("[GET] /admin/hooks/inspect")
    ir = requests.get(base+"/admin/hooks/inspect?method=POST&path=/api/service/compute"); print(ir.status_code); print(ir.text)
    print("[POST] /api/pay/checkout (jwt)")
    token = gen_jwt()
    r = requests.post(base+"/api/pay/checkout", headers={"Authorization":"Bearer "+token, "Content-Type":"application/json"}, data=json.dumps({"tenant_id":1, "amount": "1.00", "currency":"CNY", "provider":"alipay"}))
    print(r.status_code); print(r.text)
    if r.status_code in (200,201):
        data = r.json().get('data') or {}
        order_no = (data.get('order_no') or r.json().get('order_no'))
        print("[POST] /api/pay/callback/alipay")
        cb = {"order_no": order_no, "status":"succeeded", "txn_id":"tx_"+str(int(time.time()*1000)), "sig":""}
        rc = post_json(base+"/api/pay/callback/alipay", cb); print(rc.status_code); print(rc.text)
    print("[POST] /api/pay/checkout (wechat+jwt)")
    r = requests.post(base+"/api/pay/checkout", headers={"Authorization":"Bearer "+token, "Content-Type":"application/json"}, data=json.dumps({"tenant_id":1, "amount": "1", "currency":"CNY", "provider":"wechat"}))
    print(r.status_code); print(r.text)
    if r.status_code in (200,201):
        data = r.json().get('data') or {}
        order_no = (data.get('order_no') or r.json().get('order_no'))
    print("[POST] /api/pay/callback/wechat (with headers)")
    headers = {"Content-Type":"application/json","Wechatpay-Timestamp": str(int(time.time())), "Wechatpay-Nonce": "n"+str(int(time.time()*1000)), "Wechatpay-Signature": ""}
    cb = {"order_no": order_no, "status":"succeeded", "txn_id":"tx_"+str(int(time.time()*1000)), "amount":"1"}
    rc = requests.post(base+"/api/pay/callback/wechat", headers=headers, data=json.dumps(cb)); print(rc.status_code); print(rc.text)
    print("[POST] /api/usage/reports (groupBy=tenant)")
    r = post_json(base+"/api/usage/reports", {"rangeStart":"2025-01-01","rangeEnd":"2025-12-31","groupBy":"tenant"}); print(r.status_code); print(r.text)
    print("[POST] /api/usage/reports (groupBy=endpoint)")
    r = post_json(base+"/api/usage/reports", {"rangeStart":"2025-01-01","rangeEnd":"2025-12-31","groupBy":"endpoint"}); print(r.status_code); print(r.text)
    print("[POST] /api/usage/reports (groupBy=provider)")
    r = post_json(base+"/api/usage/reports", {"rangeStart":"2025-01-01","rangeEnd":"2025-12-31","groupBy":"provider"}); print(r.status_code); print(r.text)
    print("[GET] /api/usage/export.csv (groupBy=tenant)")
    r = requests.get(base+"/api/usage/export.csv?groupBy=tenant&rangeStart=2025-01-01&rangeEnd=2025-12-31"); print(r.status_code); print(r.text.splitlines()[0])
    print("[GET] /api/usage/export.csv (groupBy=endpoint)")
    r = requests.get(base+"/api/usage/export.csv?groupBy=endpoint&rangeStart=2025-01-01&rangeEnd=2025-12-31"); print(r.status_code); print(r.text.splitlines()[0])
    print("[GET] /api/usage/export.csv (groupBy=provider)")
    r = requests.get(base+"/api/usage/export.csv?groupBy=provider&rangeStart=2025-01-01&rangeEnd=2025-12-31"); print(r.status_code); print(r.text.splitlines()[0])
    print("[GET] /api/admin/remote/test (jwt)")
    r = requests.get(base+"/api/admin/remote/test", headers={"Authorization":"Bearer "+token}); print(r.status_code); print(r.text)
    print("FAILED:", failed)
    sys.exit(1 if failed else 0)

if __name__ == "__main__":
    main()