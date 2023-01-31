import axios from "axios";
const api = "http://127.0.0.1:8000/api/"
export default {
  async get_worst_estudiantes() {
    // const config = {
    //   headers: { Authorization: 'Bearer '+token }
    // };
    // console.log(config)
    try {
      return axios.get(api + "mejor").then((data) => {
        // console.log("XXX")
        // console.log(data.data.data)
        // console.log("XXX")
        return data.data.data
      })

    } catch (e) {
      if (e.response) {
        // Request made and server responded
        console.log('error response')
        console.log(e.response.data);
        console.log(e.response.status);
        console.log(e.response.headers);
      } else if (e.request) {
        // The request was made but no response was received
        console.log('error request')
        console.log(e.request);
      } else {
        // Something happened in setting up the request that triggered an Error
        console.log('error other')
        console.log('Error', e.message);
      }
    }

  },
  create(todo) {
    return axios.post(api, todo);
  },
}