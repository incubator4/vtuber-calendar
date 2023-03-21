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
  tag_id: number;
  is_active: boolean;
}

interface ITag {
  id: number;
  name: string;
}

interface VtuberCalendar extends ICalendar, IVtuber {
  tags: Array<number>;
}

interface IMilestone {
  id: number;
  vid: number;
  date: Date;
  event: string;
  description: string;
  is_dispaly: boolean;
  is_deleted: boolean;
}
