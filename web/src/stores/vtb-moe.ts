import { defineStore } from "pinia";
import axios from "axios";

export const useVtbmoeStore = defineStore("vtb-moe", () => {
  const vtuberDetail = ref<VtuberDetail>();

  const getDetail = (uid: number) =>
    axios
      .get<VtuberDetail>(`https://api.vtbs.moe/v1/detail/${uid}`)
      .then((res) => {
        vtuberDetail.value = res.data;
      });

  const clearDetail = () => {
    vtuberDetail.value = undefined;
  };

  return { vtuberDetail, getDetail, clearDetail };
});
