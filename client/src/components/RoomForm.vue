<template>
    <MacWindow :width="370" :height="400">
        <div class="room-form">
            <h2 style="text-align: center">登录</h2>
            <a-form :model="CreateRoomForm" @finish="onCreateRoom" :wrapper-col="{ span: 20 }" :label-col="{ span: 3 }" hideRequiredMark>
                <a-form-item label="方式">
                    <a-radio-group v-model:value="CreateRoomForm.layout">
                        <a-radio value="create">创建新聊天室</a-radio>
                        <a-radio value="join">加入聊天室</a-radio>
                    </a-radio-group>
                </a-form-item>
                <a-form-item label="名称" :rules="{ required: true, message: '请输入聊天室名称' }" name="name">
                    <a-input v-model:value="CreateRoomForm.name" placeholder="聊天室名称" />
                </a-form-item>
                <a-form-item label="密码">
                    <a-input type="password" v-model:value="CreateRoomForm.password" placeholder="不设密码则公开" />
                </a-form-item>
                <a-form-item :wrapper-col="{ span: 24 }">
                    <a-button htmlType="submit" type="primary" :loading="CreateRoomForm.loading" block>{{ CreateRoomForm.layout === "create" ? "创建" : "加入" }}</a-button>
                </a-form-item>
            </a-form>
        </div>
    </MacWindow>
</template>

<script>
import { reactive } from "vue"
import MacWindow from "./MacWindow.vue"
import { post } from "../plugins/axios"

export default {
    name: "RoomForm",
    components: { MacWindow },
    setup() {
        const CreateRoomForm = reactive({
            name: "",
            password: "",
            layout: "create",
            loading: false,
        })

        const onCreateRoom = () => {
            CreateRoomForm.loading = true
            let user_id = window.localStorage.getItem("USER_ID")
            post("/room", { name: CreateRoomForm.name, password: CreateRoomForm.password, user_id: user_id })
                .then(res => {
                    CreateRoomForm.loading = false
                    window.localStorage.setItem("ROOM_ID", res.message)
                })
                .catch(() => {
                    CreateRoomForm.loading = false
                })
        }

        return { CreateRoomForm, onCreateRoom }
    },
}
</script>

<style lang="less" scoped>
.room-form {
    padding: 20px;
}
</style>
