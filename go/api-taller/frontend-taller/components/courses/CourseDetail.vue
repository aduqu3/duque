<template>
  <div>
    <h1 class="temporal">
      <strong>{{ curso.curso }}</strong>
    </h1>
    <p>Promedio general de la materia {{ nota }}</p>
  </div>
</template>

<script>
import api from '/api/api';
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
        var res = await api.get_course_avg(this.curso.ID)
        console.log(res)
        this.nota = (res)
      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>

<style scoped>
h1{
  color:green
}
</style>