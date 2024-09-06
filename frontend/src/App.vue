<template>
  <div>
    <h2>BFS Array</h2>
    <input v-model="bfsInput" placeholder="Enter BFS sequence" />
    <button @click="generateTree">Generate Tree</button>

    <h3>Binary tree in JSON</h3>
    <pre>{{ treeJson }}</pre>

    <h3>Binary Tree</h3>
    <binary-tree :treeData="treeData"></binary-tree>
  </div>
</template>

<script>
import BinaryTree from './components/BinaryTree.vue';

export default {
  components: {
    BinaryTree
  },
  data() {
    return {
      bfsInput: "",
      treeJson: {},
      treeData: null
    };
  },
  methods: {
    async generateTree() {
      const bfsArray = this.bfsInput.split(",");  // 不转换为数字，保持为字符串
      const res = await fetch(`${process.env.VUE_APP_API_URL}/api/generateTree`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ bfs: bfsArray })
      });
      const data = await res.json();
      this.treeJson = data;
      this.treeData = data;
    }
  }
};
</script>
