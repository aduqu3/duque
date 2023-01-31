# Env

## File env

create file in root dir with name **.env.local** add this content

    NUXT_ENV_URL_SERVER=https://roush.loducode.com
    NUXT_ENV_PATH_API=/api/
    NUXT_ENV_URL_SERVER_API=https://roush.loducode.com/api/
    NUXT_ENV_ENV=Local

### for use envs

    console.log(process.env.URL_SERVER_API)
    console.log(process.env.PATH_API)
