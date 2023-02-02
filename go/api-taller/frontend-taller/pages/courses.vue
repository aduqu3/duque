
<template>
  <div>
    <div>
      <NavBar></NavBar>  
    </div> 
    <div>
      <Course></Course>
    </div> 
  </div>
</template>

<script>
import NavBar from '../components/common/NavBar.vue';
import Course from '../components/courses/Course.vue';
export default {
  name: "CoursePage",
  components: {
    Course, NavBar
  }
}
</script>

<style scoped>
</style>