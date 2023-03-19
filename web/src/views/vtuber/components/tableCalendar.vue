<script setup lang="ts">
import moment from "moment";
import EditDialog from "./EditDialog.vue";
import ImageRender from "@/components/imageRender/index.vue";
import { Hide, View } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { useVtuberStore, useCalendarStore } from "@/stores";
const vtuberStore = useVtuberStore();
const calendarStore = useCalendarStore();

const emits = defineEmits(["update"]);

const defaultEvent: ICalendar = {
  title: "",
  start_time: moment().toDate(),
  end_time: moment().add(2, "h").toDate(),
  cid: vtuberStore.currentVtuber?.id as number,
  tag_id: 0,
  is_active: true,
};

let eventModel = reactive<ICalendar>(defaultEvent);
const isReverse = ref(true);
const renderPanel = ref(false);
const dialogVisible = ref(false);
const pickDateRange = ref<[Date, Date]>([
  moment().startOf("week").toDate(),
  moment().endOf("week").toDate(),
]);
const size = ref<"default" | "large" | "small">("default");

const shortcuts = [
  {
    text: "本周",
    value: () => {
      return [
        moment().startOf("week").toDate(),
        moment().endOf("week").toDate(),
      ];
    },
  },
  {
    text: "本月",
    value: () => {
      return [
        moment().startOf("month").toDate(),
        moment().endOf("month").toDate(),
      ];
    },
  },
  {
    text: "本季度",
    value: () => {
      return [
        moment().startOf("quarter").toDate(),
        moment().endOf("quarter").toDate(),
      ];
    },
  },
  {
    text: "本年",
    value: () => {
      return [
        moment().startOf("year").toDate(),
        moment().endOf("year").toDate(),
      ];
    },
  },
  {
    text: "下周",
    value: () => {
      return [
        moment().startOf("week").add(1, "week").toDate(),
        moment().endOf("week").add(1, "week").toDate(),
      ];
    },
  },
];

const tagName = (tag_id: number) => {
  const tag = calendarStore.tags.find((t) => t.id === tag_id);
  return tag ? tag.name : "-";
};

const tableData = computed(() => {
  const data = calendarStore.calendars
    .sort((a, b) => (a.start_time > b.start_time ? 1 : -1))
    .filter((event, index) => {
      const date = new Date(event.start_time);
      const [from, to] = pickDateRange.value;
      const res =
        date.getTime() > from.getTime() && date.getTime() < to.getTime();
      return res;
    });
  return isReverse.value ? data.reverse() : data;
});

const tableDataReverse = computed(() => tableData.value.reverse());

const onNew = () => {
  eventModel = {
    ...defaultEvent,
  };
  const d = moment();
  eventModel.cid = vtuberStore.currentVtuber?.id as number;
  eventModel.start_time = d.toDate();
  eventModel.end_time = d.add(2, "h").toDate();
  dialogVisible.value = true;
};

const onEdit = (event: ICalendar) => {
  eventModel = {
    ...event,
  };
  dialogVisible.value = true;
};

const onDelete = (id: number) => {
  calendarStore.deleteCalendar(id).then(() => {
    ElMessage({
      message: "删除成功",
      type: "success",
    });
    emits("update");
  });
};

const onRender = () => {
  renderPanel.value = true;
};
</script>
<template>
  <el-row style="margin-top: 20px; margin-bottom: 5px">
    <el-col :sm="4" :xs="8">
      <el-form-item label="倒序">
        <el-switch v-model="isReverse"></el-switch>
      </el-form-item>
    </el-col>
    <el-col :sm="8" :xs="16">
      <el-button type="primary" @click="onNew">新增</el-button>
      <el-button @click="onRender">渲染</el-button>
    </el-col>
    <el-col :sm="8" :xs="24">
      <el-date-picker
        v-model="pickDateRange"
        type="daterange"
        unlink-panels
        range-separator="To"
        start-placeholder="Start date"
        end-placeholder="End date"
        :shortcuts="shortcuts"
        :size="size"
      />
    </el-col>
  </el-row>

  <el-table :data="tableData" style="width: 100%">
    <el-table-column label="Active" width="70">
      <template #default="{ row }">
        <el-icon v-if="row.is_active"><View /></el-icon>
        <el-icon v-else><Hide /></el-icon>
      </template>
    </el-table-column>
    <el-table-column label="Datetime" width="180">
      <el-table-column label="Date" width="120">
        <template #default="{ row }">
          {{ moment(row.start_time).format("YYYY-MM-DD") }}
        </template>
      </el-table-column>
      <el-table-column label="Time" width="80">
        <template #default="{ row }">
          {{ moment(row.start_time).format("HH:mm") }}
        </template>
      </el-table-column>

      <el-table-column label="Duration" width="100">
        <template #default="{ row }">
          {{ moment(row.end_time).diff(moment(row.start_time), "hour") }}小时
        </template>
      </el-table-column>
    </el-table-column>

    <el-table-column prop="title" label="Title" />
    <el-table-column prop="tag_id" label="Tag" width="80">
      <template #default="{ row }">
        <el-tag> {{ tagName(row.tag_id) }} </el-tag>
      </template>
    </el-table-column>
    <el-table-column fixed="right" label="Action" width="180">
      <template #default="{ row }">
        <el-button
          type="primary"
          @click="
            () => {
              onEdit(row);
            }
          "
        >
          编辑</el-button
        >
        <el-button
          type="danger"
          @click="
            () => {
              onDelete(row.id);
            }
          "
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
  <EditDialog
    v-model="dialogVisible"
    :event="eventModel"
    @update="
      () => {
        emits('update');
      }
    "
  >
  </EditDialog>
  <el-dialog v-if="renderPanel" title="预览" width="90%" v-model="renderPanel">
    <ImageRender :data="tableDataReverse" />
  </el-dialog>
</template>
