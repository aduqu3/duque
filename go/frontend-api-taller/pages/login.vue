<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import FormLogin from "~/components/login/FormLogin";

</script>
<template>
  <div class="flex">
    <div class="md:w-1/2 hidden md:flex items-center">
      <img class="xl:w-full xl:p-28  w-full p-24 " src="~/assets/img/img.png"/>
    </div>
    <div class="md:w-1/2 w-full grid justify-items-center items-center xl:mt-0 mt-12">
      <div class="w-full  grid justify-items-center items-center">
        <img class="w-1/6 mb-10" src="~/assets/img/img.png"/>
        <FormLogin></FormLogin>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "LoginUser",
  created() {
    if(this.$store.getters.getAuthenticated){
      this.$router.push({name: 'index'});
    }
  }


}
</script>

<style scoped></style>
