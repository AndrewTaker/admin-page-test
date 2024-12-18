// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  css: ["./assets/stylesheet.css"],
  runtimeConfig: ({
    public: {
      apiUrl: process.env.API_BASE_URL || "http://golang-backend:6969",
    },
  }),
});
