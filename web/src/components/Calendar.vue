<script setup lang="tsx">
import { useVtuberStore, useCalendarStore } from "@/stores";
import moment from "moment";
import { groupBy } from "lodash";
import Calendar from "./Calendar/index.vue";
import { useLocalStorage } from "@vueuse/core";
import { useRoute } from "vue-router";

const vtuberStore = useVtuberStore();
const calendarStore = useCalendarStore();
const route = useRoute();

const localStorage = useLocalStorage<Array<number>>("vtuber", []);

const getUID = (cid: number) =>
  (vtuberStore.vtubers.find((c) => c.id === cid) || { uid: 0 }).uid;

const startOfWeek = moment().startOf("week").toDate();
const endofWeek = moment().endOf("week").toDate();

const currentWeekRange = computed(() => {
  const dates = [];
  for (let i = 0; i <= 6; i++) {
    dates.push(moment(startOfWeek).add(i, "days"));
  }
  return dates;
});

const currentDay = ref(moment().days() === 0 ? 7 : moment().days());

vtuberStore.fetchAll().then(() => {
  checkAll.value = true;
  isIndeterminate.value = false;

  if (route.query["vtuber"]) {
    console.log(route.query["vtuber"]);
  } else {
    checkboxVtubers.value = localStorage.value;
  }
});

calendarStore.listCalendar({
  // start: moment(startOfWeek).format("yyyy-MM-DD"),
  end: moment(endofWeek).add(1, "days").format("yyyy-MM-DD"),
});

const currentWeekCal = computed(() =>
  groupBy(
    calendarStore.calendars
      // .filter(({ start_time }) => {
      //   const date = new Date(start_time);
      //   return date > startOfWeek && date < endofWeek;
      // })
      .sort((a, b) => (a.start_time > b.start_time ? 1 : -1)),
    (event) => {
      const index = new Date(event.start_time).getDay();
      return index === 0 ? 7 : index;
    }
  )
);

interface attr {
  customData: {
    time: number;
  };
}
const sortAttr = (a: attr, b: attr) =>
  a.customData.time > b.customData.time ? 1 : -1;

const matchWeek = (index: number) => {
  switch (index) {
    case 1:
      return "星期一";
    case 2:
      return "星期二";
    case 3:
      return "星期三";
    case 4:
      return "星期四";
    case 5:
      return "星期五";
    case 6:
      return "星期六";
    case 7:
      return "星期日";
  }
};

const checkboxVtubers = ref<Array<number>>([]);

const checkAll = ref(false);
const isIndeterminate = ref(true);
const handleCheckAllChange = (val: boolean) => {
  checkboxVtubers.value = val ? vtuberStore.vtubers.map((c) => c.id) : [];
  isIndeterminate.value = false;
};

const handleCheckedVtubersChange = (value: number[]) => {
  localStorage.value = value;
};
</script>

<template>
  <div class="wrapper">
    <div class="check-wrapper">
      <el-select
        style="width: 100%"
        v-model="checkboxVtubers"
        @change="handleCheckedVtubersChange"
        filterable
        multiple
        placeholder="Select"
      >
        <el-option
          v-for="c in vtuberStore.vtubers"
          :key="c.id"
          :label="c.name"
          :value="c.id"
        >
          <span style="float: left">{{ c.name }}</span>
          <span
            style="
              float: right;
              color: var(--el-text-color-secondary);
              font-size: 13px;
            "
            >{{ c.uid }}</span
          >
        </el-option>
      </el-select>
    </div>
    <Calendar
      :calendars="
        calendarStore.calendars.filter((cal) =>
          checkboxVtubers.includes(cal.vid)
        )
      "
    />
  </div>
</template>

<style scoped>
.wrapper {
  margin: 10px;
}

.check-wrapper {
  /* margin: 5px; */
  margin-bottom: 5px;
}
.carousel {
  margin-top: 40px;
}

.el-carousel__item h3 {
  /* color: #475669; */
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  /* background-color: #99a9bf; */
}

.el-carousel__item:nth-child(2n + 1) {
  /* background-color: #d3dce6; */
}
</style>
