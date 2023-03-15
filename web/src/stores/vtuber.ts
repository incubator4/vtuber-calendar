import { defineStore } from "pinia";
import request from "@/tools/axios";

export const useVtuberStore = defineStore("vtuber", () => {
  const vtubers = ref<Array<IVtuber>>([]);

  const currentVtuber = ref<IVtuber>();

  const milestones = ref<IMilestone[]>([]);

  const fetchAll = () => {
    return request.get<IVtuber[]>("/vtubers").then((res) => {
      vtubers.value = res.data;
    });
  };

  const getVtuber = (uid: number) => {
    return request.get<IVtuber>(`/vtubers/${uid}`).then((res) => {
      currentVtuber.value = res.data;
    });
  };

  const listMilestones = (vid: number) => {
    return request.get<IMilestone[]>(`/milestones/${vid}`).then((res) => {
      milestones.value = res.data;
    });
  };

  return {
    vtubers,
    currentVtuber,
    milestones,
    fetchAll,
    getVtuber,
    listMilestones,
  };
});
