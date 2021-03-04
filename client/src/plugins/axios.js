import axios from "axios"
import { message } from "ant-design-vue"

let pending = []
const CancelToken = axios.CancelToken

const removePending = (config, f) => {
    // 获取请求的url
    const flagUrl = config.url
    // 判断该请求是否在请求队列中
    if (pending.indexOf(flagUrl) !== -1) {
        // 如果在请求中，并存在f,f即axios提供的取消函数
        if (f) {
            f("取消重复请求") // 执行取消操作
        } else {
            pending.splice(pending.indexOf(flagUrl), 1) // 把这条记录从数组中移除
        }
    } else {
        // 如果不存在在请求队列中，加入队列
        if (f) {
            pending.push(flagUrl)
        }
    }
}
const instance = axios.create({
    baseURL: "http://127.0.0.1:8199",
    crossDomain: true,
})

/* request拦截器 */
instance.interceptors.request.use(
    config => {
        // neverCancel 配置项，允许多个请求
        if (!config.neverCancel) {
            // 生成cancelToken
            config.cancelToken = new CancelToken(c => {
                removePending(config, c)
            })
        }
        // 在这里可以统一修改请求头，例如 加入 用户 token 等操作
        // if (store.state.system.token) {
        //     config.headers["Authorization"] = "Bearer " + store.state.system.token
        // }

        return config
    },
    error => {
        Promise.reject(error)
    }
)

/* response 拦截器 */
instance.interceptors.response.use(
    response => {
        // 移除队列中的该请求，注意这时候没有传第二个参数f
        removePending(response.config)

        // if (response.headers.authorization && response.headers.authorization.length > 0) {
        //     store.commit("refreshToken", response.headers.authorization)
        // }

        return response.data
    },
    error => {
        let msg = "服务器繁忙，请稍后再试"
        pending = []
        if (error.message === "取消重复请求") {
            return Promise.reject(error)
        }
        if (error.response) {
            let { status, data } = error.response
            if (status === 401) {
                // store.commit("logout")
                // router.replace({ name: "login" })
                msg = data.message
            }

            if (process.env.NODE_ENV === "development") {
                msg = data.message
            }
        }

        message.error(msg)

        return Promise.reject(error)
    }
)

export function get(url, params) {
    return instance({ url: url, params: params, method: "get" })
}

export function post(url, params) {
    return instance({ url: url, data: params, method: "post" })
}

export function put(url, params) {
    return instance({ url: url, data: params, method: "put" })
}

export function destroy(url, params) {
    return instance({ url: url, data: params, method: "delete" })
}

export default instance
