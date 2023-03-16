<script setup lang="ts">
import { useVtbmoeStore } from "@/stores";

const vmStore = useVtbmoeStore();
</script>
<template>
  {{ vmStore.user }}
  <div v-if="vmStore.user">
    <img class="bannner" style="width: 100%" :src="vmStore.user.top_photo" />
    <div style="position: absolute; left: 40px; top: 20px">
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
  <el-skeleton v-else style="width: 100%" :row="10"> </el-skeleton>
</template>
<style scoped>
.bannner {
  width: 100%;
  position: relative;
  overflow: hidden;
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
