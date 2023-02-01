
<template>
  <div>
    <h1>Estudiante con mayor edad para el genero {{ data }}</h1>
    <p>{{ estudiante_.ID }}</p>
    <p>{{ estudiante_.nombre }} {{ estudiante_.apellido }}</p>
  </div>
</template>

<script>
import api from '/api/api';
export default {
  name: "OldGenderStudent",
  components: {
  },
  props: {
    data: {},
  },
  data: () => ({
    errors: [],
    estudiante_: {}
  }),
  created() {
    this.courses();
  }
  , methods: {
    async courses() {
      try {
        console.log(this.data)
        var res = await api.get_more_old_gender_student(this.data)
        console.log(res)
        this.estudiante_ = (res)
        //   return {moduls_}
      } catch (err) {
        console.log(err)
      }
    }
  }
};
</script>

<style scoped>

</style>