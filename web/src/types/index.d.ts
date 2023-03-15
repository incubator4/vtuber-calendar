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
  cid: number;
  tag_id: number;
  is_active: boolean;
}

interface ITag {
  id: number;
  name: string;
}

interface ImageConfig {
  url: string;
  width: number;
  height: number;
}

interface OffSet {
  x: number;
  y: number;
}

interface OffSetConfig extends OffSet {
  prefix: string;
  suffix: string;
}

interface FontConfig {
  family: string;
  size: number;
  color: string;
  stroke: string;
  layout: string;
  style: string;
}

interface ImageRenderConfig {
  id?: number;
  name: string;
  image: ImageConfig;
  text_offset: OffSet;
  text_group: Array<Array<number>>;
  text_group_offset: OffSet;
  row: OffSetConfig;
  col: OffSetConfig;
  font: FontConfig;
}

type VtuberCalendar = ICalendar & IVtuber;

interface IMilestone {
  id: number;
  vid: number;
  date: Date;
  event: string;
  description: string;
  is_dispaly: boolean;
  is_deleted: boolean;
}
