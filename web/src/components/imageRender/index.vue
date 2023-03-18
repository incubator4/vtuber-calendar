<script setup lang="ts">
import { groupBy } from "lodash";
import { useImageRenderConfig } from "@/stores";

import CanvasStage from "./CanvasStage.vue";

const props = defineProps<{ data: ICalendar[] }>();
const image2 = ref<HTMLImageElement>();

const configStore = useImageRenderConfig();

const renderStage = ref();

const stageSize = reactive({
  width: 1920,
  height: 1180,
});

const percent = ref(20);

const computeSize = computed(() => {
  const ratio = (stageSize.width * percent.value) / 50 / stageSize.width;
  return {
    width: stageSize.width * ratio,
    height: stageSize.height * ratio,
    scale: {
      x: ratio,
      y: ratio,
    },
  };
});

const resize = () => {
  const width = (stageSize.width * percent.value) / 50;
  const ratio = width / stageSize.width;
};

const computedData = computed(() =>
  groupBy(
    props.data
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

const load = () => {
  const image = new window.Image();

  image.setAttribute("crossOrigin", "anonymous");
  image.src = currentConfig.value.image.url;
  stageSize.height = currentConfig.value.image.height;
  stageSize.width = currentConfig.value.image.width;
  image.onload = () => {
    // set image only when it is loaded
    const width = image.width;
    const ratio = width / image.width;
    image.width = width;
    image.height = image.height * ratio;
    image2.value = image;
    render.value = true;
  };
};

const onSave = () => {
  const s = renderStage.value;
  s.onSave();
};

let currentConfig = ref<ImageRenderConfig>({
  name: "",
  image: {
    url: "",
    width: 0,
    height: 0,
  },
  text_offset: {
    x: 0,
    y: 0,
  },
  text_group: [[]],
  text_group_offset: {
    x: 0,
    y: 0,
  },
  row: {
    x: 0,
    y: 0,
    prefix: "",
    suffix: "",
  },
  col: {
    x: 0,
    y: 0,
    prefix: "",
    suffix: "",
  },
  font: {
    family: "",
    size: 0,
    color: "",
    stroke: "",
    stroke_width: 2,
    style: "",
    layout: "",
  },
});

const currentConfigId = ref();

const configChange = (id: number) => {
  const config = configStore.configs.find((c) => c.id === id);
  if (config) {
    currentConfig.value = { ...config };
  }
  load();
};

const curFontStyle = computed<Array<String>>({
  get() {
    return currentConfig.value.font.style.split(" ");
  },
  set(vals) {
    currentConfig.value.font.style = vals.join(" ");
  },
});

const activeNames = ref(["1"]);
const handleChange = (val: string[]) => {
  console.log(val);
};

const defaultNoEvent = ref("无日程");
</script>
<template>
  <el-container>
    <el-header>
      <el-row>
        <el-col :span="6">
          <el-select
            v-model="currentConfigId"
            @change="configChange"
            placeholder=""
          >
            <el-option
              v-for="config in configStore.configs"
              :label="config.name"
              :value="config.id"
            >
            </el-option>
          </el-select>
        </el-col>
        <el-col :span="0"></el-col>
        <el-col :span="16">
          <el-slider v-model="percent" @change="resize" show-input />
        </el-col>
      </el-row>
    </el-header>
    <el-container>
      <el-aside width="220px">
        <el-collapse v-model="activeNames" @change="handleChange">
          <el-collapse-item title="字体设置" name="1">
            <el-form-item label="字体">
              <el-select v-model="currentConfig.font.family" placeholder="">
                <el-option label="默认" value="sans-serif"></el-option>
                <el-option label="灵动指书" value="灵动指书" />
                <el-option
                  label="SmileySans-Oblique"
                  value="SmileySans-Oblique"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item label="字号">
              <el-input-number
                style="width: 100%"
                v-model="currentConfig.font.size"
                placeholder=""
              />
            </el-form-item>
            <el-row>
              <el-col :span="12">
                <el-form-item label="颜色">
                  <el-color-picker
                    v-model="currentConfig.font.color"
                    show-alpha
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item label="边框">
                  <el-color-picker
                    v-model="currentConfig.font.stroke"
                    show-alpha
                  />
                </el-form-item>
              </el-col>
              <el-col :span="12">
                <el-form-item v-show="false" label="方向">
                  <el-select v-model="currentConfig.font.layout" placeholder="">
                    <el-option label="横向" value="vertical"></el-option>
                    <el-option label="纵向" value="horizontal"></el-option>
                  </el-select>
                </el-form-item>
              </el-col>
            </el-row>
            <el-form-item label="样式">
              <el-input-number
                style="width: 100%"
                :max="10"
                v-model="currentConfig.font.stroke_width"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="样式">
              <el-checkbox-group v-model="curFontStyle">
                <el-checkbox label="bold" />
                <el-checkbox label="italic" />
              </el-checkbox-group>
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="文本设置" name="2">
            <template #title> 文本设置 </template>
            <el-form-item label="无日程设置">
              <el-input v-model="defaultNoEvent" placeholder=""></el-input>
            </el-form-item>
            <el-form-item label="横向偏移">
              <el-input-number
                size="small"
                v-model="currentConfig.text_offset.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向偏移">
              <el-input-number
                size="small"
                v-model="currentConfig.text_offset.y"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.text_group_offset.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.text_group_offset.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="每日间距设置" name="3">
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.row.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.row.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
          <el-collapse-item title="事件间距设置" name="4">
            <el-form-item label="横向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.col.x"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="纵向间距">
              <el-input-number
                size="small"
                v-model="currentConfig.col.y"
                placeholder=""
              />
            </el-form-item>
          </el-collapse-item>
        </el-collapse>
        <el-button type="primary" @click="onSave">保存</el-button>
      </el-aside>
      <el-main>
        <div style="text-align: center">
          <CanvasStage
            :image="image2"
            :size="computeSize"
            :config="currentConfig"
            :no-event="defaultNoEvent"
            :data="computedData"
          />
          <CanvasStage
            ref="renderStage"
            v-show="false"
            :image="image2"
            :size="stageSize"
            :config="currentConfig"
            :no-event="defaultNoEvent"
            :data="computedData"
          />
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<style scoped>
.title-tooltip {
  width: 110px;
}
</style>
