<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import ListUser from "~/components/home/ListUser.vue";
import InfUser from "~/components/home/InfUser.vue";
</script>

<template>
  <div class="flex flex-col md:flex-row">
    <div class="md:w-1/2 p-10 lg:p-20">
      <InfUser></InfUser>
    </div>
    <div class="md:w-1/2 p-10 lg:p-20">
      <ListUser></ListUser>
    </div>
  </div>

</template>

<script>
export default {
  name: "HomeUser",
  name1: "InfUser",
}
</script>

<style scoped>

</style>
