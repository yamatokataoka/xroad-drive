module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    port: 8080,
    https: false,
    proxy: {
      '^/api': {
        target: process.env.VUE_APP_API_URL || 'http://localhost:8082',
        changeOrigin: true
      }
    }
  }
}