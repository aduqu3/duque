import axios from "axios";
import Cookies from "js-cookie";
const api = process.env.NUXT_ENV_URL_SERVER_API;


export default {
  get_token(username, password) {
    let headers_ = {
      "Accept-Language": "es"
    }
    try {
      return axios.post(api + "token/", {
        username: username,
        password: password,

      }, {headers: headers_},).then((data) => {
        Cookies.set('access', String(data.data.access), {
          expires: 365
        });
        Cookies.set('refresh', String(data.data.refresh), {
          expires: 365
        });

        // store.commit('setAccess', data.data.access)
        // store.commit('setRefresh', data.data.refresh)
        // store.commit('setAuthenticated', true)
        return data.status
      }).catch((error) => {
        if (error.response.status == 401) {
          // store.commit('setAuthenticated', false)
          return error.response.data.message
        }
        return error.response.status
      })
    } catch (e) {
      // store.commit('setAuthenticated', false)
      return e.response.status
    }

  },
  verify_token(access) {
    let headers_ = {
      "Accept-Language": "es"
    }
    try {
      return axios.post(api + "token/verify/", {
        token: access,
      }, {headers: headers_},).then((data) => {
        return data.status
      }).catch((error) => {
        if (error.response.status == 401) {
          return error.response.data.message
        }
        return error.response.status
      })
    } catch (e) {
      return e.response.status
    }

  },
  refresh_token(refresh) {
    let headers_ = {
      "Accept-Language": "es"
    }
    try {
      return axios.post(api + "token/refresh/", {
        refresh: refresh,
      }, {headers: headers_},).then((data) => {
        return data.data
      }).catch((error) => {
        if (error.response.status == 401) {
          return error.response.data.message
        }
        return error.response.status
      })
    } catch (e) {
      return e.response.status
    }

  },
  create(todo) {
    return axios.post(api, todo);
  },
};
