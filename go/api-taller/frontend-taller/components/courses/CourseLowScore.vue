<template>
  <div v-if="curso.curso">

    <table class="table">
      <thead>
        <tr>
          <th scope="col">Estudiante</th>
          <th scope="col">Nota</th>
        </tr>
      </thead>
      <tbody v-for="(item) in students.slice(0, 10)" :key="item.key">
        <tr v-if="item.ID">
          <th scope="row">{{ item.Estudiante.nombre }} {{ item.Estudiante.apellido }}</th>
          <th scope="row">{{ item.nota }}</th>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import api from '/api/api';
export default {
  name: "CourseLowScore",
  props: {
    curso: {},
  },
  data: () => ({
    errors: [],
    students: []
  })
  , created() {
    this.get_course_students();
  }
  , methods: {
    async get_course_students() {
      try {
        var res = await api.get_student_lowest_score(this.curso.ID)
        console.log(res)
        this.students = (res)
      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>

<style scoped>

</style>