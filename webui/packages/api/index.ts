import axios from 'axios'

const api = axios.create({ baseURL: '' })

api.interceptors.request.use(cfg => {
  const t = localStorage.getItem('jwt')
  if (t) cfg.headers.Authorization = 'Bearer ' + t
  return cfg
})

export function setTenant(id?: string|number){
  if(!id) return
  api.defaults.headers.common['X-Tenant-Id'] = String(id)
}

export default api