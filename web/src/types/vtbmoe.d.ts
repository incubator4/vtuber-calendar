interface LastLive {
  online: number;
  time: number;
}

interface VtuberDetail {
  mid: number;
  uuid: string;
  uname: string;
  video: number;
  roomid: number;
  sign: string;
  notice: string;
  face: string;
  rise: number;
  topPhoto: string;
  archiveView: number;
  follower: number;
  liveStatus: number;
  recordNum: number;
  guardNum: number;
  lastLive: LastLive;
  guardChange: number;
  guardType: number[];
  online: number;
  title: string;
  time: number;
  liveStartTime: number;
}
