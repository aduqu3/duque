export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'base',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      {charset: 'utf-8'},
      {name: 'viewport', content: 'width=device-width, initial-scale=1'},
      {hid: 'description', name: 'description', content: ''},
      {name: 'format-detection', content: 'telephone=no'}
    ],
    link: [
      {rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}
    ]
  },

  // mode: 'spa', // spa/universal

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [],


  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
    '@nuxtjs/dotenv'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    //'~/modules/auth',
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},

  auth: {
    // localStorage: true,
    strategies: {
      local: {
        scheme: "refresh",
        token: {
          property: "access",
          // maxAge: 290
        },
        refreshToken: {
          property: "refresh",
          data: "refresh",
          // maxAge: false
        },
        user: {
          property: false,  // `user` property is now `user.property`
          autoFetch: true
        },

        endpoints: {
          login: {
            url: '/api/token/',
            method: 'post',
            propertyName: 'access',
            altProperty: 'refresh'
          }, refresh: {
            url: '/api/token/refresh/',
            method: 'post',
            propertyName: 'access',
          },
          logout: {},
          user: {url: '/api/persons/get-data-person/', method: 'get'},

        },
        tokenRequired: true,
        tokenType: 'Bearer'
      }
    },
    redirect: {
      login: "/login",
      logout: "/login",
      callback: false,
      home: "/"
    },

  },
  router: {
    middleware: ['auth']
  },
  axios: {
    baseURL: 'https://roush.loducode.com'
  },
  tailwindcss: {
    // Options
  }

}
