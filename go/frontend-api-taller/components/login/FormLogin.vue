<template>
  <form class="flex flex-col items-center w-full" action @submit.prevent="login">
    <input
      v-model="user.username"
      type="text"
      class="border pl-1 mt-2 rounded w-4/5 xl:w-1/2" placeholder="Username"
      :class="border"

    />
    <input
      v-model="user.password"
      type="password"
      class="border pl-1 mt-2 rounded w-4/5 xl:w-1/2"
      :class="border" placeholder="Password"
    />
    <p v-if="errors.length">
    <ul>
      <li
        v-for="(error) in errors " :key="error.key"
        class="text-red-500 mt-4">{{ error }}
      </li>
    </ul>
    </p>
    <button
      type="submit"
      class="bg-blue-500 text-white pt-2 pb-2 pl-4 pr-4 rounded-full mt-4 xl:w-1/5 w-1/4"
    >
      Log in
    </button>
  </form>

</template>

<script>

export default {
  name: "FormLogin",

  data: () => ({
    errors: [],
    user: {
      username:"root",
      password:"asdfgfds"
    },

  }), computed: {
    border() {
      return this.errors.length > 0 ? 'border-red-500' : 'border-gay-500'
    }

  }
  , async created() {
    // if (this.$store.getters.getAuthenticated) {
    //   this.$router.push('/')
    // }
  }, methods: {
    async login() {
      try {
        await this.$auth.loginWith('local', {
          data: this.user
        })

      } catch (err) {
        console.log(err)
      }
    }
  }
}
</script>

<style scoped></style>
