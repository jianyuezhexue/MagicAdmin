/**
 * 网站配置文件
 */

const config = {
  appName: 'MagicAdmin',
  appLogo: 'http://chuantu.xyz/t6/742/1666241837x2890373684.jpg',
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
