const http = require('http')
const crypto = require('crypto')
const fs = require('fs')

async function fetchJson(method, url, body, headers={}){
  const u = new URL(url)
  const opts = { method, headers: { ...headers } }
  return new Promise((resolve,reject)=>{
    const req = http.request({ hostname: u.hostname, port: u.port||80, path: u.pathname+u.search, method: opts.method, headers: opts.headers }, res =>{
      let data=''
      res.on('data', c=> data+=c)
      res.on('end', ()=> resolve({ status: res.statusCode, text: data }))
    })
    req.on('error', reject)
    if(body){ const b = JSON.stringify(body); opts.headers['Content-Type']='application/json'; opts.headers['Content-Length']=Buffer.byteLength(b); req.write(b) }
    req.end()
  })
}

function genJWT(secret, roles=['admin'], expMinutes=30){
  const h = Buffer.from(JSON.stringify({alg:'HS256',typ:'JWT'})).toString('base64url')
  const exp = Math.floor(Date.now()/1000 + expMinutes*60)
  const p = Buffer.from(JSON.stringify({roles,exp})).toString('base64url')
  const unsigned = `${h}.${p}`
  const sig = crypto.createHmac('sha256', secret).update(unsigned).digest('base64url')
  return `${unsigned}.${sig}`
}

async function main(){
  const base = process.argv[2] || 'http://localhost:8080'
  let failed = 0
  let r
  r = await fetchJson('POST', base+'/api/text/reverse', {text:'Hello World'})
  console.log('[POST] /api/text/reverse', r.status, r.text)
  const jwt = genJWT(process.env.JWT_SECRET || 'devsecret')
  r = await fetchJson('GET', base+'/api/secure/ping', null, { Authorization: 'Bearer '+jwt })
  console.log('[GET] /api/secure/ping', r.status, r.text)
  r = await fetchJson('POST', base+'/admin/migrate', null, { Authorization: 'Bearer '+jwt })
  console.log('[POST] /admin/migrate', r.status, r.text)
  const apiKey = process.env.API_KEY || 'ak_test_123'
  const apiSecret = process.env.API_SECRET || 'sk_test_456'
  const ts = new Date().toISOString().replace(/\.\d+Z$/,'Z')
  const nonce = 'n'+Date.now()
  const body = { x:2, y:3 }
  const bodyHex = crypto.createHash('sha256').update(JSON.stringify(body)).digest('hex')
  const msg = `POST\n/api/service/compute\n${ts}\n${nonce}\n${bodyHex}`
  const sig = crypto.createHmac('sha256', apiSecret).update(msg).digest('base64url')
  r = await fetchJson('POST', base+'/api/service/compute', body, { 'X-Api-Key': apiKey, 'X-Timestamp': ts, 'X-Nonce': nonce, 'X-Signature': sig, 'Content-Type':'application/json', 'X-Tenant-Id': process.env.TENANT_ID || '1' })
  console.log('[POST] /api/service/compute', r.status, r.text)
  r = await fetchJson('GET', base+'/admin/hooks/inspect?method=POST&path=/api/service/compute')
  console.log('[GET] /admin/hooks/inspect', r.status, r.text)
  process.exit(failed?1:0)
}

main().catch(e=>{ console.error(e); process.exit(1) })