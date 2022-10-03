/**
 * 网站配置文件
 */

const config = {
  appName: '渠道数据中台',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用MagicAdmin`
      )
    )
  }
}

export default config
