<script lang="ts" setup>
import { ChatLineRound, Male } from "@element-plus/icons-vue";
import { useStatusStore } from "@/stores";
import type { TrafficPackage } from "tencentcloud-sdk-nodejs/tencentcloud/services/lighthouse/v20200324/lighthouse_models";

const packageSet = computed<TrafficPackage>(() => {
  const p: TrafficPackage = {
    TrafficPackageId: "",
    TrafficUsed: 0,
    TrafficPackageTotal: 0,
    TrafficPackageRemaining: 0,
    TrafficOverflow: 0,
    StartTime: "",
    EndTime: "",
    Deadline: "",
    Status: "",
  };
  if (!store.traffic || store.traffic.TotalCount === 0) {
    return p;
  }

  const { TrafficPackageSet } = store.traffic.InstanceTrafficPackageSet[0];

  if (!TrafficPackageSet || TrafficPackageSet.length === 0) {
    return p;
  }

  return TrafficPackageSet[0];
});

const loading = ref(true);

const store = useStatusStore();
store.getTraffic().then(() => {
  loading.value = false;
});

const options: Array<{
  unit: "GB" | "MB" | "KB";
  format: (input: number) => number;
}> = [
  {
    unit: "GB",
    format: (input) => input / (1024 * 1024 * 1024),
  },
  {
    unit: "MB",
    format: (input) => input / (1024 * 1024),
  },
];

const currentOption = ref<"GB" | "MB" | "KB">("GB");

const currentFormat = computed(() => {
  const option = options.find((o) => o.unit === currentOption.value);
  return option ? option.format : (input: number) => input;
});
</script>

<template>
  <div style="height: 80vh; text-align: center; padding: 100px">
    <el-card>
      <template #header>
        <el-space>
          流量详情
          <el-select style="width: 80px" v-model="currentOption" placeholder="">
            <el-option
              v-for="{ unit } in options"
              :label="unit"
              :value="unit"
            ></el-option>
          </el-select>
        </el-space>
      </template>
      <el-skeleton v-if="loading" :rows="2" animated />
      <el-row v-else>
        <el-col :span="6">
          <el-statistic
            title="已用流量"
            :value="currentFormat(packageSet.TrafficUsed)"
          >
            <template #suffix>/{{ currentOption }}</template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic
            title="总流量"
            :value="currentFormat(packageSet.TrafficPackageTotal)"
          >
            <template #suffix>/{{ currentOption }}</template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic
            title="剩余"
            :value="currentFormat(packageSet.TrafficPackageRemaining)"
          >
            <template #suffix>/{{ currentOption }}</template>
          </el-statistic>
        </el-col>
        <el-col :span="6">
          <el-statistic
            title="使用百分比"
            :precision="2"
            :value="
              (packageSet.TrafficUsed * 100) / packageSet.TrafficPackageTotal
            "
          >
            <template #suffix>%</template>
          </el-statistic>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<style scoped>
.el-col {
  text-align: center;
}
</style>
