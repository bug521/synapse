import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router'
import {naiveUi} from "./utils/naive-ui.ts";

const app = createApp(App)

// 使用 Pinia 状态管理
app.use(createPinia())

// 使用路由
app.use(router)
app.use(naiveUi)

app.mount('#app')
