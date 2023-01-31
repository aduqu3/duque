<template>
    <div class="grid items-center justify-items-center">
      <div class="flex flex-col md:flex-row md:space-x-20 mt-20 p-4 md:p-0">
        <Worse></Worse>
      </div>
  
    </div>
  </template>
  
  <script>
  import Worse from "~/components/home/WorstStudents";
  export default {
    name: "StudentsInf",
    components: {Worse}
  }
  </script>
  
  <style scoped>
  </style>