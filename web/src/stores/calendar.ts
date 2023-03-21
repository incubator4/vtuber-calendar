import { defineStore } from "pinia";
import request from "@/tools/axios";
import qs from "qs";

export const useCalendarStore = defineStore("calendar", () => {
  const calendars = ref<Array<VtuberCalendar>>([]);
  const tags = ref<Array<ITag>>([]);

  const getCalendar = (uid: number) => {
    return request.get<VtuberCalendar[]>(`/cal/${uid}`).then((res) => {
      calendars.value = res.data;
    });
  };

  const updateCalendar = (id: number, model: ICalendar) => {
    return request.put(`/cal/${id}`, model);
  };

  const createCalendar = (model: ICalendar) => {
    return request.post("/cal", model);
  };

  const listCalendar = (params?: {
    start?: string;
    end?: string;
    uid?: Array<string>;
    all?: boolean;
  }) => {
    return request
      .get<VtuberCalendar[]>(`/cal`, {
        params,
        paramsSerializer: {
          serialize: (params) => {
            return qs.stringify(params, { arrayFormat: "repeat" });
          },
        },
      })
      .then((res) => {
        calendars.value = res.data;
      });
  };

  const deleteCalendar = (id: number) => request.delete(`/cal/${id}`);

  const listTags = () => {
    return request.get<ITag[]>("/tags").then((res) => {
      tags.value = res.data;
    });
  };

  const clearCalendar = () => {
    calendars.value = [];
  };

  return {
    calendars,
    tags,
    getCalendar,
    updateCalendar,
    createCalendar,
    clearCalendar,
    listCalendar,
    deleteCalendar,
    listTags,
  };
});
