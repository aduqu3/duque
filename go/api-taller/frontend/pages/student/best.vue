<template>
  <div>
    <div>
      <NavBar></NavBar>  
    </div> 
    <div class="grid items-center justify-items-center">
    <div class="flex flex-col md:flex-row md:space-x-20 mt-20 p-4 md:p-0">
      <BetsStudent></BetsStudent>
    </div>
  </div>
  </div>
</template>

<script>
import NavBar from '../../components/common/NavBar.vue';
import BetsStudent from '../../components/Student/BetsStudent.vue';
export default {
  name: "BestInf",
  components: {BetsStudent, NavBar}
}
</script>

<style scoped>
</style>