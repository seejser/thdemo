import { createApp } from 'vue'
import NutUI from '@nutui/nutui-taro'
import '@nutui/nutui-taro/dist/style.css'

import './app.scss'
//import * as NutUI from '@nutui/nutui-taro';
console.log(NutUI); // 打印出模块的所有导出
const App = createApp({})
App.use(NutUI)

export default App
