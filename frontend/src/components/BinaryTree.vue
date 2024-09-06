<template>
  <div id="tree"></div>
</template>

<script>
import * as d3 from "d3";

export default {
  props: ["treeData"],
  watch: {
    treeData: function (newVal) {
      if (newVal) {
        this.drawTree(newVal);
      }
    }
  },
  mounted() {
    if (this.treeData) {
      this.drawTree(this.treeData);
    }
  },
  methods: {
    drawTree(treeData) {
      // 清空之前的图形
      d3.select("#tree").selectAll("*").remove();

      const width = 600;
      const height = 400;

      const svg = d3
          .select("#tree")
          .append("svg")
          .attr("width", width)
          .attr("height", height);

      const g = svg.append("g").attr("transform", "translate(50, 50)");

      // 调整比例以确保树节点均匀分布
      const treeLayout = d3
          .tree()
          .size([width - 100, height - 100])  // 横向调整大小
          .separation((a, b) => {
            return a.parent === b.parent ? 1 : 2; // 增加兄弟节点间距，确保间隔
          });

      // 构建树数据结构
      const root = d3.hierarchy(treeData, d => {
        const children = [];

        // 如果左子节点存在
        if (d.left) {
          children.push(d.left);
        } else if (d.right) {
          // 如果没有左子节点但是有右子节点，则在左侧填充一个null节点
          children.push({ val: "null" });
        }

        // 如果右子节点存在
        if (d.right) {
          children.push(d.right);
        } else if (d.left) {
          // 如果没有右子节点但是有左子节点，则在右侧填充一个null节点
          children.push({ val: "" });
        }

        return children.length ? children : null;
      });

      treeLayout(root);

      // 绘制连接线
      g.selectAll(".link")
          .data(root.links())
          .enter()
          .append("line")
          .attr("class", "link")
          .attr("x1", d => d.source.x)
          .attr("y1", d => d.source.y)
          .attr("x2", d => d.target.x)
          .attr("y2", d => d.target.y)
          .style("stroke", d => d.target.data.val === "" ? "none" : "#ccc"); // 不绘制连接到null节点的线

      // 绘制节点
      const nodes = g.selectAll(".node")
          .data(root.descendants())
          .enter()
          .append("g")
          .attr("class", "node")
          .attr("transform", d => `translate(${d.x},${d.y})`);

      // 绘制圆形节点，忽略val为"null"的节点
      nodes.append("circle")
          .attr("r", 20)
          .style("fill", "#fff")
          .style("stroke", "black")
          .style("display", d => d.data.val === "" ? "none" : "block");  // 不显示null节点

      // 节点文本，忽略val为"null"的节点
      nodes.append("text")
          .attr("dy", 5)
          .attr("text-anchor", "middle")
          .text(d => d.data.val)
          .style("displasy", d => d.data.val === "" ? "none" : "block");  // 不显示null节点文本
    }
  }
};
</script>

<style scoped>
#tree {
  width: 600px;
  height: 400px;
  border: 1px solid #ccc;
}
</style>
