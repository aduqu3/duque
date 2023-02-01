
<template>
  <header>
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
  </header>
</template>

<script>
export default {
  name: "NavBar",
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