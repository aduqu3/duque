
<template>
  <div>
    <div>
      <NavBar></NavBar>  
    </div> 
    <div style="display: grid; justify-content: center;">
      <old-gender-student data="male"></old-gender-student>
    </div>
  </div>
</template>

<script>
import NavBar from '../../components/common/NavBar.vue';
import OldGenderStudent from '../../components/Student/OldGenderStudent.vue';
export default {
  name: "AgeMale",
  components: {
    NavBar, OldGenderStudent
  }
}
</script>

<style scoped>
</style>