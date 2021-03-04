import { Button, Form, Input, Divider, List, Radio } from "ant-design-vue"

const components = [Button, Form, Input, Divider, List, Radio]

export default function(app) {
    components.forEach(component => {
        app.use(component)
    })
}
