<template>
  <a-list item-layout="vertical" size="large" :pagination="pagination" :data-source="listData">
    <div slot="footer"><b>ant design vue</b> footer part</div>
    <a-list-item slot="renderItem" key="item.title" slot-scope="item">
      <template v-for="{ type, text } in actions" slot="actions">
        <span :key="type">
          <a-icon :type="type" style="margin-right: 8px" />
          {{ text }}
        </span>
      </template>
      <img
        slot="extra"
        width="272"
        alt="logo"
        src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png"
      />
      <a-list-item-meta :description="item.description">
        <a slot="title" :href="item.href">{{ item.title }}</a>
        <a-avatar slot="avatar" :src="item.avatar" />
      </a-list-item-meta>
      {{ item.content }}
    </a-list-item>
  </a-list>
</template>
<script>
const listData = [];
for (let i = 0; i < 23; i += 1) {
  listData.push({
    href: 'https://www.lshblog.com/',
    title: `Golang Gin Vue Blog ${i}`,
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    description:
      'Golang Gin Vue Blog  description.',
    content:
      'Golang Gin Vue Blog  content.',
  });
}

export default {
  data() {
    return {
      listData,
      pagination: {
        onChange: (page) => {
          console.log(page);
        },
        pageSize: 3,
      },
      actions: [
        { type: 'star-o', text: '137' },
        { type: 'like-o', text: '101' },
        { type: 'message', text: '67' },
      ],
    };
  },
};
</script>
<style></style>
