<template>
    <MacWindow>
        <div class="video-window">
            <div class="left">
                <video src=""></video>
                <div>等待连接...</div>
            </div>
            <div class="right">
                <div class="rooms">
                    <a-divider orientation="left">公共聊天室</a-divider>
                    <a-list item-layout="horizontal" :data-source="rooms">
                        <template #renderItem="{ item }">
                            <a-list-item>
                                <a-list-item-meta description="Ant Design, a design language for background applications, is refined by Ant UED Team">
                                    <template #title>
                                        <a href="https://www.antdv.com/">{{ item.title }}</a>
                                    </template>
                                </a-list-item-meta>
                            </a-list-item>
                        </template>
                    </a-list>
                </div>
            </div>
        </div>
    </MacWindow>
</template>

<script>
import { computed } from "vue"
export default {
    name: "VideoRoom",
    setup() {
        let room_id = computed(() => window.localStorage.getItem("ROOM_ID"))

        return { room_id }
    },
}
</script>

<style scoped></style>
