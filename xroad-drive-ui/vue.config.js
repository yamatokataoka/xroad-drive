module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    port: 8080,
    https: false,
    proxy: {
      '^/api': {
        target: process.env.VUE_APP_API_URL || 'http://localhost:8082'
      },
      '^/xroad-metadata-proxy': {
        target: process.env.VUE_APP_METADATA_PROXY_URL || 'http://localhost:8083',
        pathRewrite: { '^/xroad-metadata-proxy': '' }
      },
      '^/security-server': {
        target: process.env.VUE_APP_SECURITY_SERVER_URL || 'http://localhost:8085',
        pathRewrite: { '^/security-server': '' }
      }
    }
  }
}