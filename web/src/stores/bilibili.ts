import { defineStore } from "pinia";
import request from "@/tools/axios";

export const useVtbmoeStore = defineStore("vtb-moe", () => {
  const user = ref<BiliUser>();

  const clearUserInfo = () => {
    user.value = undefined;
  };

  const getUserInfo = (uid: number) =>
    request.get<BiliUser>(`/bilibili/${uid}`, {}).then((resp) => {
      user.value = resp.data;
    });

  return { user, getUserInfo, clearUserInfo };
});
