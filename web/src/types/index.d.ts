interface IVtuber {
  id: number;
  name: string;
  uid: number;
  live_id: number;
  main_color: string;
}

interface ICalendar {
  id?: number;
  title: string;
  start_time: Date;
  end_time: Date;
  vid: number;
  is_active: boolean;
  tags: Array<ITag>;
}

interface ITag {
  id: number;
  name: string;
}

type VtuberCalendar = ICalendar & Vtuber;

interface IMilestone {
  id: number;
  vid: number;
  date: Date;
  event: string;
  description: string;
  is_dispaly: boolean;
  is_deleted: boolean;
}
