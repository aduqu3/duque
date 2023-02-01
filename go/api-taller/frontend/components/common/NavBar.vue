
<template>
  <header>
    <nav>
      <ul>
        <li>
          <nuxt-link exact to="/">Home</nuxt-link>
        </li>
        <li v-for="(item) in courses_ " :key="item.key" class="text-red-500 mt-4">

          <!-- <nuxt-link v-if="item.curso.id != ''" to="/course"> -->
          <!-- <template v-if="item.curso"> -->
          <course-detail :curso="item"></course-detail>
          <!-- </template> -->

          <!-- </nuxt-link> -->

          <!-- <nuxt-link v-if="item.curso != ''" to="/course">{{ item.curso }}</nuxt-link> -->

        </li>
        <li>
          <nuxt-link to="/women">Women</nuxt-link>
        </li>
        <li>
          <nuxt-link to="/men">Men</nuxt-link>
        </li>
      </ul>
    </nav>
  </header>
</template>

<script>
import estudiantes from '/api/estudiante';
import CourseDetail from '../CourseDetail.vue';
// import CoursePage from '../../pages/course.vue'
export default {
  name: "NavBar",
  components: {
    CourseDetail
  },
  props: {
    name: String,
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
        var res = await estudiantes.get_courses()
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
.navarea {
  overflow: hidden;
}

.capsule {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

nav {
  width: 100vw;
  height: 60px;
  background: white;
}

ul {
  padding-left: 0;
  display: flex;
  list-style: none outside none;
  justify-content: center;
  align-items: center;
}

li {
  padding: 0 50px;
}

a,
a:visited,
a:active {
  text-decoration: none;
  color: black;
}

.cartitem {
  position: relative;
  float: right;
}

.cartcount {
  font-family: 'Barlow', sans-serif;
  position: absolute;
  background: #ff2211;
  color: white;
  text-align: center;
  padding-top: 4px;
  width: 20px;
  height: 20px;
  font-size: 10px;
  margin: -5px 0 0 20px;
  border-radius: 1000px;
  font-weight: 700;
}
</style>