<template>
  <div class="about">
    <el-card class="card">
      <template #header>
        <h3>关于网站</h3>
      </template>

      <ul>
        <li>个人制作的第三方网站，主要功能是做日历订阅</li>
        <li>网站所有内容不代表开发者本人的主观意愿和倾向</li>
        <li>本网站没有任何商业和企划相关联，主播也是出于个人意愿合作</li>
      </ul>
    </el-card>
    <el-card class="card">
      <template #header>
        <h3>建站初衷</h3>
      </template>
      <ul>
        <li>早期是为了方便自己能够更好的获取主播的日程信息</li>
        <li>后面是为了管理和帮助各位DD们提供一个日历订阅网站</li>
        <li>目前方便主播自行编辑日程和渲染出周表，同步日程订阅</li>
      </ul>
    </el-card>

    <el-card class="card">
      <template #header>
        <h3>技术栈</h3>
      </template>
      <ul>
        <li>前端: Vue3 & Typescript</li>
        <li>后端: Gorm & Postgres</li>
        <li>版本控制工具: Github & Github Actions</li>
        <li>部署: Kubernetes</li>
        <li>数据监控: 腾讯云API & Prometheus</li>
      </ul>
      <br />
      意见和建议可以通过 QQ: 337376577 和
      <el-link
        href="https://github.com/incubator4/vtuber-calendar"
        target="_blank"
        >Github</el-link
      >联系我
    </el-card>
  </div>
</template>

<style scoped>
.about {
  height: 80vh;
  /* text-align: center; */
  padding: 50px;
  /* margin: 40px; */
}

.card {
  margin: 20px;
}
</style>
