
<template>
  <div>
    <h1>Estudiante con mayor edad para el genero {{ data }}</h1>
    <p>{{ estudiante_.nombre }} {{ estudiante_.apellido }}</p>
    <p>{{ estudiante_.about }}</p>
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