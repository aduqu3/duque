<template>
  <div class="">
    <img class="" src="~/assets/img/img.png"/>
    <div class="flex flex-col">
      <div class="grid justify-items-center items-center mt-6">
        <div class="relative -mb-10">
          <img class=" w-80 h-80" src="~/assets/img/img.png"/>
          <a style="top:-102px; left:10px; position: relative">
            <img class="w-20" src="~/assets/img/cheque.png"/>
          </a>
          <a style="top:-182px; left:210px; position: relative">
            <img class="w-20" src="~/assets/img/cancelar.png">
          </a>
        </div>
      </div>
      <div class="-mt-28">
        <div class="mt-2">Name:</div>
        <div class="mt-2">Description:</div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "InfUser"
}
</script>

<style scoped>

</style>
