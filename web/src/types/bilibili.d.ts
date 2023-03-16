interface BiliUser {
  mid: number;
  name: string;
  sex: string;
  face: string;
  sign: string;
  rank: number;
  level: number;
  jointime: number;
  moral: number;
  silence: number;
  coins: number;
  fans_badge: boolean;
  official: {
    role: number;
    title: string;
    desc: string;
    type: number;
  };
  vip: {
    type: number;
    status: number;
    due_date: number;
    vip_pay_type: number;
    theme_type: number;
    label: {
      path: string;
      text: string;
      label_theme: string;
      text_color: string;
      bg_style: number;
      bg_color: string;
      border_color: string;
    };
    avatar_subscript: number;
    nickname_color: string;
    role: number;
    avatar_subscript_url: string;
  };
  pendant: {
    pid: number;
    name: string;
    image: string;
    expire: number;
    image_enhance: string;
    image_enhance_frame: string;
  };
  nameplate: {
    nid: number;
    name: string;
    image: string;
    image_small: string;
    level: string;
    condition: string;
  };
  is_followed: boolean;
  top_photo: string;
  sys_notice: {
    id: number;
    content: string;
    url: string;
    notice_type: number;
    icon: string;
    text_color: string;
    bg_color: string;
  };
  live_room: {
    roomStatus: number;
    roundStatus: number;
    liveStatus: number;
    url: string;
    title: string;
    cover: string;
    online: number;
    roomid: number;
    broadcast_type: number;
    online_hidden: number;
  };
  birthday: string;
  school: {
    name: string;
  };
  profession: {
    name: string;
  };
  series: {
    user_upgrade_status: number;
    show_upgrade_window: boolean;
  };
}
