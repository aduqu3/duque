<template>
  <div v-if="curso.curso">
    <table class="styled-table">
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
.styled-table {
  border-collapse: collapse;
  margin: 25px 0;
  font-size: 0.9em;
  font-family: sans-serif;
  min-width: 400px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);
}

.styled-table thead tr {
  background-color: #009879;
  color: #ffffff;
  text-align: left;
}

.styled-table thead tr {
  background-color: #009879;
  color: #ffffff;
  text-align: left;
}

.styled-table tbody tr {
  border-bottom: 1px solid #dddddd;
}

.styled-table tbody tr:nth-of-type(even) {
  background-color: #f3f3f3;
}
</style>