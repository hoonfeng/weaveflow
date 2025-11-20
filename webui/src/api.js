import axios from 'axios'

const api = axios.create({
  baseURL: ''
})

api.interceptors.request.use(cfg => {
  const t = localStorage.getItem('jwt')
  if (t) cfg.headers.Authorization = 'Bearer ' + t
  return cfg
})

export default api