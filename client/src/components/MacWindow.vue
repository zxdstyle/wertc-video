<template>
    <div class="mac-window" id="window" :style="windowStyle">
        <div class="operate">
            <div v-if="closeable" @click="$router.push({ name: 'home' })" class="item close"></div>
            <div v-if="minable" class="item mini"></div>
            <div v-if="maxable" class="item max"></div>
        </div>

        <slot></slot>
    </div>
</template>

<script>
import { computed } from "vue"

export default {
    name: "MacWindow",
    props: {
        // 窗口宽度
        width: {
            type: Number,
            default: 100,
        },
        // 窗口高度
        height: {
            type: Number,
            default: 100,
        },
        // 窗口是否可以最小化
        minable: {
            type: Boolean,
            default: true,
        },
        // 窗口是否可以最大化
        maxable: {
            type: Boolean,
            default: true,
        },
        // 窗口是否可以关闭
        closeable: {
            type: Boolean,
            default: true,
        },
        // 窗口弧度
        radius: {
            type: [Number, String],
            default: 10,
        },
    },
    setup(props) {
        let windowStyle = computed(() => {
            return {
                width: props.width + "px",
                height: props.height + "px",
                maxHeight: "80%",
                maxWidth: "80%",
                borderRadius: props.radius + "px",
            }
        })

        return { windowStyle }
    },
}
</script>

<style lang="less" scoped>
.mac-window {
    background: #fff;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 80px;
    margin: auto;
    padding: 10px;

    & > .operate {
        display: flex;
        align-items: flex-start;
        width: 100%;
        height: 22px;

        .item {
            width: 13px;
            height: 13px;
            border-radius: 50%;
            margin-right: 8px;
            display: flex;
            justify-content: center;
            align-items: center;

            &.close {
                background: #f06c5a;
            }
            &.mini {
                background: #f9be45;
            }
            &.max {
                background: #61c350;
            }
        }
    }
}
</style>
