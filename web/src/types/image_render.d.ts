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
  stroke_width: number;
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
