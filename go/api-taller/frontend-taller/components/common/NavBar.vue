
<template>
  <div style="display: grid; justify-content: center;">
    <nav>
      <ul>
        <li>
          <nuxt-link exact to="/">Home</nuxt-link>
        </li>
        <li>
          <nuxt-link exact to="/courses">Cursos</nuxt-link>
        </li>
        <li>
          <nuxt-link exact to="/student/age_female">Mujer con mayor edad</nuxt-link>
        </li>
        <li>
          <nuxt-link exact to="/student/age_male">Hombre con Mayor edad</nuxt-link>
        </li>
        <li>
          <nuxt-link exact to="/student/best">Mejor Estudiante</nuxt-link>
        </li>
        <li>
          <nuxt-link exact to="/info">Pruebas</nuxt-link>
        </li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: "NavBar",
};
</script>

<style scoped>
li {
  display: inline;
}

li {
  float: left;
}

a {
  display: block;
  padding: 8px;
  background-color: #04AA6D;
  
}
ul {
  background-color: #04AA6D;
;
}
ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  background-color: #333;
}
ul {
  position: -webkit-sticky; /* Safari */
  position: sticky;
  top: 0;
}
ul {
  border: 1px solid #e7e7e7;
  background-color: #f3f3f3;
}

li a {
  color: #666;
}

li {
  float: left;
}

li a {
  display: block;
  color: white;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
  background-color: #111;
}

.active {
  background-color: #04AA6D;
}
li {
  border-right: 1px solid #bbb;
}

li:last-child {
  border-right: none;
}

</style>