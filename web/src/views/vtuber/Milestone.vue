<script setup lang="ts">
import moment from "moment";
import { useVtuberStore } from "@/stores";
const props = defineProps<{ vid: number }>();
const store = useVtuberStore();
store.listMilestones(props.vid);

const getBetweenDays = (d: Date): string => {
  const days = moment().diff(moment(d), "day");
  if (days > 0) {
    return `已经${days}天`;
  } else {
    return `还有${days}天`;
  }
};
</script>
<template>
  <el-timeline>
    <el-timeline-item
      v-for="m in store.milestones"
      :timestamp="`${moment(m.date).format('YYYY-MM-DD')} ${getBetweenDays(
        m.date
      )}`"
      placement="top"
    >
      <el-card>
        <h4>{{ m.event }}</h4>
        <p>{{ m.description }}</p>
      </el-card>
    </el-timeline-item>
  </el-timeline>
</template>
