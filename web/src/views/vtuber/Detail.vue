<script setup lang="ts">
import {
  useVtuberStore,
  useVtbmoeStore,
  useCalendarStore,
  useImageRenderConfig,
} from "@/stores";
import { useDark } from "@vueuse/core";
import { useScreen } from "vue-screen";

import Milestone from "./Milestone.vue";
import type { TabsPaneContext } from "element-plus";
import banner from "./components/banner.vue";
import tableCalendar from "./components/tableCalendar.vue";
import Info from "./components/Info.vue";

const props = defineProps({ uid: String });

const vtuberStore = useVtuberStore();
const calendarStore = useCalendarStore();
const configStore = useImageRenderConfig();
const vmStore = useVtbmoeStore();
const screen = useScreen();
const isDark = useDark();

const loadData = () => {
  const uid = +(props.uid as string);

  calendarStore.listCalendar({ uid: [uid.toString()] }).then(() => {
    state.total = calendarStore.calendars.length;
  });
  vtuberStore.getVtuber(uid).then(() => {
    const id = vtuberStore.currentVtuber?.id as number;
    configStore.list(id);
  });
  calendarStore.listTags();
  vmStore.getDetail(uid);
};

onMounted(() => {
  vmStore.clearDetail();
  calendarStore.clearCalendar();
  loadData();
});

watch(
  () => props.uid,
  (newUID, oldUID) => {
    if (newUID !== "" && newUID !== undefined && newUID !== oldUID) {
      const uid = +(newUID as string);
      loadData();
    }
  }
);

const state = reactive({
  page: 1,
  limit: 10,
  total: 0,
});

const activeName = ref("Calendar");

const handleClick = (tab: TabsPaneContext, event: Event) => {
  console.log(tab.props.name, event);
};
</script>

<template>
  <main>
    <banner></banner>
    <div style="padding: 20px; min-height: 50vh">
      <el-tabs v-model="activeName" type="border-card" @tab-click="handleClick">
        <el-tab-pane label="日程" name="Calendar">
          <table-calendar @update="loadData"></table-calendar>
        </el-tab-pane>
        <el-tab-pane label="里程碑" name="milestore">
          <Milestone
            v-if="vtuberStore.currentVtuber?.id"
            :vid="vtuberStore.currentVtuber.id"
          />
        </el-tab-pane>
        <el-tab-pane label="个人信息" name="third">
          <info />
        </el-tab-pane>
      </el-tabs>
    </div>
  </main>
</template>

<style scoped>
.button {
  padding: 10px 20px;
  border: 1px solid #ddd;
  color: #333;
  background-color: #fff;
  border-radius: 4px;
  font-size: 14px;
  font-family: "微软雅黑", arail;
  cursor: pointer;
}
</style>
