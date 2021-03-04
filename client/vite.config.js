import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import styleImport from "vite-plugin-style-import";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    styleImport({
      libs: [
        {
          libraryName: "ant-design-vue",
          esModule: true,
          resolveStyle: name => `ant-design-vue/es/${name}/style/css`
        }
      ]
    })
  ],
  resolve: {
    alias: {
      "/@/": `${resolve(__dirname, ".", "src")}/`
    }
  }
});
