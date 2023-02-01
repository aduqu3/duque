<template>
  <div v-if="nota">
    <!-- <h1>Hola detail {{ curso.ID }}</h1> -->
    <p>{{ curso.curso }}</p>
    <p>{{ nota }}</p>
  </div>
</template>

<script>
//   import axios from "axios";
import estudiantes from '/api/estudiante';
export default {
  name: "CourseDetail",
  props: {
    curso: {},
  },
  data: () => ({
    errors: [],
    nota: String
  })
  , created() {
    this.course_avg();
  }
  , methods: {
    async course_avg() {
      try {
        // const api = "http://127.0.0.1:8000/api/"
        // await axios.get(api + "avg/"+this.id).then((data) => {
        //   console.log(data)
        //   return data
        // })

        var res = await estudiantes.get_course_avg(this.curso.ID)
        console.log(res)
        this.nota = (res)
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