<script setup lang="ts">
import { useVtuberStore } from "@/stores";
import Avatar from "@/components/icons/Avatar.vue";
import { useScreen } from "vue-screen";
import useClipboard from "vue-clipboard3";
import { useRouter } from "vue-router";
import {
  Calendar,
  ChatDotRound,
  CopyDocument,
  More,
  Monitor,
} from "@element-plus/icons-vue";

const { toClipboard } = useClipboard();
const router = useRouter();

const icsLink = (uid: number) =>
  `webcal://${window.location.host}/api/ics/${uid}`;

const open = (uid: number) => {
  window.open(icsLink(uid));
};

const onLiveRoom = (live_id: number) => {
  window.open(`https://live.bilibili.com/${live_id}`);
};

const onClipboard = (uid: number) => {
  toClipboard(icsLink(uid));
  alert("日历链接复制成功");
};

const onQQGroup = (group: number) => {
  toClipboard("" + group);
  alert(`粉丝群 ${group} 复制成功`);
};

const onDetail = (uid: number) => {
  router.push({
    name: "VtuberDetail",
    params: {
      uid,
    },
  });
};

const size = ref("");
const screen = useScreen();
const store = useVtuberStore();
store.fetchAll();
</script>
<template>
  <div style="max-width: 1000px; margin: 0 auto">
    <div :style="{ height: `${screen.height}px`, width: 'auto' }">
      <el-scrollbar>
        <el-row>
          <template
            v-for="{ uid, name, id, live_id, main_color } in store.vtubers"
            :key="id"
            class="scrollbar-demo-item"
            :style="{
              color: main_color,
              border: `1px solid ${main_color}`,
            }"
          >
            <el-col :xs="24" :sm="12" :lg="8" :xxl="6">
              <el-card style="margin: 10px">
                <el-row>
                  <el-col :span="8">
                    <a
                      :href="`https://space.bilibili.com/${uid}`"
                      style="font-size: 16px"
                      target="_blank"
                      class="title"
                      ><Avatar class="avatar" :uid="uid"></Avatar
                    ></a>

                    <el-button style="width: 120px" :color="main_color">
                      <div style="font-size: 16px" class="title">
                        {{ name }}
                      </div>
                    </el-button>
                  </el-col>
                  <el-col :span="16">
                    <ul style="overflow: auto">
                      <li class="scrollbar-demo-item">
                        <el-button
                          @click="
                            () => {
                              open(uid);
                            }
                          "
                          class="button"
                          :icon="Calendar"
                        >
                          订阅
                        </el-button>
                      </li>
                      <li class="scrollbar-demo-item">
                        <el-button
                          class="button"
                          :icon="CopyDocument"
                          @click="() => onClipboard(uid)"
                        >
                          复制
                        </el-button>
                      </li>
                      <li class="scrollbar-demo-item">
                        <el-button
                          class="button"
                          @click="() => onLiveRoom(live_id)"
                          :icon="Monitor"
                          >直播间</el-button
                        >
                      </li>
                      <li class="scrollbar-demo-item">
                        <el-button
                          class="button"
                          @click="() => onDetail(uid)"
                          :icon="More"
                          >更多信息</el-button
                        >
                      </li>
                    </ul>
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </template>
        </el-row>
      </el-scrollbar>
    </div>
  </div>
</template>
<style scoped>
.avatar {
  width: 120px;
  height: 120px;
}

.title {
  /* height: 60px; */
  /* padding: 0; */
  color: white;
}
.scrollbar-demo-item {
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  /* height: 180px; */
  /* min-width: 350px; */
  max-width: 350px;
  margin: 5px;
  text-align: center;
  border-radius: 4px;
  /* background: var(--el-color-primary-light-9); */
  color: var(--el-color-primary);
}

.cell-item {
  display: flex;
  min-width: 60px;
  text-align: center;
  align-items: center;
}
.cell-context {
  text-align: center;
}
.button {
  width: 160px;
}
</style>
