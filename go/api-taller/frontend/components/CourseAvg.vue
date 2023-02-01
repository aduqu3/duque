<template>
    <div class="flex flex-col">
      <h1 class="border-2 w-60 rounded-md p-2">
        Curso
      </h1>
      <div>
        <form class="flex flex-col items-center w-full" action @submit.prevent="course_avg">
          <input
            v-model="id"
            type="text"
            class="border pl-1 mt-2 rounded w-4/5 xl:w-1/2" placeholder="course id"

          />
          <!-- <input
            v-model="user.password"
            type="password"
            class="border pl-1 mt-2 rounded w-4/5 xl:w-1/2"
            :class="border" placeholder="Password"
          /> -->
          <!-- <p v-if="errors.length">
          <ul>
            <li
              v-for="(error) in errors " :key="error.key"
              class="text-red-500 mt-4">{{ error }}
            </li>
          </ul>
          </p> -->
          <button
            type="submit"
            class="bg-blue-500 text-white pt-2 pb-2 pl-4 pr-4 rounded-full mt-4 xl:w-1/5 w-1/4"
          >
            Submit
          </button>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  import estudiantes from '/api/estudiante';
  export default {
    name: "CourseInfo",
    data: () => ({
      errors: [],
      id: ""
    })
    , created() {
    //   this.course_avg();
    }
    , methods: {
      async course_avg() {
        try {
          // const api = "http://127.0.0.1:8000/api/"
          // await axios.get(api + "avg/"+this.id).then((data) => {
          //   console.log(data)
          //   return data
          // })

          var res = await estudiantes.get_course_avg(this.id)
          console.log(res)
          // this.moduls_ = (res)
        //   return {moduls_}
        } catch (err) {
          console.log(err)
        }
      }
    }
  }
  </script>
  
  <style scoped>
  </style>