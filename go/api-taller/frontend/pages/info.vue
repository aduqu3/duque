<template>
    <div class="grid items-center justify-items-center">
      <div class="flex flex-col md:flex-row md:space-x-20 mt-20 p-4 md:p-0">
        <BetsStudent></BetsStudent>
        <CourseAvg></CourseAvg>
      </div>
  
    </div>
  </template>
  
  <script>
  import BetsStudent from "~/components/BetsStudent";
  import CourseAvg from "~/components/CourseAvg";
  export default {
    name: "Inf",
    components: {BetsStudent, CourseAvg}
  }
  </script>
  
  <style scoped>
  </style>