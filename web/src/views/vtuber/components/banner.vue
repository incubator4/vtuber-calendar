<script setup lang="ts">
import { useVtbmoeStore } from "@/stores";

const vmStore = useVtbmoeStore();
</script>
<template>
  <div v-if="vmStore.user">
    <div class="banner-cover">
      <img class="bannner" style="width: 100%" :src="vmStore.user.top_photo" />
      <div class="context">
        <el-container style="max-width: 600px">
          <el-aside width="70px">
            <el-avatar
              style="margin-top: 30px"
              size="large"
              :src="vmStore.user.face"
            />
            <el-tag
              v-if="
                vmStore.user.live_room.roomStatus &&
                vmStore.user.live_room.liveStatus === 1
              "
            >
              <el-link
                target="_blank"
                :href="`https://live.bilibili.com/${vmStore.user.live_room.roomid}`"
                >直播中</el-link
              >
            </el-tag>
          </el-aside>
          <el-main style="color: white">
            <div class="card">
              <h2 style="font-weight: bold">
                {{ vmStore.user.name }}
              </h2>
              {{ vmStore.user.sign }}
            </div>
          </el-main>
        </el-container>
      </div>
    </div>
  </div>
  <el-skeleton v-else style="width: 100%" :row="10"> </el-skeleton>
</template>
<style lang="less" scoped>
.banner-cover {
  position: relative;
  overflow: hidden;
  min-height: 120px;

  .bannner {
    position: absolute;
    width: 1000px; /* 设置宽度为100% */
    height: 150px; /* 高度自适应 */
    object-fit: cover; /* 关键属性, 取值包括：fill | contain | cover | scale-down | none */
    object-position: center; /* 将图片位置设置为中心部位 */
  }
  .context {
    height: 100%;
    margin-left: 20px;
    margin-top: auto;
    margin-bottom: auto;
  }
}

.card {
  margin-top: 5px;
  border-radius: 4px;
  padding: 4px;
  width: 100%;
  /* min-height: 40px; */

  background-color: rgba(1, 1, 1, 0.5);
}
</style>
