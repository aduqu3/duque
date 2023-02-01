<template>
  <div class="flex flex-col">
    <h1 class="border-2 w-60 rounded-md p-2">
      Mejor Estudiante
    </h1>
    <div v-if="moduls_.estudiante" class="mt-8 flex flex-col">
      <h1 class="mt-3 w-60 rounded-md bg-yellow-500 p-2 text-white border-yellow-500">
        {{ moduls_.estudiante.nombre }} {{ moduls_.estudiante.apellido }}
      </h1>
      <table class="table-fixed">
        <thead class="text-center">
          <tr>
            <th>Curso</th>
            <th>Nota</th>
          </tr>
        </thead>
        <tbody v-for="item in moduls_.cursos" :key="item.id">
          <tr>
            <td>{{ item.Curso.curso }}</td>
            <td>{{ item.nota }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import api from '/api/api';
export default {
  name: "BetsStudentInfo",
  data() {
    return {
      moduls_: {},
    }
  }
  , created() {
    this.person();
  }
  , methods: {
    async person() {
      try {
        var res = await api.get_best_estudiantes()
        //   console.log(res)
        this.moduls_ = (res)
      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>

<style scoped>

</style>