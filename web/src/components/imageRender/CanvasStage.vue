<script setup lang="ts">
import type { Dictionary } from "lodash";
import moment from "moment";
import type konva from "konva";

const props = defineProps<{
  image?: HTMLImageElement;
  config: ImageRenderConfig;
  size: {
    width: number;
    height: number;
    scale?: {
      x: number;
      y: number;
    };
  };
  data: Dictionary<ICalendar[]>;
  noEvent: string;
}>();

const onSave = () => {
  const s = stage.value.getNode() as unknown as konva.Stage;
  const dataURL = s.toDataURL();
  const link = document.createElement("a");
  link.download = "image.png";
  link.href = dataURL;
  document.body.appendChild(link);
  link.click();
};

defineExpose({
  onSave,
});

const timeFormat = (d: Date) => {
  const m = moment(d);
  return m.format("HH:mm");
};

const stage = ref();
</script>
<template>
  <v-stage ref="stage" :config="size">
    <v-layer ref="layer">
      <v-image
        ref="image"
        :config="{
          image,
        }"
      />
    </v-layer>
    <v-layer :config="{ ...config.text_offset }">
      <v-group
        v-for="(subGroup, gIndex) in config.text_group"
        :config="{
          x: config.text_group_offset.x * gIndex,
          y: config.text_group_offset.y * gIndex,
        }"
      >
        <v-group
          v-for="(group, index) in subGroup"
          :config="{
            x: config.row.x * +index,
            y: config.row.y * +index,
          }"
        >
          <v-group v-if="data[group]">
            <v-text
              v-for="(cal, i) in data[group]"
              :config="{
                text: timeFormat(cal.start_time) + cal.title,
                fontSize: config.font.size,
                fontFamily: config.font.family,
                fontStyle: config.font.style,
                stroke: config.font.stroke,
                strokeWidth: config.font.stroke_width,
                fill: config.font.color,
                x: config.col.x * i,
                y: config.col.y * i,
              }"
            ></v-text>
          </v-group>
          <v-group v-else>
            <v-text
              :config="{
                text: noEvent,
                align: 'center',
                fontSize: config.font.size,
                fontFamily: config.font.family,
                fontStyle: config.font.style,
                stroke: config.font.stroke,
                strokeWidth: config.font.stroke_width,
                fill: config.font.color,
              }"
            ></v-text>
          </v-group>
        </v-group>
      </v-group>
    </v-layer>
  </v-stage>
</template>
