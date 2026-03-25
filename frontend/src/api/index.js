import request from './request'

export const userAPI = {
  register(data) {
    return request.post('/user/register', data)
  },
  
  login(data) {
    return request.post('/user/login', data)
  },
  
  getProfile() {
    return request.get('/user/profile')
  },
  
  updateProfile(data) {
    return request.put('/user/profile', data)
  },
  
  sendCode(phone) {
    return request.post('/user/send-code', null, { params: { phone } })
  }
}

export const collegeAPI = {
  getList(params) {
    return request.get('/colleges', { params })
  },
  
  getDetail(id) {
    return request.get(`/colleges/${id}`)
  },
  
  getMajors(id, params) {
    return request.get(`/colleges/${id}/majors`, { params })
  }
}

export const majorAPI = {
  getList(params) {
    return request.get('/majors', { params })
  },
  
  getDetail(id) {
    return request.get(`/majors/${id}`)
  }
}

export const scoreAPI = {
  query(params) {
    return request.get('/scores', { params })
  },
  
  getProbability(params) {
    return request.get('/scores/probability', { params })
  }
}

export const rankingAPI = {
  getCollegeRankings(params) {
    return request.get('/rankings/colleges', { params })
  },
  
  getMajorRankings(params) {
    return request.get('/rankings/majors', { params })
  }
}

export const qaAPI = {
  chat(data) {
    return request.post('/qa/chat', data)
  },
  
  getHistory(params) {
    return request.get('/qa/history', { params })
  },
  
  getRecommend(data) {
    return request.post('/qa/recommend', data)
  },
  
  getColleges(params) {
    return request.get('/qa/colleges', { params })
  },
  
  getMajors(params) {
    return request.get('/qa/majors', { params })
  }
}
