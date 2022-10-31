/**
 * 网站配置文件
 */

const config = {
  appName: 'MagicAdmin',
  appLogo: 'https://i.postimg.cc/Wb026mfD/logo.jpg',
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
