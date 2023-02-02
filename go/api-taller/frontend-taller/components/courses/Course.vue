
<template>
  <div style="display: grid; justify-content: center;">
    <div v-for="(item) in courses_" :key="item.key">
      <course-detail :curso="item"></course-detail>
      <course-low-score :curso="item"></course-low-score>
    </div>
  </div>
</template>

<script>
import api from '/api/api';
import CourseDetail from './CourseDetail.vue';
import CourseLowScore from './CourseLowScore.vue';
export default {
  name: "CourseInf",
  components: {
    CourseDetail, CourseLowScore
  },
  data: () => ({
    errors: [],
    courses_: []
  }),
  created() {
    this.courses();
  }
  , methods: {
    async courses() {
      try {
        var res = await api.get_courses()
        console.log(res)
        this.courses_ = (res)
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