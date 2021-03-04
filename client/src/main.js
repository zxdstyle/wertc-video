import { createApp } from "vue"
import App from "./App.vue"
import "./assets/style/style.less"

const app = createApp(App)

import registerAntd from "./plugins/antd"
registerAntd(app)

app.mount("#app")
