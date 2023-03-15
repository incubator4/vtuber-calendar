<script setup lang="ts">
import { useVtuberStore, useCalendarStore } from "@/stores";
import moment from "moment";
import { cloneDeep } from "lodash";
import { ElMessageBox, ElMessage } from "element-plus";
import { useScreen } from "vue-screen";
import type { FormInstance, FormRules } from "element-plus";

const screen = useScreen();
const vtuberStore = useVtuberStore();
const calendarStore = useCalendarStore();

const ruleFormRef = ref<FormInstance>();

const props = defineProps<{
  modelValue: boolean;
  event: ICalendar;
}>();

const emit = defineEmits(["update:modelValue", "update"]);

const minusHour = (start: Date, end: Date) => {
  const diffTime = end.getTime() - start.getTime();
  return Math.floor(diffTime / (1000 * 60 * 60));
};

const calcValue = computed<boolean>({
  get() {
    return props.modelValue;
  },
  set(val) {
    emit("update:modelValue", val);
  },
});

const editEvent = ref<ICalendar>(props.event);

const rules = reactive<FormRules>({
  title: [{ required: true, message: "Please input Title", trigger: "blur" }],
  start_time: [
    {
      type: "date",
      required: true,
      message: "Please pick a date",
      trigger: "change",
    },
  ],
});

watch(
  () => props.modelValue,
  (val) => {
    if (val) {
      editEvent.value = props.event;
      console.log(props.event);
      // editWeekCal.value = cloneDeep(currentWeekCal.value);
    }
  }
);

const beforeClose = (done: () => void) => {
  ElMessageBox.confirm("Are you sure you want to close this?")
    .then(() => {
      done();
    })
    .catch(() => {
      // catch error
    });
};

const submitForm = async (formEl: FormInstance | undefined) => {
  console.log(formEl);
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      const model = editEvent.value;

      const result = model.id
        ? calendarStore.updateCalendar(model.id, model)
        : calendarStore.createCalendar(model);
      result.then(() => {
        ElMessage({
          message: model.id ? "更新成功" : "新建成功",
          type: "success",
        });
        emit("update");
        handleClose();
      });
    } else {
      console.log("error submit!", fields);
    }
  });
};

const handleClose = () => {
  emit("update:modelValue", false);
};
const formSize = ref("default");

const colGrid = {
  xs: 24,
  sm: 12,
};
</script>
<template>
  <el-dialog
    v-model="calcValue"
    :title="event.id ? '修改' : '新增'"
    :width="screen.width > 768 ? '600px' : '90%'"
  >
    <el-form
      ref="ruleFormRef"
      :model="editEvent"
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
      :size="formSize"
      status-icon
    >
      <el-form-item label="标题" prop="title">
        <el-input style="width: 90%" v-model="editEvent.title"></el-input>
      </el-form-item>
      <el-form-item label="类型">
        <el-select style="width: 90%" v-model="editEvent.tag_id" placeholder="">
          <el-option label="暂无" :value="0"></el-option>
          <el-option
            v-for="tag in calendarStore.tags"
            :label="tag.name"
            :value="tag.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="时间" prop="start_time">
        <el-date-picker
          v-model="editEvent.start_time"
          type="datetime"
          placeholder="Pick a date"
          style="width: 90%"
        />
      </el-form-item>
      <el-form-item label="时长">
        <el-select
          style="width: 90%"
          :model-value="
            minusHour(
              new Date(editEvent.start_time),
              new Date(editEvent.end_time)
            )
          "
          @change="(val: number) => {
                editEvent.end_time = moment(editEvent.start_time).add(val, 'hour').toDate()
              }"
          placeholder=""
        >
          <el-option
            v-for="item in 8"
            :label="item + '小时'"
            :value="item"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="启用" prop="is_active">
        <el-switch v-model="editEvent.is_active"></el-switch>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button type="danger" @click="handleClose">取消</el-button>
        <el-button
          type="primary"
          @click="
            () => {
              submitForm(ruleFormRef);
            }
          "
        >
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<style scoped>
.event {
  border: 2px;
  border-radius: 2px;
  margin-top: 5px;
}
</style>
