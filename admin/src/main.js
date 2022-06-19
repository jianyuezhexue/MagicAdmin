import { createApp } from 'vue'

import 'element-plus/dist/index.css'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import { store } from '@/store/index'
import router from '@/router/index'

import App from './App.vue'

createApp(App)
    .use(store)
    .use(router)
    .use(ElementPlus, { locale: zhCn })
    .mount('#app')
