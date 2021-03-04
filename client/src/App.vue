<template>
    <RoomForm v-if="room_id.length === 0"></RoomForm>
    <VideoRoom v-else></VideoRoom>
</template>

<script>
import { provide, computed } from "vue"
import RoomForm from "./components/RoomForm.vue"
import VideoRoom from "./components/VideoRoom.vue"
export default {
    components: { RoomForm, VideoRoom },
    setup() {
        const socket = new WebSocket("ws://127.0.0.1:8199/ws")

        provide("socket", socket)

        socket.onmessage = e => {
            const { type, content } = JSON.parse(e.data)
            switch (type) {
                // 链接成功，保存UserId
                case "connect":
                    window.localStorage.setItem("USER_ID", content)
            }
        }

        let room_id = computed(() => window.localStorage.getItem("ROOM_ID"))

        return { room_id }
    },
}
</script>

<style lang="less">
.video-window {
    height: 100%;
    padding-left: 10px;
    padding-top: 6px;
    display: flex;

    & > .left {
        height: 100%;
        width: 80%;
        position: relative;

        video {
            width: 100%;
            height: 96%;
            background: #000;
            border-radius: 10px;
        }

        div {
            position: absolute;
            top: 0;
            left: 0;
            right: 0;
            bottom: 0;
            margin: auto;
            color: #fff;
        }
    }

    & > .right {
        width: 20%;
        text-align: center;
        padding: 20px;
    }
}
</style>
