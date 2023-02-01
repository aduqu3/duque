<template>
  <div v-if="curso.curso">
    <!-- <p>{{ curso.curso }}</p> -->
    <div v-for="(item) in students.slice(0,10)" :key="item.key" class="text-red-500 mt-4">
      <p v-if="item.ID">
        {{ item.Estudiante.nombre }} {{ item.Estudiante.apellido }} {{ item.nota }}
      </p>
    </div>
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